package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 0

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()

			v := counter
			runtime.Gosched()
			fmt.Println("Goroutines", runtime.NumGoroutine())
			v++
			counter = v

			mu.Unlock()
			wg.Done()
		}()
	}

	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("counter", counter)

	defer printCounter()
}

func printCounter() {
	fmt.Println("counter", counter)
}
