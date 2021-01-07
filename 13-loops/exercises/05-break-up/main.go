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
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
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
	i := num1
	for {
		if i%2 == 0 {
			fmt.Print(i)
			sum += int(i)

			if i != num2 {
				fmt.Print(" + ")
			}
		}
		i++
		if i > num2 {
			break
		}
	}
	fmt.Println(" = ", sum)
}
