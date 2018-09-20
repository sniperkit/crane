// Copyright (c) 2016 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	// external
	"github.com/palantir/pkg/cobracli"
	"github.com/spf13/cobra"

	// internal
	"github.com/sniperkit/snk.fork.michaelsauter-crane/pkg/repack/parser"
)

const (
	debugFlagName     = "debug"
	configFlagName    = "config"
	outputDirFlagName = "output-dir"
	pkgFlagName       = "pkg"
)

var (
	debugFlagVal  bool
	configFlagVal string
	outputDirVal  string
	pkgFlagVal    string
)

// SnkrCmd represents the base command when called without any subcommands
var SnkrCmd = &cobra.Command{
	Use:   "snkr",
	Short: "Re-package main packages into a library package",
	Long: `snkr is used to re-package Go programs with
a "main" package into a library package that can be called
directly in-process.

snkr requires a configuration YML file that specifies
the packages that should be converted from "main" packages
into library packages. An output directory and the name of
the package for the generated source files should also be
specified.

Here is an example configuration file:

packages:
  gofmt:
    main: cmd/gofmt
  crane:
    main: github.com/sniperkit/snk.fork.michaelsauter-crane/cmd/crane
    distance-to-project-pkg: 2
    omit-vendor-dirs: true

An example invocation is of the form:

  snkr --config repack.yml --output-dir your_package_output_dir --pkg your_package_name

This invocation would fusion the inputs specified in "repack.yml" and would
write the generated source into the "your_package_output_dir" directory with the package
name "your_package_name".

examples:
- $ snkr --config config.yml --output-dir docker-squash --pkg main
- $ snkr --config snkr.yml --output-dir mount --pkg mount
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := parser.LoadConfig(configFlagVal)
		if err != nil {
			return err
		}
		return parser.Run(cfg, outputDirVal, pkgFlagVal)
	},
}

func Execute() int {
	return cobracli.ExecuteWithDebugVarAndDefaultParams(SnkrCmd, &debugFlagVal)
}

func init() {
	SnkrCmd.Flags().BoolVar(&debugFlagVal, debugFlagName, false, "run in debugFlagVal mode")

	SnkrCmd.Flags().StringVar(&configFlagVal, configFlagName, "", "configuration file that specifies packages to be amalgomated")
	if err := SnkrCmd.MarkFlagRequired(configFlagName); err != nil {
		panic(err)
	}

	SnkrCmd.Flags().StringVar(&outputDirVal, outputDirFlagName, "", "directory in which amalgomated output is written")
	if err := SnkrCmd.MarkFlagRequired(outputDirFlagName); err != nil {
		panic(err)
	}

	SnkrCmd.Flags().StringVar(&pkgFlagVal, pkgFlagName, "", "package name of the amalgomated source that is generated")
	if err := SnkrCmd.MarkFlagRequired(pkgFlagName); err != nil {
		panic(err)
	}

}
