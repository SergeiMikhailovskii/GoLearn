package main

import "fmt"

func main() {
	var amount int
	_, _ = fmt.Scanf("%d", &amount)

	arr := make([]int, amount)

	for i := 0; i < amount; i++ {
		_, _ = fmt.Scanf("%d", &arr[i])
	}

	for i := amount - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

}
