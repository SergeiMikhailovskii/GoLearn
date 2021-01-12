package main

import "fmt"

func main() {
	var amount int
	_, _ = fmt.Scanf("%d", &amount)

	arr := make([]int, amount)

	for i := range arr {
		_, _ = fmt.Scanf("%d", &arr[i])
	}

	swaps := 0
	for i, el := range arr {
		if i != el-1 {
			for j := i + 1; j < len(arr); j++ {
				if arr[j] == i+1 {
					arr[i] = arr[j]
					arr[j] = el
					swaps++
					break
				}
			}
		}
	}

	fmt.Println(swaps)

}
