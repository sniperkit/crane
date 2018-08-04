package crane

/*
import (
	"bytes"
	_ "crypto/sha256"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/opencontainers/go-digest"
	"github.com/opencontainers/image-spec/specs-go/v1"
)

// TODO: need a refactor to split the big function to several
// small functions
func doConvert(in io.Reader, out string) (retErr error) {
	tmpDir, err := ioutil.TempDir("", "docker2oci-docker-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	if err := unpack(tmpDir, in); err != nil {
		return err
	}

	manifestPath := filepath.Join(tmpDir, manifestFileName)
	manifestFile, err := os.Open(manifestPath)
	if err != nil {
		return err
	}
	defer manifestFile.Close()

	var manifests []manifestItem
	if err := json.NewDecoder(manifestFile).Decode(&manifests); err != nil {
		return err
	}

	if err := createLayoutFile(out); err != nil {
		return err
	}

	var index v1.Index
	index.SchemaVersion = 2

	for _, m := range manifests {
		var manifest v1.Manifest

		manifest.SchemaVersion = 2

		configPath := filepath.Join(tmpDir, m.Config)
		config, err := ioutil.ReadFile(configPath)
		if err != nil {
			return err
		}
		img, err := NewFromJSON(config)
		if err != nil {
			return err
		}
		ociConfig := v1.Image{
			Created:      &img.Created,
			Author:       img.Author,
			Architecture: img.Architecture,
			OS:           img.OS,
			Config: v1.ImageConfig{
				User:         img.Config.User,
				ExposedPorts: img.Config.ExposedPorts,
				Env:          img.Config.Env,
				Entrypoint:   []string(img.Config.Entrypoint),
				Cmd:          []string(img.Config.Cmd),
				Volumes:      img.Config.Volumes,
				WorkingDir:   img.Config.WorkingDir,
				Labels:       img.Config.Labels,
				StopSignal:   img.Config.StopSignal,
			},
			RootFS: v1.RootFS{
				Type:    img.RootFS.Type,
				DiffIDs: img.RootFS.DiffIDs,
			},
			History: img.History,
		}
		des, err := createConfigFile(out, ociConfig)
		if err != nil {
			return err
		}
		des.MediaType = v1.MediaTypeImageConfig
		manifest.Config = des
		for i, _ := range img.RootFS.DiffIDs {
			layerPath := filepath.Join(tmpDir, m.Layers[i])
			f, err := os.Open(layerPath)
			if err != nil {
				return err
			}
			defer f.Close()
			des, err := createLayerBlob(out, f)
			if err != nil {
				return err
			}
			// TODO: detect the tar format, so we know the mediaType
			des.MediaType = v1.MediaTypeImageLayer
			manifest.Layers = append(manifest.Layers, des)
		}
		des, err = createManifestFile(out, manifest)
		if err != nil {
			return err
		}
		des.MediaType = v1.MediaTypeImageManifest
		des.Platform = &v1.Platform{
			Architecture: ociConfig.Architecture,
			OS:           ociConfig.OS,
		}
		des.Annotations = make(map[string]string)
		// FIXME: a image may have multiple tags
		// TODO: validate the tag
		for _, tag := range m.RepoTags {
			strs := strings.Split(tag, ":")
			if len(strs) != 2 {
				continue
			}
			des.Annotations["org.opencontainers.image.ref.name"] = strs[1]
		}
		index.Manifests = append(index.Manifests, des)
	}
	err = createIndexFile(out, index)
	if err != nil {
		return err
	}

	return nil
}

func createLayoutFile(root string) error {
	var layout v1.ImageLayout
	layout.Version = v1.ImageLayoutVersion
	contents, err := json.Marshal(layout)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(root, v1.ImageLayoutFile), contents, 0644)
}

func createLayerBlob(root string, inTar io.Reader) (v1.Descriptor, error) {
	return createBlob(root, inTar)
}

func createIndexFile(root string, index v1.Index) error {
	content, err := json.Marshal(index)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(root, "index.json"), content, 0644)
}

func createManifestFile(root string, manifest v1.Manifest) (v1.Descriptor, error) {
	content, err := json.Marshal(manifest)
	if err != nil {
		return v1.Descriptor{}, err
	}

	return createBlob(root, bytes.NewBuffer(content))
}

func createConfigFile(root string, config v1.Image) (v1.Descriptor, error) {
	content, err := json.Marshal(config)
	if err != nil {
		return v1.Descriptor{}, err
	}

	return createBlob(root, bytes.NewBuffer(content))
}

func createBlob(root string, stream io.Reader) (v1.Descriptor, error) {
	name := filepath.Join(root, "blobs", "sha256", ".tmp-blob")
	err := os.MkdirAll(filepath.Dir(name), 0700)
	if err != nil {
		return v1.Descriptor{}, err
	}

	f, err := os.Create(name)
	if err != nil {
		return v1.Descriptor{}, err
	}
	defer f.Close()

	digester := digest.SHA256.Digester()
	tee := io.TeeReader(stream, digester.Hash())
	size, err := io.Copy(f, tee)
	if err != nil {
		return v1.Descriptor{}, err
	}

	if err := f.Sync(); err != nil {
		return v1.Descriptor{}, err
	}

	if err := f.Chmod(0644); err != nil {
		return v1.Descriptor{}, err
	}

	if err := digester.Digest().Validate(); err != nil {
		return v1.Descriptor{}, err
	}

	err = os.Rename(name, filepath.Join(filepath.Dir(name), digester.Digest().Hex()))
	if err != nil {
		return v1.Descriptor{}, err
	}

	return v1.Descriptor{
		Digest: digester.Digest(),
		Size:   size,
	}, nil
}
*/
