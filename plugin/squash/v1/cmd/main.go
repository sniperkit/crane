package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"

	// external
	"github.com/satori/go.uuid"

	// internal
	squash "github.com/sniperkit/snk.fork.michaelsauter-crane/plugin/squash/v1"
)

func main() {
	var from, input, output, tempdir, tag, tmpRoot string
	var keepTemp, version, last bool
	flag.StringVar(&input, "i", "", "Read from a tar archive file, instead of STDIN")
	flag.StringVar(&output, "o", "", "Write to a file, instead of STDOUT")
	flag.StringVar(&tag, "t", "", "Repository name and tag for new image")
	flag.StringVar(&from, "from", "", "Squash from layer ID (default: first FROM layer)")
	// flag.StringVar(&tmpRoot, "tmpRoot", "", "Tempdir")
	flag.BoolVar(&last, "last", false, "Squash from last found layer ID (Inverts order for automatic root-layer selection")
	flag.BoolVar(&keepTemp, "keepTemp", false, "Keep temp dir when done. (Useful for debugging)")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	flag.BoolVar(&version, "v", false, "Print version information and quit")

	flag.Usage = func() {
		fmt.Printf("\nUsage: docker-squash [options]\n\n")
		fmt.Printf("Squashes the layers of a tar archive on STDIN and streams it to STDOUT\n\n")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	// set verbosity
	squash.Verbose = verbose

	if version {
		fmt.Println(buildVersion)
		return
	}

	var err error

	if tmpRoot != "" {
		rid := uuid.Must(uuid.NewV4())
		fmt.Printf("Declared UUIDv4: %s\n", rid)
		dsTmpBasename := fmt.Sprintf("docker-squash-%s", rid)
		dsTmpDir := path.Join(tmpRoot, dsTmpBasename)
		squash.CreateDirIfNotExist(dsTmpDir)
		tempdir = dsTmpDir
	} else {
		tempdir, err = ioutil.TempDir("", "docker-squash")
	}
	if err != nil {
		squash.Fatal(err)
	}

	// fmt.Println("tempdir: ", tempdir)
	// os.Exit(1)
	// docker save sniperkit/flogo-docker:v0.5.6-snk-devzone-ai | sudo docker-squash -t sniperkit/flogo-docker:v0.5.6-devzone-ai-squashed -verbose | docker load

	var tagRegExp = regexp.MustCompile("^(.*):(.*)$")
	if tag != "" && strings.Contains(tag, ":") {
		parts := tagRegExp.FindStringSubmatch(tag)
		if parts[1] == "" || parts[2] == "" {
			squash.Fatalf("bad tag format: %s\n", tag)
		}
	}

	signals = make(chan os.Signal, 1)

	if !keepTemp {
		wg.Add(1)
		signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGTERM)
		go squash.Shutdown(tempdir)
	}

	export, err := squash.LoadExport(input, tempdir)
	if err != nil {
		squash.Fatal(err)
	}

	// Export may have multiple branches with the same parent.
	// We can't handle that currently so abort.
	for _, v := range export.Repositories {
		commits := map[string]string{}
		for tag, commit := range *v {
			commits[commit] = tag
		}
		if len(commits) > 1 {
			squash.Fatal("This image is a full repository export w/ multiple images in it.  " +
				"You need to generate the export from a specific image ID or tag.")
		}

	}

	var start *ExportedImage
	if last {
		start = export.LastSquash()
		// Can't find a previously squashed layer, use last FROM
		if start == nil {
			start = export.LastFrom()
		}
	} else {
		start = export.FirstSquash()
		// Can't find a previously squashed layer, use first FROM
		if start == nil {
			start = export.FirstFrom()
		}
	}
	// Can't find a FROM, default to root
	if start == nil {
		start = export.Root()
	}

	if from != "" {

		if from == "root" {
			start = export.Root()
		} else {
			start, err = export.GetById(from)
			if err != nil {
				squash.Fatal(err)
			}
		}
	}

	if start == nil {
		squash.Fatalf("no layer matching %s\n", from)
		return
	}

	// extract each "layer.tar" to "layer" dir
	err = export.ExtractLayers()
	if err != nil {
		fatal(err)
		return
	}

	// insert a new layer after our squash point
	newEntry, err := export.InsertLayer(start.LayerConfig.Id)
	if err != nil {
		squash.Fatal(err)
		return
	}

	squash.Debugf("Inserted new layer %s after %s\n", newEntry.LayerConfig.Id[0:12],
		newEntry.LayerConfig.Parent[0:12])

	if verbose {
		e := export.Root()
		for {
			if e == nil {
				break
			}
			cmd := strings.Join(e.LayerConfig.ContainerConfig().Cmd, " ")
			if len(cmd) > 60 {
				cmd = cmd[:60]
			}

			if e.LayerConfig.Id == newEntry.LayerConfig.Id {
				squash.Debugf("  -> %s %s\n", e.LayerConfig.Id[0:12], cmd)
			} else {
				squash.Debugf("  -  %s %s\n", e.LayerConfig.Id[0:12], cmd)
			}
			e = export.ChildOf(e.LayerConfig.Id)
		}
	}

	// squash all later layers into our new layer
	err = export.SquashLayers(newEntry, newEntry)
	if err != nil {
		squash.Fatal(err)
		return
	}

	squash.Debugf("Tarring up squashed layer %s\n", newEntry.LayerConfig.Id[:12])
	// create a layer.tar from our squashed layer
	err = newEntry.TarLayer()
	if err != nil {
		fatal(err)
	}

	squash.Debugf("Removing extracted layers\n")
	// remove our expanded "layer" dirs
	err = export.RemoveExtractedLayers()
	if err != nil {
		squash.Fatal(err)
	}

	if tag != "" {
		tagPart := "latest"
		repoPart := tag
		parts := tagRegExp.FindStringSubmatch(tag)
		if len(parts) > 2 {
			repoPart = parts[1]
			tagPart = parts[2]
		}
		tagInfo := TagInfo{}
		layer := export.LastChild()

		tagInfo[tagPart] = layer.LayerConfig.Id
		export.Repositories[repoPart] = &tagInfo

		squash.Debugf("Tagging %s as %s:%s\n", layer.LayerConfig.Id[0:12], repoPart, tagPart)
		err := export.WriteRepositoriesJson()
		if err != nil {
			fatal(err)
		}
	}

	ow := os.Stdout
	if output != "" {
		var err error
		ow, err = os.Create(output)
		if err != nil {
			fatal(err)
		}
		squash.Debugf("Tarring new image to %s\n", output)
	} else {
		squash.Debugf("Tarring new image to STDOUT\n")
	}

	// Update manifest
	squash.UpdateManifest(export, tag)

	// bundle up the new image
	err = export.TarLayers(ow)
	if err != nil {
		fatal(err)
	}

	squash.Debug("Done. New image created.")
	// print our new history
	export.PrintHistory()

	signals <- os.Interrupt
	wg.Wait()
}
