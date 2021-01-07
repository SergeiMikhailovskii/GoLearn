package main

import (
	"fmt"
	"math"
)

func main() {
	var size int
	lDiag, rDiag := 0, 0
	_, _ = fmt.Scanf("%d", &size)
	arr := make([][]int, size)

	for i := range arr {
		arr[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			_, _ = fmt.Scanf("%d", &arr[i][j])
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				lDiag += arr[i][j]
			}
			if j == size-i-1 {
				rDiag += arr[i][j]
			}
		}
	}

	fmt.Println(int(math.Abs(float64(lDiag - rDiag))))
}
