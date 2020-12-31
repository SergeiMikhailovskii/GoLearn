package main

import "fmt"

func main() {
	var input int
	_, _ = fmt.Scanf("%d", &input)
	if input%2 == 1 {
		fmt.Println("Weird")
	} else if input > 1 && input < 6 {
		fmt.Println("Not Weird")
	} else if input > 5 && input < 21 {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}
