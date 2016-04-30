package main

import (
	"fmt"
	"os"

	"github.com/kr/binarydist"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: bsdiff OLD NEW > PATCH\n")
		os.Exit(1)
	}
	old, err := os.Open(os.Args[0])
	must(err)
	new, err := os.Open(os.Args[1])
	must(err)
	binarydist.Diff(old, new, os.Stdout)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
