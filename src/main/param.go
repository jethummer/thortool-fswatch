package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help  bool
	path  string
	multiple bool
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.BoolVar(&multiple, "m", false, "permit multiple process")

	flag.StringVar(&path, "w", "..", "watch root path")

	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `fswatch version: 0.0.1
Usage: fswatch [-h] [-m] [-w watchpath]
`)
	flag.PrintDefaults()
}

func ParseArgs() {

	flag.Parse()

	if help || len(path) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	dir = path
	kill = !multiple
}
