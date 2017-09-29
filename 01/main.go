// write a function which takes an integer.
// The function will have two returns
// The first return should be the argument
// divided by 2. The second return should
// be a bool that let’s us know whether or
// not the argument was even. For example:
// a. half(1) returns (0, false)
// b. half(2) returns (1, true)

package main

import (
	"fmt"
	"strconv"
)

func half(n int64) (int64, bool) {
	return n / 2, n%2 == 0
}

func text(n int64, b bool) string {
	return strconv.FormatInt(n, 10) + ", " + strconv.FormatBool(b)
}

func main() {
	fmt.Println("half(1) = ", text(half(1)))
	fmt.Println("half(2) = ", text(half(2)))
	fmt.Println("half(100) = ", text(half(100)))
	fmt.Println("half(150) = ", text(half(150)))
}
