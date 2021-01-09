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
)

func main() {

	digitPrinter(1337)
}

func digitPrinter(num int) {

	numbers := make([]int, 0)

	for num > 0 {
		digit := num % 10
		num /= 10
		numbers = append(numbers, digit)
	}

	digitCount := len(numbers)
	for i := 0; i <= 10; i++ {

		for j, _ := range numbers {

			v := numbers[digitCount-j-1]

			fmt.Print((*digits[v])[i] + " ")

		}

		fmt.Println("")
	}

}
