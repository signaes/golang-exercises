package main

import "fmt"

func main() {
	const DAYS = 365
	const WEEKS = DAYS / 7
	const MONTHS = DAYS / 30
	var days [DAYS]int
	var weeks [DAYS / 7]int

	months := [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for i, _ := range days {
		days[i] = i + 1
	}

	for i, _ := range weeks {
		weeks[i] = i + 1
	}

	fmt.Println(days)
	fmt.Println(weeks)
	fmt.Println(MONTHS)
	fmt.Println(months)
}
