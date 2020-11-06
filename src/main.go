package main

import (
	"flag"
	"fmt"
	"os"
)

var suffix string

func init() {
	flag.StringVar(&suffix, "suffix", "", "suffix to print")
	flag.Parse()
}

func main() {
	world := os.Getenv("world")
	fmt.Printf("hello %s %s\n", world, suffix)
}
