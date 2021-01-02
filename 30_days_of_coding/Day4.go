package main

import "fmt"

var amount int

func main() {
	_, _ = fmt.Scanf("%d", &amount)

	arr := make([]int, amount)
	for i := 0; i < amount; i++ {
		_, _ = fmt.Scanf("%d", &arr[i])
	}

	for _, personAge := range arr {
		person := new(person)
		person = person.NewPerson(personAge)
		person.AmIOld()

		for i := 0; i < 3; i++ {
			person = person.YearPasses()
		}

		person.AmIOld()
		fmt.Println()
	}

}

type person struct {
	age int
}

func (receiver person) NewPerson(age int) *person {
	receiver.age = validateAge(age)
	return &receiver
}

func (receiver person) YearPasses() *person {
	receiver.age++
	return &receiver
}

func (receiver person) AmIOld() {
	if receiver.age < 13 {
		fmt.Println("You are young.")
	} else if receiver.age > 12 && receiver.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}
}

func validateAge(age int) int {
	if age < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		return 0
	}
	return age
}
