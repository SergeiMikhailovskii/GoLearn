package main

import (
	"fmt"
	"math"
)

func main() {
	var amount int
	_, _ = fmt.Scanf("%d", &amount)

	arr := make([]int, amount)

	for i := range arr {
		_, _ = fmt.Scanf("%d", &arr[i])
	}

	for _, el := range arr {
		if el == 1 {
			fmt.Println("Not prime")
			continue
		}
		isPrime := false
		for i := 2; i <= int(math.Sqrt(float64(el))); i++ {
			if el%i == 0 {
				fmt.Println("Not prime")
				isPrime = true
				break
			}
		}
		if !isPrime {
			fmt.Println("Prime")
		}
	}

}
