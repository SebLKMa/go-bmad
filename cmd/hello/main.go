package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: hello <name>")
		os.Exit(1)
	}
	fmt.Printf("hello %s\n", os.Args[1])
}
