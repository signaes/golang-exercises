package main

import "fmt"

func main() {
	c := make(chan int)
	var x []int

	go send(c, 100)
	receive(c, &x)

	fmt.Println(x)
}

func send(c chan<- int, times int) {
	for i := 0; i < times; i++ {
		c <- i
	}

	close(c)
}

func receive(c <-chan int, x *[]int) {
	for v := range c {
		*x = append(*x, v)
	}
}
