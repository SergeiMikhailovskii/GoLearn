package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	hiddenNumber := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Println("Enter the number")
		inputStr, _ := reader.ReadString('\n')
		inputInt, _ := strconv.Atoi(strings.TrimSpace(inputStr))
		if inputInt == hiddenNumber {
			fmt.Println("You are right")
			fmt.Println("You win!")
			return
		} else if inputInt > hiddenNumber {
			fmt.Println("Hidden number is less")
		} else {
			fmt.Println("Hidden number is more")
		}
	}
	fmt.Println("You lose. Hidden number was", hiddenNumber)

}
