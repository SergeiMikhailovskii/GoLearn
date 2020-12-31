package main

import (
	"fmt"
)

func main() {
	var size int
	_, _ = fmt.Scanf("%d", &size)

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		_, _ = fmt.Scanf("%d", &arr[i])
	}

	steps := 0

	for i := 0; i < size; i++ {
		if i == size-1 {
			break
		} else if i == size-2 {
			steps++
		} else {
			if arr[i+2] == 0 {
				i++
				steps++
			} else {
				steps++
			}
		}
	}

	fmt.Println(steps)
}
