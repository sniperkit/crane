package core

import (
	"runtime"
	// snk "github.com/sniperkit/snk.fork.michaelsauter-crane/plugin/mount/snk"
)

type AcceleratedMount interface {
	Run()
	Reset()
	Logs(follow bool)
	VolumeArg() string
	Volume() string
}

type acceleratedMount struct {
	RawVolume  string
	configPath string
}

func NewAcceleratedMount(rawVolume, configPath string) *acceleratedMount {
	return &acceleratedMount{
		RawVolume:  rawVolume,
		configPath: configPath,
	}
}

func (am *acceleratedMount) Volume() string {
	return ""
}

func (am *acceleratedMount) Run() {
}

func (am *acceleratedMount) Reset() {
}

func (am *acceleratedMount) Logs(follow bool) {
}

func (am *acceleratedMount) VolumeArg() string {
	return ""
}

var proOnly = "Accelerated bind mounts are not available in the free version, please purchase the pro version: https://www.craneup.tech"

func accelerationEnabled() bool {
	if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
		printInfof("%s\n", proOnly)
	}
	return false
}
