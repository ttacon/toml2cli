package main

import (
	"errors"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// processAndGenerate processes `inFile` and writes the generated `main.go` file
// contents to `outFile`. It returns an error if:
//
//  - We can't read or process the input file.
//  - We can't generate the `main.go` file (i.e. invalid generator).
//  - We can't output the generated file.
func processAndGenerate(inFile, outFile string) error {
	raw, err := ioutil.ReadFile(inFile)
	if err != nil {
		return err
	}

	var m MetaInfo
	if err := toml.Unmarshal(raw, &m); err != nil {
		return err
	}

	return generateOutFile(outFile, &m)
}

// MetaInfo contains the information needed to generate the CLI boilerplate.
// It consists of two main sections:
//
//  - General info: general information about the top level CLI entrypoint.
//  - Commands: information about the commands to generate.
type MetaInfo struct {
	General  map[string]string        `toml:"general"`
	Commands []map[string]interface{} `toml:"command"`
}

// CLIGenerator is the function interface that a new CLI generator needs to
// implement.
type CLIGenerator func(string, *MetaInfo) error

// Register all generators here.
var knownGenerators = map[string]CLIGenerator{
	"github.com/urfave/cli":      urfaveGenerator,
	"github.com/abiosoft/ishell": abiosoftIShellGenerator,
}

var (
	// ErrUnknownGenerator is the error returned when we try to generate a
	// CLI from a config that references an unknown generator.
	ErrUnknownGenerator = errors.New("unknown generator")

	// ErrNoGeneratorProvided is the error that is returned when no
	// generator in the `[general]` section of the found config file.
	ErrNoGeneratorProvided = errors.New("no generator provided")
)

// generateOutFile calls out to the config's generator to generate the CLI
// boilerplate.
func generateOutFile(outFile string, m *MetaInfo) error {
	generatorName, ok := m.General["generator"]
	if !ok {
		return ErrNoGeneratorProvided
	}

	generator, ok := knownGenerators[generatorName]
	if !ok {
		return ErrUnknownGenerator
	}

	return generator(outFile, m)
}
