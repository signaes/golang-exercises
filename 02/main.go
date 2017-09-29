// Modify the previous program to use a func expression

package main

import (
	"fmt"
	"strconv"
)

func main() {
	half := func(n int64) (int64, bool) {
		return n / 2, n%2 == 0
	}

	text := func(n int64, b bool) string {
		return strconv.FormatInt(n, 10) + ", " + strconv.FormatBool(b)
	}

	fmt.Println("half(1) = ", text(half(1)))
	fmt.Println("half(2) = ", text(half(2)))
	fmt.Println("half(100) = ", text(half(100)))
	fmt.Println("half(150) = ", text(half(150)))
}
