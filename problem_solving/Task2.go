package main

import "fmt"

func main() {
	var amount int
	_, _ = fmt.Scanf("%d", &amount)

	sum := 0

	for i := 0; i < amount; i++ {
		value := 0
		_, _ = fmt.Scanf("%d", &value)
		sum += value
	}

	fmt.Println(sum)
}
