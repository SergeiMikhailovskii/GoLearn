package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var input int
	_, _ = fmt.Scanf("%d", &input)
	binaryStr := ""
	max := 0

	for input != 0 {
		binaryStr = strconv.Itoa(input%2) + binaryStr
		input /= 2
	}

	re := regexp.MustCompile("(1+)")

	variants := re.FindAll([]byte(binaryStr), -1)

	for _, element := range variants {
		if len(element) > max {
			max = len(element)
		}
	}
	fmt.Println(max)

}
