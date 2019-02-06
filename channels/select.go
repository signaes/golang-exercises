package main

import "fmt"

var primes []int

func main() {
	even := make(chan int)
	odd := make(chan int)
	prime := make(chan int)
	quit := make(chan int)

	go send(1000, even, odd, prime, quit)
	receive(even, odd, prime, quit)

	fmt.Println("bye")
	fmt.Println(primes)
}

func isPrime(x int) bool {
	var b bool
	l := []int{2, 3, 5, 7}

	for _, v := range l {
		b = x == v || x%v != 0

		if !b {
			return b
		}
	}

	return b
}

func send(times int, e, o, p, q chan<- int) {
	for i := 0; i < times; i++ {
		if isPrime(i) {
			p <- i
			primes = append(primes, i)
		}

		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	close(e)
	close(o)

	q <- 0
}

func receive(e, o, p, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("even:", v)
		case v := <-o:
			fmt.Println("odd:", v)
		case v := <-p:
			fmt.Println("prime:", v)
		case v := <-q:
			fmt.Println("quit:", v)
			return
		}
	}
}
