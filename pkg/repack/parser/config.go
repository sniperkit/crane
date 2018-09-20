package parser

import (
	"io/ioutil"

	// external
	"github.com/jinzhu/configor"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type SrcPkg struct {
	// default
	MainPkg              string `yaml:"main" json:"main" toml:"main" ini:"main"`
	DistanceToProjectPkg int    `yaml:"distance-to-project-pkg" json:"distance-to-project-pkg" toml:"distance-to-project-pkg" ini:"distance-to-project-pkg"`
	OmitVendorDirs       bool   `yaml:"omit-vendor-dirs" json:"omit-vendor-dirs" toml:"omit-vendor-dirs" ini:"omit-vendor-dirs"`
}

type ExtraRules struct {
	// extra
	Vendors VendorRules `yaml:"vendors" json:"vendors" toml:"vendors" ini:"vendors"`
	Assets  AssetRules  `yaml:"assets" json:"assets" toml:"assets" ini:"assets"`
	Outputs OutputRules `yaml:"outputs" json:"outputs" toml:"outputs" ini:"outputs"`
}

type VendorRules struct {
	OmitVendorDirs bool     `yaml:"omit-vendor-dirs" json:"omit-vendor-dirs" toml:"omit-vendor-dirs" ini:"omit-vendor-dirs"`
	Exclude        []string `yaml:"exclude-vendor-dirs" json:"exclude-vendor-dirs" toml:"exclude-vendor-dirs" ini:"exclude-vendor-dirs"`
}

type AssetRules struct {
	Bindata bool     `yaml:"bindata" json:"bindata" toml:"bindata" ini:"bindata"`
	Dirs    []string `yaml:"dirs" json:"dirs" toml:"dirs" ini:"dirs"`
	Exclude []string `yaml:"exclude" json:"exclude" toml:"exclude" ini:"exclude"`
}

type OutputRules struct {
	Templates []string `yaml:"templates" json:"templates" toml:"templates" ini:"templates"`
}

type Config struct {
	Pkgs map[string]SrcPkg `yaml:"packages" json:"packages" toml:"packages" ini:"packages"`
}

func LoadConfigor(configPaths ...string) (Config, error) {
	var cfg Config
	configor.New(&configor.Config{
		Environment: "production",
		Debug:       true,
		Verbose:     true,
	}).Load(&cfg, configPaths...)
	pp.Println("Config: %#v", cfg)
	return cfg, nil
}

func LoadConfig(configPath string) (Config, error) {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, errors.Wrapf(err, "failed to read file %s", configPath)
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return Config{}, errors.Wrapf(err, "failed to unmarshal file %s", configPath)
	}

	if len(cfg.Pkgs) == 0 {
		return Config{}, errors.Errorf("configuration read from file %s with content %q was empty", configPath, string(file))
	}

	for name, pkg := range cfg.Pkgs {
		if name == "" {
			return Config{}, errors.Errorf("config cannot contain a blank name: %v", cfg)
		}

		if pkg.MainPkg == "" {
			return Config{}, errors.Errorf("config for package %s had a blank main package directory: %v", name, cfg)
		}
	}
	return cfg, nil
}
