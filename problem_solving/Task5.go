package main

import "fmt"

func main() {
	var amount int
	var input string

	position := 0
	valleys := 0

	_, _ = fmt.Scanf("%d", &amount)
	_, _ = fmt.Scanf("%s", &input)

	arr := make([]string, amount)

	for i, symbolInt := range input {
		arr[i] = string(symbolInt)
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == "U" {
			position++
		} else {
			position--
		}
		if position == 0 && arr[i] == "U" {
			valleys++
		}

	}

	fmt.Println(valleys)

}
