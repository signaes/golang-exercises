/*
	Write a function, foo, which can be called in
	all of these ways

	func main() {
		foo(1, 2)
		foo(1, 2, 3)
		aSlice := []int{1, 2, 3, 4}
		foo(aSlice...)
		foo()
	}
*/

package main

import "fmt"

func foo(numbers ...int) int {
	var total int

	if len(numbers) > 0 {
		for _, number := range numbers {
			total += number
		}
	}

	fmt.Println(total)
	return total
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
