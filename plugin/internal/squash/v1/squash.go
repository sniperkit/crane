package squash

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"strings"
	"sync"
	"syscall"
)

var (
	buildVersion string
	signals      chan os.Signal
	wg           sync.WaitGroup
)

func Shutdown(tempdir string) {
	defer wg.Done()
	<-signals
	Debugf("Removing tempdir %s\n", tempdir)
	err := os.RemoveAll(tempdir)
	if err != nil {
		fatal(err)
	}

}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			Panic(err)
		}
	}
}

func UpdateManifest(export *Export, tag string) {
	manifestor, err := NewManifestor(export, tag)
	if err != nil {
		Fatal(err)
	}

	manifestor.GenerateLayers()
	manifestor.UpdateManifest()
	manifestor.UpdateConfig()

	err = manifestor.SaveChanges()
	if err != nil {
		Fatal(err)
	}
}
