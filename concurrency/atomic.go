package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64 = 0

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("value", atomic.LoadInt64(&counter))
			runtime.Gosched()
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
