package main

import "fmt"

func main() {
	var amount int
	_, _ = fmt.Scanf("%d", &amount)

	positives, negatives, zeros := 0, 0, 0

	for i := 0; i < amount; i++ {
		var input int
		_, _ = fmt.Scanf("%d", &input)
		if input > 0 {
			positives++
		} else if input < 0 {
			negatives++
		} else {
			zeros++
		}
	}

	fmt.Println(fmt.Sprintf("%.6f", float64(positives)/float64(amount)))
	fmt.Println(fmt.Sprintf("%.6f", float64(negatives)/float64(amount)))
	fmt.Println(fmt.Sprintf("%.6f", float64(zeros)/float64(amount)))

}
