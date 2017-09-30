// Write a function with one variadic parameter
// that finds the greatest number in a list of numbers

package main

import "fmt"

func main() {
	max := func(nums ...int) int {
		bigger := nums[0]

		for _, num := range nums {
			if num > bigger {
				bigger = num
			}
		}

		return bigger
	}

	s := []int{20, 30, 40, 60}

	fmt.Println(max(s...))
	fmt.Println(max(1, 40, 30, 34, 60, 90, 2))
	fmt.Println(max(1, 400, 30, 34, 60, 90, 2))
	fmt.Println(max(0, 2, 17, 3, 43, 78, 8))
	fmt.Println(max(200, 92, 17, 3, 43, 78, 8))
	fmt.Println(max(200, 92, 17, 3600, 43, 78, 8))
}
