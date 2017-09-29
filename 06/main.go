package main

import (
	"fmt"
	"math/big"
)

var stop = big.NewInt(4000000)

func fibonacciEvensSum(limit *big.Int) *big.Int {
	first := big.NewInt(0)
	second := big.NewInt(1)
	total := big.NewInt(0)
	i := big.NewInt(0)

	for i.Cmp(limit) < 0 && second.Cmp(stop) == -1 {
		first.Add(first, second)
		first, second = second, first

		rem := big.NewInt(0)
		rem = rem.Rem(first, big.NewInt(2))

		if rem.Cmp(big.NewInt(0)) == 0 {
			total.Add(total, first)
		}

		i = i.Add(i, big.NewInt(1))
	}

	return total
}

func main() {
	ten := big.NewInt(10)
	limit := big.NewInt(4000000)

	fmt.Println(fibonacciEvensSum(ten))   // should return 44
	fmt.Println(fibonacciEvensSum(limit)) // should return 4613732
}

/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
