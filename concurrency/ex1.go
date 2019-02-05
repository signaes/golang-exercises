package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("a")
		wg.Done()
	}()

	go func() {
		fmt.Println("b")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Bye")
}
