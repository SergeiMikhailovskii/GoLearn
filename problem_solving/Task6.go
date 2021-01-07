package main

import "fmt"

func main() {
	var amount int
	var sum uint64
	_, _ = fmt.Scanf("%d", &amount)
	arr := make([]uint64, amount)

	for i := 0; i < amount; i++ {
		_, _ = fmt.Scanf("%d", &arr[i])
	}

	for _, value := range arr {
		sum += value
	}

	fmt.Println(sum)
}
