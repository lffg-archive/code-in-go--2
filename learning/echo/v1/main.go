package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	fmt.Println(s)
}
