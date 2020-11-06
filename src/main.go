package main

import (
	"fmt"
	"os"
)

func main() {
	world := os.Getenv("world")
	fmt.Printf("hello %s\n", world)
}
