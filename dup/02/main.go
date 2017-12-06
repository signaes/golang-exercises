package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, lines)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, lines)
			f.Close()
		}
	}

	for line, n := range lines {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, lines map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		lines[input.Text()]++
	}
}
