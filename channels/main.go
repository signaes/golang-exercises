package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)

	c <- 30
	c <- 50

	for i := 0; i < 200; i++ {
		inc(&c, i)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func inc(c *chan int, i int) {
	go func() {
		v := <-*c
		fmt.Println("iteration:", i, "value:", v)
		v++
		*c <- v
	}()
}
