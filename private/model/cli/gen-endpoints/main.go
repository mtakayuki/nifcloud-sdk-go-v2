// +build codegen

// Command gen-endpoints parses a JSON description of the NIFCLOUD endpoint
// discovery logic and generates a Go file which returns an endpoint.

// NOTE: This file was imported from https://github.com/aws/aws-sdk-go-v2/blob/v2.0.0-preview.5/private/model/cli/gen-endpoints/main.go

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
)

// Generates the endpoints from json description
//
// Args:
//  -model The definition file to use
//  -out The output file to generate
func main() {
	var modelName, outName string
	flag.StringVar(&modelName, "model", "", "Endpoints definition model")
	flag.StringVar(&outName, "out", "", "File to write generated endpoints to.")
	flag.Parse()

	if len(modelName) == 0 || len(outName) == 0 {
		exitErrorf("model and out both required.")
	}

	modelFile, err := os.Open(modelName)
	if err != nil {
		exitErrorf("failed to open model definition, %v.", err)
	}
	defer modelFile.Close()

	outFile, err := os.Create(outName)
	if err != nil {
		exitErrorf("failed to open out file, %v.", err)
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			exitErrorf("failed to successfully write %q file, %v", outName, err)
		}
	}()

	if err := endpoints.CodeGenModel(modelFile, outFile); err != nil {
		exitErrorf("failed to codegen model, %v", err)
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
