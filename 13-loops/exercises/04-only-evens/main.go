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
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {

	args := os.Args
	sum := 0
	if len(args) != 3 {
		fmt.Println("Usage main [min] [max]")
		return
	}
	num1, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		fmt.Println("Invalid [min] number")
		return
	}

	num2, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		fmt.Println("Invalid [min] number")
		return
	}

	if num1 > num2 {
		fmt.Println("Make sure [min] < [max]")
		return
	}

	if num2%2 != 0 {
		num2 -= 1
	}
	for i := num1; i <= num2; i++ {
		if i%2 == 0 {
			fmt.Print(i)
			sum += int(i)

			if i != num2 {
				fmt.Print(" + ")
			}
		}
	}
	fmt.Println(" = ", sum)
}
