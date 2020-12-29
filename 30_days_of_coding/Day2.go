package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	localI, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	localF, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	localS := scanner.Text()

	fmt.Println(uint64(localI) + i)
	fmt.Printf("%.1f\n", localF+d)
	fmt.Println(s + localS)
	// Declare second integer, double, and String variables.

	// Read and save an integer, double, and String to your variables.

	// Print the sum of both integer variables on a new line.

	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

}
