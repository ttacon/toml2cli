package main

import (
	"flag"
	"log"
)

var (
	inFile  = flag.String("in-file", "", "file to generate a CLI for")
	outFile = flag.String("out-file", "", "file to output generated CLI to")
)

// One day this will bootstrap itself...
func main() {
	flag.Parse()

	if len(*inFile) == 0 {
		log.Fatal("must provide the file to generate CLI from")
	}
	// An empty outfile equals stdout.

	if err := processAndGenerate(*inFile, *outFile); err != nil {
		log.Fatal(err)
	}
}
