package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var amount int
	var key, value string
	users := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	_, _ = fmt.Scanf("%d", &amount)

	for i := 0; i < amount; i++ {
		_, _ = fmt.Scanf("%s %s", &key, &value)
		users[key] = value
	}

	for true {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if input == "" {
			return
		}
		if val, ok := users[input]; ok {
			fmt.Println(input + "=" + val)
		} else {
			fmt.Println("Not found")
		}
	}

}
