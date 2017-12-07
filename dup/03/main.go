package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := make(map[string]map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		lines := make(map[string]int)

		for _, line := range strings.Split(string(data), "\n") {
			lines[line]++
		}

		files[filename] = lines
	}

	for file, lines := range files {
		for line, n := range lines {
			if n > 1 {
				fmt.Printf("%s: %d ocurrences of: \"%s\"\n", file, n, string(line))
			}
		}
	}
}
