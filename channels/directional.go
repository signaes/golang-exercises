package main

import "fmt"

func main() {
	b := make(chan int, 2)
	c := make(chan<- int, 1)
	r := make(<-chan int, 1)

	c <- 100

	var v int

	go func() {
		v = <-r
	}()

	fmt.Printf("%T\n", (chan<- int)(b))
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", r)
	fmt.Println(v)

	go send(b, 1000)

	var x int

	receive(b, &x)

	fmt.Println("x:", x)
}

func send(c chan<- int, i int) {
	c <- i
}

func receive(c <-chan int, x *int) {
	*x = <-c
}
