package main

/*

Creates bsdiff patch files.

USAGE: bsdiff OLD NEW > patch

*/

import (
	"fmt"
	"os"

	"github.com/kr/binarydist"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "USAGE: bsdiff OLD NEW > PATCH\n\n")
		os.Exit(1)
	}
	old, err := os.Open(os.Args[1])
	must(err)
	new, err := os.Open(os.Args[2])
	must(err)
	err = binarydist.Diff(old, new, os.Stdout)
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
