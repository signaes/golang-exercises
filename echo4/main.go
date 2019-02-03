// based on gopl.io/ch2/echo4

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
var length = flag.Bool("l", false, "show output length")

func main() {
	flag.Parse()
	o := strings.Join(flag.Args(), *sep)
	fmt.Println(o)

	if *length {
		fmt.Println(len(o))
	}

	if !*n {
		fmt.Println()
	}
}
