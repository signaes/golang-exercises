package main

import "fmt"

func main() {
	func(arg string) {
		fmt.Println("Anonymous", arg)
	}("ok")

	num := func(i int) int {
		y := 1

		for x := i; x < 100; x++ {
			y += x
		}

		return y
	}(20)

	f := func(x int, y int) []int {
		return []int{x, y}
	}

	add := func(a int, b ...int) int {
		x := a

		for _, n := range b {
			x += n
		}

		return x
	}

	add1 := func(x int) int {
		return add(1, x)
	}

	fmt.Println(add(1, 4, 6, 7, 8))
	fmt.Println(add1(2))

	fmt.Println(num)

	fmt.Println(f(10, 20))

	fmt.Println(refun()())

	frf := func() func() int {
		return funrefun()
	}

	frfr := frf()

	fmt.Printf("%T\n", frf)
	fmt.Println(frf())
	fmt.Printf("%T\n", frfr)
	fmt.Println(frfr())

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	even := filter(numbers, func(x int) bool {
		return x%2 == 0
	})
	filterOdd := func(o int) bool {
		return o%2 != 0
	}
	odd := filter(numbers, filterOdd)

	fmt.Println(even)
	fmt.Println(odd)
}

func refun() func() int {
	return func() int {
		return 1000
	}
}

func funrefun() func() int {
	return func() int {
		return 2000
	}
}

func filter(x []int, f func(y int) bool) []int {
	var r []int

	for _, n := range x {
		if f(n) {
			r = append(r, n)
		}
	}

	return r
}
