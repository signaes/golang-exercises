package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())

	printGoroutines()

	wg.Add(1)
	go a()
	printGoroutines()

	b()
	printGoroutines()

	wg.Wait()
}

func printGoroutines() {
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a", i)
	}

	wg.Done()
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b", i)
	}

}
