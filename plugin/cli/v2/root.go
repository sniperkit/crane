package cmd

import (
	"fmt"
	"os"

	// internal
	"github.com/sniperkit/crane/pkg/crane"
	// "github.com/sniperkit/crane/plugin/service"
	// "github.com/sniperkit/crane/plugin/pipeline"
	// "github.com/sniperkit/crane/plugin/macro"

	"github.com/blevesearch/bleve"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var (
	allowed []string
	cfg     crane.Config
	db      *gorm.DB
	index   bleve.Index
)

var options struct {
	language string
	output   string
	service  string
	tag      string
	verbose  bool
}

// RootCmd is the root command for limo
var RootCmd = &cobra.Command{
	Use:   "crane",
	Short: "Lift containers with ease - https://www.craneup.tech",
	Long:  `crane allows you to manage your ocker containers.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.StringVarP(&options.language, "language", "l", "", "language")
	flags.StringVarP(&options.output, "output", "o", "color", "output type")
	flags.StringVarP(&options.service, "service", "s", "github", "service")
	flags.StringVarP(&options.tag, "tag", "t", "", "tag")
	flags.BoolVarP(&options.verbose, "verbose", "v", false, "verbose output")
}

/*
func getSettings() (*crane.Settings, error) {
    if settings == nil {
        var err error
        if settings, err = crane.ReadConfig(); err != nil {
            return nil, err
        }
    }
    return configuration, nil
}

func getDatabase() (*gorm.DB, error) {
    if db == nil {
        cfg, err := getConfiguration()
        if err != nil {
            return nil, err
        }
        db, err = model.InitDB(cfg.DatabasePath, options.verbose)
        if err != nil {
            return nil, err
        }
    }
    return db, nil
}

func getIndex() (bleve.Index, error) {
    if index == nil {
        cfg, err := getConfiguration()
        if err != nil {
            return nil, err
        }
        index, err = model.InitIndex(cfg.IndexPath)
        if err != nil {
            return nil, err
        }
    }
    return index, nil
}

func getOutput() output.Output {
    o := output.ForName(options.output)
    oc, err := getConfiguration()
    if err == nil {
        o.Configure(oc.GetOutput(options.output))
    }
    return o
}

func getService() (service.Service, error) {
    return service.ForName(options.service)
}

func fatalOnError(err error) {
    if err != nil {
        getOutput().Fatal(err.Error())
    }
}
*/
