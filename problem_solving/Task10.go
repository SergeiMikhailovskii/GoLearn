package main

import "fmt"

func main() {
	var size, rotationLeft int
	_, _ = fmt.Scanf("%d", &size)
	_, _ = fmt.Scanf("%d", &rotationLeft)

	arr := make([]int, size)
	newArr := make([]int, size)
	for i := 0; i < size; i++ {
		_, _ = fmt.Scanf("%d", &arr[i])
	}

	for i := rotationLeft; i < size; i++ {
		newArr[i-rotationLeft] = arr[i]
	}
	for i := 0; i < rotationLeft; i++ {
		newArr[size-rotationLeft+i] = arr[i]
	}

	for _, el := range newArr {
		fmt.Print(el, " ")
	}
	fmt.Println()
}
