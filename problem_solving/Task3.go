package main

import "fmt"

func main() {
	var a [3]int
	var b [3]int
	_, _ = fmt.Scanf("%d %d %d", &a[0], &a[1], &a[2])
	_, _ = fmt.Scanf("%d %d %d", &b[0], &b[1], &b[2])

	aScore := 0
	bScore := 0

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			aScore++
		} else if a[i] < b[i] {
			bScore++
		}
	}

	fmt.Println(aScore, bScore)
}
