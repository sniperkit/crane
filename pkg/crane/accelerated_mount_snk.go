// +build !pro
package crane

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

/*
	Refs:
	- http://markshust.com/2018/01/30/performance-tuning-docker-mac
	- https://blog.docker.com/2016/06/docker-for-mac-splice/
	- http://markshust.com/2017/03/02/making-docker-mac-faster-overlay2-filesystem
		- docker info | grep "Storage Driver"
		- {"storage-driver": "overlay2"}
	- https://docs.docker.com/docker-for-mac/osxfs-caching/
		- docker run -v /Users/yallop/project:/project:cached alpine command
	- https://docs.docker.com/docker-for-mac/osxfs-caching/
	- https://docs.docker.com/docker-for-mac/osxfs/#performance-issues-solutions-and-roadmap
	- https://stories.amazee.io/docker-on-mac-performance-docker-machine-vs-docker-for-mac-4c64c0afdf99
	- https://docs.docker.com/compose/compose-file/#caching-options-for-volume-mounts-docker-for-mac
*/

type AcceleratedMount interface {
	ContainerName() string
	Volume() string
	Autostart() bool
	Exists() bool
	Running() bool
	Start(debug bool)
	Stop()
	Status() string
	// 2.x
	Run()
	Reset()
	Logs(follow bool)
	VolumeArg() string
}

type acceleratedMount struct {
	RawVolume    string
	RawFlags     []string `json:"flags" yaml:"flags"`
	RawImage     string   `json:"image" yaml:"image"`
	Uid          int      `json:"uid" yaml:"uid"`
	Gid          int      `json:"gid" yaml:"gid"`
	RawAutostart bool     `json:"autostart" yaml:"autostart"`
	configPath   string
	cName        string
	volume       string
}

func accelerationEnabled() bool {
	/*
	   if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
	       printInfof("%s\n", proOnly)
	   }
	*/
	return true
}

func (a *acceleratedMount) ContainerName() string {
	if a.cName == "" {
		syncIdentifierParts := []string{
			a.configPath,
			a.Volume(),
			a.image(),
			strings.Join(a.flags(), " "),
			strconv.Itoa(a.Uid),
			strconv.Itoa(a.Gid),
		}
		syncIdentifier := []byte(strings.Join(syncIdentifierParts, ":"))
		digest := fmt.Sprintf("%x", md5.Sum(syncIdentifier))
		a.cName = "crane_unison_" + digest
	}
	return a.cName
}

func (a *acceleratedMount) Volume() string {
	if a.volume == "" {
		v := expandEnv(a.RawVolume)
		parts := strings.Split(v, ":")
		if !path.IsAbs(parts[0]) {
			parts[0] = a.configPath + "/" + parts[0]
		}
		a.volume = strings.Join(parts, ":")
	}
	return a.volume
}

func (a *acceleratedMount) Autostart() bool {
	return a.RawAutostart
}

func (a *acceleratedMount) Exists() bool {
	return containerID(a.ContainerName()) != ""
}

func (a *acceleratedMount) Running() bool {
	return a.serverRunning() && a.clientRunning()
}

func (a *acceleratedMount) Status() string {
	status := "-"
	if a.Exists() {
		status = "stopped"
		if a.Running() {
			status = "running"
		}
	}
	return status
}

func (a *acceleratedMount) Start(debug bool) {
	unisonArgs := []string{}

	// Start sync container if needed
	if a.Exists() {
		if a.serverRunning() {
			verboseLog("Unison sync server for " + a.hostDir() + " already running")
		} else {
			verboseLog("Starting unison sync server for " + a.hostDir())
			dockerArgs := []string{"start", a.ContainerName()}
			executeHiddenCommand("docker", dockerArgs)
		}
	} else {
		checkUnisonRequirements()
		verboseLog("Starting unison sync server for " + a.hostDir())
		dockerArgs := []string{
			"run",
			"--name", a.ContainerName(),
			"-d",
			"-P",
			"-e", "UNISON_DIR=" + a.containerDir(),
			"-e", "UNISON_UID=" + strconv.Itoa(a.Uid),
			"-e", "UNISON_GID=" + strconv.Itoa(a.Gid),
			"-v", a.containerDir() + ":cached",
			a.image(),
		}

		/*
			- mac, osxfs, volumes
			- docker run -v /Users/yallop/project:/project:cached alpine command
			- docker run -v /Users/yallop/project:/project:cached  -v /host/another-path:/mount/another-point:consistent alpine command
		*/

		executeHiddenCommand("docker", dockerArgs)
		fmt.Printf("Doing initial snyc for %s ...\n", a.hostDir())
		unisonArgs = a.unisonArgs()
		initialSyncArgs := []string{}
		for _, a := range unisonArgs {
			if !strings.HasPrefix(a, "-repeat") {
				initialSyncArgs = append(initialSyncArgs, a)
			}
		}
		// Wait a bit for the Unison server to start
		time.Sleep(3 * time.Second)
		if debug {
			executeCommand("unison", initialSyncArgs, os.Stdout, os.Stderr)
		} else {
			executeCommand("unison", initialSyncArgs, nil, nil)
		}
	}

	// Start unison in background if not already running
	if a.clientRunning() {
		verboseLog("Unison sync client for " + a.hostDir() + " already running")
	} else {
		verboseLog("Starting unison sync client for " + a.hostDir())
		unisonArgs = a.unisonArgs()
		if !debug { // Line will be logged later within executeCommand anyway
			verboseLog("unison " + strings.Join(unisonArgs, " "))
		}
		if !isDryRun() {
			// Wait a bit for the Unison server to start
			time.Sleep(3 * time.Second)
			if debug {
				executeCommand("unison", unisonArgs, os.Stdout, os.Stderr)
			} else {
				cmd := exec.Command("unison", unisonArgs...)
				cmd.Dir = cfg.Path()
				cmd.Stdout = nil
				cmd.Stderr = nil
				cmd.Stdin = nil
				cmd.Start()
			}
		}
	}
}

func (a *acceleratedMount) Run() {
}

func (a *acceleratedMount) Reset() {
}

func (a *acceleratedMount) Logs(follow bool) {
}

func (a *acceleratedMount) VolumeArg() string {
	return ""
}

func (a *acceleratedMount) Stop() {
	verboseLog("Stopping unison sync for " + a.hostDir())

	// stop container (also stops Unison sync)
	dockerArgs := []string{"kill", a.ContainerName()}
	executeHiddenCommand("docker", dockerArgs)
}

func (a *acceleratedMount) serverRunning() bool {
	return a.Exists() && inspectBool(a.ContainerName(), "{{.State.Running}}")
}

func (a *acceleratedMount) clientRunning() bool {
	args := []string{"-f", "unison " + a.hostDir() + " socket://localhost:" + a.publishedPort()}
	_, err := commandOutput("pgrep", args)
	return err == nil
}

func (a *acceleratedMount) unisonArgs() []string {
	unisonArgs := []string{a.hostDir(), "socket://localhost:" + a.publishedPort() + "/"}
	return append(unisonArgs, a.flags()...)
}

func (a *acceleratedMount) image() string {
	if len(a.RawImage) > 0 {
		return expandEnv(a.RawImage)
	}
	allowedVersions := []string{"2.48.4"}
	versionOut, err := commandOutput("unison", []string{"-version"})
	if err != nil {
		return "michaelsauter/unison:2.48.4"
		// return "michaelsauter/crane-sync:latest" // 3.2.0
	}
	// `unison -version` returns sth like "unison version 2.48.4"
	versionParts := strings.Split(versionOut, " ")
	installedVersion := versionParts[len(versionParts)-1]
	if !includes(allowedVersions, installedVersion) {
		panic(StatusError{errors.New("Unison version " + installedVersion + " is not supported. You need to install: " + strings.Join(allowedVersions, ", ")), 69})
	}
	return "michaelsauter/unison:" + installedVersion
}

func (a *acceleratedMount) flags() []string {
	if len(a.RawFlags) > 0 {
		f := []string{}
		for _, rawFlag := range a.RawFlags {
			f = append(f, expandEnv(rawFlag))
		}
		return f
	}
	return []string{"-auto", "-batch", "-ignore=Name {.git}", "-confirmbigdel=false", "-prefer=newer", "-repeat=watch"}
}

func (a *acceleratedMount) hostDir() string {
	parts := strings.Split(a.Volume(), ":")
	return parts[0]
}

func (a *acceleratedMount) containerDir() string {
	parts := strings.Split(a.Volume(), ":")
	return parts[1]
}

func (a *acceleratedMount) publishedPort() string {
	args := []string{"port", a.ContainerName(), "5000/tcp"}
	published, err := commandOutput("docker", args)
	if err != nil {
		printErrorf("Could not detect port of container %a. Sync will not work properly.", a.ContainerName())
		return ""
	}
	parts := strings.Split(published, ":")
	return parts[1]
}

func checkUnisonRequirements() {
	_, err := commandOutput("which", []string{"unison"})
	if err != nil {
		panic(StatusError{errors.New("`unison` is not installed or not in your $PATH.\nSee https://github.com/michaelsauter/crane/wiki/Unison-installation."), 69})
	}

	_, err = commandOutput("which", []string{"unison-fsmonitor"})
	if err != nil {
		panic(StatusError{errors.New("`unison-fsmonitor` is not installed or not in your $PATH.\nSee https://github.com/michaelsauter/crane/wiki/Unison-installation."), 69})
	}
}
