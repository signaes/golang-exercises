package main

import (
	"fmt"
	"os"
)

func main() {
	var s, space string

	// os.Args -> []string
	for _, arg := range os.Args[1:] {
		s += space + arg
		space = " "
	}

	fmt.Println(s)
}
