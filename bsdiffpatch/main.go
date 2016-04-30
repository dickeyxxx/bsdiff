package main

/*

Patch a file with a bsdiff patch

USAGE: bsdiff BINARY PATCH > BINARY

*/

import (
	"fmt"
	"os"

	"github.com/kr/binarydist"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "USAGE: bsdiff BINARY PATCH > BINARY\n")
		os.Exit(1)
	}
	old, err := os.Open(os.Args[1])
	must(err)
	patch, err := os.Open(os.Args[2])
	must(err)
	err = binarydist.Patch(old, os.Stdout, patch)
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
