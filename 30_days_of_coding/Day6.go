package main

import "fmt"

func main() {
	var amount int
	_, _ = fmt.Scanf("%d", &amount)

	arr := make([]string, amount)
	for i := 0; i < amount; i++ {
		_, _ = fmt.Scanf("%s", &arr[i])
	}

	for i := 0; i < len(arr); i++ {
		evenStr := ""
		oddStr := ""

		for j := 0; j < len(arr[i]); j++ {
			if j%2 == 0 {
				evenStr += string(arr[i][j])
			} else {
				oddStr += string(arr[i][j])
			}
		}

		fmt.Println(evenStr, oddStr)
	}

}
