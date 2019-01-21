package main

import "fmt"

func main() {
	s := "Is it a pointer?"
	ps := &s

	fmt.Printf("%T %v\n", s, s)
	fmt.Printf("%T %v\n", ps, ps)

	*ps = "Yes"

	fmt.Printf("%T %v\n", s, s)

	assign(ps, "Reassign")

	fmt.Printf("%T %v\n", s, s)

	fmt.Printf("%T %v\n", s, s)
}

func assign(s *string, value string) {
	*s = value
}
