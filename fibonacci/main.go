package main

import "fmt"

func fibonacci(limit int) []int {
	var first, second int = 0, 1
	a := make([]int, limit)

	for i := 0; i < limit; i++ {
		if i == 1 {
			first = a[i-1]
			second = a[i-1]
		} else if i > 1 {
			first = a[i-2]
			second = a[i-1]
		}

		sum := first + second
		a[i] = sum
	}

	return a
}

func main() {
	fmt.Println(fibonacci(10))
}
