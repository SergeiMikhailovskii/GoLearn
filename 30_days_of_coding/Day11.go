package main

import (
	"fmt"
	"math"
)

func main() {
	clockLength, size, maxSum := 3, 6, math.MinInt64
	arr := make([][]int, size)

	for i := range arr {
		arr[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			_, _ = fmt.Scanf("%d", &arr[i][j])
		}
	}

	for i := 0; i < size-clockLength+1; i++ {
		for j := 0; j < size-clockLength+1; j++ {
			var subMatrix = make([]int, clockLength*clockLength)
			pos := 0
			for k := i; k < i+clockLength; k++ {
				for l := j; l < j+clockLength; l++ {
					subMatrix[pos] = arr[k][l]
					pos++
				}
			}

			localSum := subMatrix[0] + subMatrix[1] + subMatrix[2] + subMatrix[4] + subMatrix[6] + subMatrix[7] + subMatrix[8]

			if localSum > maxSum {
				maxSum = localSum
			}

		}
	}

	fmt.Println(maxSum)
}
