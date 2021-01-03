package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	for i := 1; i <= 10; i++ {
		fmt.Println(a, "x", i, "=", a*i)
	}
}
