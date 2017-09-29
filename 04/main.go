// what’s the value of the expression
// (true && false) || (false && true) || !(false && false)

package main

import (
	"fmt"
	"strconv"
)

func main() {
	exp := (true && false) || (false && true) || !(false && false)

	fmt.Println(strconv.FormatBool(exp))
}
