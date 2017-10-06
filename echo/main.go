package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.Args -> []string
	fmt.Println(strings.Join(os.Args[1:], " "))
}
