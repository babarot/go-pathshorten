package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/b4b4r07/go-pathshorten"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "pathshorten: too few arguments\n")
		os.Exit(1)
	}

	for _, path := range flag.Args() {
		fmt.Println(pathshorten.Run(path))
	}
}
