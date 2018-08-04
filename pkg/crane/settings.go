package crane

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	// external
	"github.com/cep21/xdgbasedir"
	"github.com/go-yaml/yaml"
	uuid "github.com/hashicorp/go-uuid"
	"github.com/jinzhu/configor"
	"github.com/k0kubun/pp"
)

var (
	settings              *Settings
	settingsDirectoryPath string
)

type Settings struct {
	filename        string                    `json:"-" yaml:"-"`
	UUID            string                    `json:"uuid" yaml:"uuid"`
	Version         string                    `json:"version" yaml:"version"`
	LatestVersion   string                    `json:"latest_version" yaml:"latest_version"`
	NextUpdateCheck time.Time                 `json:"next_update_check" yaml:"next_update_check"`
	CheckForUpdates bool                      `json:"check_for_updates" yaml:"check_for_updates"`
	DatabasePath    string                    `default:"~/.config/crane/data" json:"database_path,omitempty" yaml:"database_path,omitempty"`
	IndexPath       string                    `default:"~/.config/crane/index" json:"index_path,omitempty" yaml:"index_path,omitempty"`
	Services        map[string]*ServiceConfig `json:"services,omitempty" yaml:"services,omitempty"`
	Outputs         map[string]*OutputConfig  `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}

func init() {
	baseDir, err := xdgbasedir.ConfigHomeDirectory()
	if err != nil {
		log.Fatal("Can't find XDG BaseDirectory")
	} else {
		settingsDirectoryPath = path.Join(baseDir, ProgramName)
	}
}

// Determine crane settings base path.
// On windows, this is %APPDATA%\\crane
// On unix, this is ${XDG_CONFIG_HOME}/crane (which usually
// is ${HOME}/.config)
func settingsPath() (string, error) {
	settingsPath := os.Getenv("CRANE_SETTINGS_PATH")
	if len(settingsPath) > 0 {
		return settingsPath, nil
	}
	if runtime.GOOS == "windows" {
		settingsPath = os.Getenv("APPDATA")
		if len(settingsPath) > 0 {
			return fmt.Sprintf("%s/crane", settingsPath), nil
		}
		return "", errors.New("Cannot detect settings path!")
	}
	settingsPath = os.Getenv("XDG_CONFIG_HOME")
	if len(settingsPath) > 0 {
		return fmt.Sprintf("%s/crane", settingsPath), nil
	}
	homeDir := os.Getenv("HOME")
	if len(homeDir) > 0 {
		return fmt.Sprintf("%s/.config/crane", homeDir), nil
	}
	return "", errors.New("Cannot detect settings path!")
}

func createSettings(filename string) error {
	uuid, _ := uuid.GenerateUUID()
	settings = &Settings{
		filename:        filename,
		UUID:            uuid,
		Version:         Version,
		LatestVersion:   Version,
		NextUpdateCheck: time.Now().Add(autoUpdateCheckInterval()),
		CheckForUpdates: true,
	}
	msg := fmt.Sprintf("Writing settings file to %s\n", filename)
	printInfof(msg)
	return settings.Write(filename)
}

func loadSettings(settingFiles ...string) {
	if len(settingFiles) == 0 {
		settingFiles = append(settingFiles, "config.yml")
	}
	var settings *Settings
	configor.Load(settings, settingFiles...)
	pp.Println("settings: %#v", settings)
}

func readSettings() error {
	// Determine settings path
	sp, err := settingsPath()
	if err != nil {
		return err
	}

	// Create settings path if it does not exist yet
	if _, err := os.Stat(sp); err != nil {
		os.MkdirAll(sp, os.ModePerm)
		if _, err := os.Stat(sp); err != nil {
			return err
		}
	}

	// Create settings file if it does not exist yet
	filename := filepath.Join(sp, "config.json")
	if _, err := os.Stat(filename); err != nil {
		return createSettings(filename)
	}

	// read settings of file
	settings = &Settings{filename: filename}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, settings)
}

func (s *Settings) ShouldCheckForUpdates() bool {
	if !s.CheckForUpdates {
		return false
	}
	return time.Now().After(settings.NextUpdateCheck)
}

// If version in settings does not match version of binary,
// we assume that the binary was updated and update the
// settings file with the new information.
func (s *Settings) CorrectVersion() error {
	if Version != s.Version {
		s.Version = Version
		return s.Update(Version)
	}
	return nil
}

func (s *Settings) Update(latestVersion string) error {
	s.NextUpdateCheck = time.Now().Add(autoUpdateCheckInterval())
	s.LatestVersion = latestVersion
	return s.Write(s.filename)
}

func (s *Settings) DelayNextUpdateCheck() error {
	s.NextUpdateCheck = time.Now().Add(time.Hour)
	return s.Write(s.filename)
}

func (s *Settings) Write(filename string) error {
	contents, _ := json.Marshal(s)
	return ioutil.WriteFile(filename, contents, 0644)
}

/*
func (s *Settings) WriteConfig() error {
	err := os.MkdirAll(settingsDirectoryPath, 0700)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(s)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(settingsFilePath(), data, 0600)
}

func settingsFilePath() string {
	return path.Join(settingsDirectoryPath, fmt.Sprintf("%s.settings.yml", ProgramName))
}
*/
