package main

import (
	"fmt"
)

func main() {
	var amount int
	_, _ = fmt.Scanf("%d", &amount)

	pairs := 0
	set := make(map[int]bool)
	for i := 0; i < amount; i++ {
		var number int
		_, _ = fmt.Scanf("%d", &number)

		if set[number] {
			set[number] = false
			pairs++
		} else {
			set[number] = true
		}
	}
	fmt.Println(pairs)
}
