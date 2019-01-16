package main

import "fmt"

func main() {
	const WEEKS = 52
	x := []int{1, 2, 3, 4, 5, 6, 7}

	weeksList := []int{}

	fmt.Println(x)

	for i := 0; i < WEEKS; i++ {
		weeksList = append(x, weeksList...)
	}

	fmt.Println(weeksList)
	fmt.Println(len(weeksList))

	y := append(weeksList[:3], weeksList[100:]...)

	fmt.Println(y)

	var year = make([]int, 7, 365)

	fmt.Println(year)
	fmt.Println(len(year))
	fmt.Println(cap(year))

	z := [][]string{[]string{"A", "B"}, []string{"C", "D"}}

	fmt.Println(z)
}
