package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	str, _, _ := reader.ReadLine()
	dateReturned := NewDate(strings.TrimRight(string(str), "\r\n"))

	str, _, _ = reader.ReadLine()
	dateDue := NewDate(strings.TrimRight(string(str), "\r\n"))

	fmt.Println(checkDates(*dateReturned, *dateDue))
}

func NewDate(dateStr string) *time.Time {
	arr := strings.Split(dateStr, " ")
	days, _ := strconv.Atoi(arr[0])
	month, _ := strconv.Atoi(arr[1])
	years, _ := strconv.Atoi(arr[2])

	arr[0] = fmt.Sprintf("%02d", days)
	arr[1] = fmt.Sprintf("%02d", month)
	arr[2] = fmt.Sprintf("%04d", years)

	dateStr = arr[0] + "-" + arr[1] + "-" + arr[2]

	timeT, _ := time.Parse("2-1-2006", dateStr)
	return &timeT
}

func checkDates(dateReturned time.Time, dateDue time.Time) int {
	if dateReturned.Before(dateDue) {
		return 0
	} else if dateReturned.Month() == dateDue.Month() && dateReturned.Year() == dateDue.Year() {
		days := time.Time{}.Add(dateReturned.Sub(dateDue)).Day() - 1
		return 15 * days
	} else if dateReturned.Year() == dateDue.Year() {
		return 500 * (int(time.Time{}.Add(dateReturned.Sub(dateDue)).Month()) - 1)
	} else {
		return 10_000
	}
}
