package squash

// This part is an extract from https://github.com/jwilder/docker-squash/
import (
	"fmt"
	"os"
	"os/exec"

	jww "github.com/spf13/jwalterweatherman"
)

type ExportedImage struct {
	Path         string
	JsonPath     string
	VersionPath  string
	LayerTarPath string
	LayerDirPath string
}

func (e *ExportedImage) CreateDirs() error {
	return os.MkdirAll(e.Path, 0755)
}

func (e *ExportedImage) TarLayer() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Chdir(e.LayerDirPath)
	if err != nil {
		return err
	}
	defer os.Chdir(cwd)

	cmd := exec.Command("sudo", "/bin/sh", "-c", fmt.Sprintf("%s cvf ../layer.tar ./", TarCmd))
	out, err := cmd.CombinedOutput()
	if err != nil {
		jww.INFO.Println(out)
		return err
	}
	return nil
}

func (e *ExportedImage) RemoveLayerDir() error {
	return os.RemoveAll(e.LayerDirPath)
}

func (e *ExportedImage) ExtractLayerDir() error {
	err := os.MkdirAll(e.LayerDirPath, 0755)
	if err != nil {
		return err
	}

	out, err := extractTar(e.LayerTarPath, e.LayerDirPath)
	if err != nil {
		jww.INFO.Println(out)
		return err
	}
	return nil
}
