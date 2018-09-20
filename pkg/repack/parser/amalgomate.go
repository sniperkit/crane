package parser

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	// external
	"github.com/pkg/errors"
)

func Run(cfg Config, outputDir, pkg string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return errors.Wrapf(err, "failed to ensure that output directory exists: %s", outputDir)
	}

	if !filepath.IsAbs(outputDir) {
		wd, err := os.Getwd()
		if err != nil {
			return errors.Wrapf(err, "failed to get working directory")
		}
		outputDir = path.Join(wd, outputDir)
	}

	// repackage main files specified in configuration
	if err := repackage(cfg, outputDir); err != nil {
		return errors.Wrapf(err, "failed to repackage files specified in configuration")
	}

	amalgomatedOutputDir := path.Join(outputDir, internalDir)

	// write output file that imports and uses repackaged files
	if err := writeOutputGoFile(cfg, outputDir, amalgomatedOutputDir, pkg); err != nil {
		return errors.Wrapf(err, "failed to write output file")
	}

	return nil
}

func writeOutputGoFile(config Config, outputDir, amalgomatedOutputDir, packageName string) error {
	fileSet := token.NewFileSet()

	var template string
	if packageName == "main" {
		template = mainTemplate
	} else {
		template = libraryTemplate
	}

	file, err := parser.ParseFile(fileSet, "", template, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "failed to parse template: %s", template)
	}
	file.Name = ast.NewIdent(packageName)

	if err := addImports(file, fileSet, amalgomatedOutputDir, config); err != nil {
		return errors.Wrap(err, "failed to add imports")
	}
	sortImports(file)

	if err := setVarCompositeLiteralElements(file, "programs", createMapLiteralEntries(config.Pkgs)); err != nil {
		return errors.Wrap(err, "failed to add const elements")
	}

	// write output to in-memory buffer and add import spaces
	var byteBuffer bytes.Buffer
	if err := printer.Fprint(&byteBuffer, fileSet, file); err != nil {
		return errors.Wrap(err, "failed to write output file to buffer")
	}
	outputWithSpaces := addImportSpaces(&byteBuffer, importBreakPaths(file))

	// write output to file
	outputFilePath := path.Join(outputDir, packageName+".go")
	if err := ioutil.WriteFile(outputFilePath, outputWithSpaces, 0644); err != nil {
		return errors.Wrapf(err, "failed to write output to path: %s", outputFilePath)
	}

	return nil
}
