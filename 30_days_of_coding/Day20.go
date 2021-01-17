package main

import "fmt"

func main() {
	var amount int
	_, _ = fmt.Scanf("%d", &amount)
	count := 0

	arr := make([]int, amount)

	for i := range arr {
		_, _ = fmt.Scanf("%d", &arr[i])
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
				count++
			}
		}
	}

	fmt.Println("Array is sorted in", count, "swaps.")
	fmt.Println("First Element:", arr[0])
	fmt.Println("Last Element:", arr[len(arr)-1])
}
