package main

import (
	"fmt"
	"regexp"
)

func main() {
	var base string
	var length uint64
	_, _ = fmt.Scanf("%s", &base)
	_, _ = fmt.Scanf("%d", &length)

	regex := regexp.MustCompile("a")
	aInBaseStrAmount := len(regex.FindAllStringIndex(base, -1))

	mult := length / uint64(len(base))

	aInStr := uint64(aInBaseStrAmount) * mult

	left := length - mult*uint64(len(base))

	leftStr := base[0:left]

	aInStr += uint64(len(regex.FindAllStringIndex(leftStr, -1)))
	fmt.Println(aInStr)
}
