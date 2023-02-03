package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) < 1 {
		fmt.Println("need arguments: -a or -s")
	}

	if args[1] == "-a" {
		assemmbler_runner()
	}

	if args[1] == "-s" {
		scrambler_runner()
	}

}
