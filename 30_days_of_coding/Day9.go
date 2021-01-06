package main

import "fmt"

func main() {
	var num int
	_, _ = fmt.Scanf("%d", &num)
	fmt.Println(factorial(num))
}

func factorial(num int) int {
	if num == 1 || num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
