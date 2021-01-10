// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	for {
		screen.Clear()
		screen.MoveTopLeft()
		hour, minute, second := getCurrentTime()
		hour_arr, minute_arr, second_arr := digitToTwoNumberArray(hour), digitToTwoNumberArray(minute), digitToTwoNumberArray(second)
		clockPrinter(hour_arr, minute_arr, second_arr)
		time.Sleep(1 * time.Second)
	}

}

func clockPrinter(hour, minute, second []int) {

	fmt.Print("\n\n")
	for i := 0; i <= 10; i++ {

		fmt.Print((*digits[hour[0]])[i])
		fmt.Print(" ")
		fmt.Print((*digits[hour[1]])[i])
		fmt.Print(" ")

		if second[1]%2 == 0 {
			fmt.Print((seperator[i]))
		} else {
			fmt.Print(strings.Repeat(" ", 13))
		}

		fmt.Print((*digits[minute[0]])[i])
		fmt.Print(" ")
		fmt.Print((*digits[minute[1]])[i])
		fmt.Print(" ")

		if second[1]%2 == 0 {
			fmt.Print((seperator[i]))
		} else {
			fmt.Print(strings.Repeat(" ", 13))

		}
		fmt.Print((*digits[second[0]])[i])
		fmt.Print(" ")
		fmt.Print((*digits[second[1]])[i])
		fmt.Print(" ")

		fmt.Println("")
	}

}

func digitToTwoNumberArray(num int) []int {

	numbers := make([]int, 0, 2)
	if num == 0 {
		return []int{0, 0}
	}

	for num > 0 {
		digit := num % 10
		num /= 10
		numbers = append(numbers, digit)
	}

	if len(numbers) == 1 {
		return []int{0, numbers[0]}
	}
	return []int{numbers[1], numbers[0]}

}

func getCurrentTime() (int, int, int) {

	currentTime := time.Now()
	return currentTime.Hour(), currentTime.Minute(), currentTime.Second()

}
