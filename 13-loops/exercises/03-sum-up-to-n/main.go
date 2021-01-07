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
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
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

	for i := num1; i <= num2; i++ {

		fmt.Print(i)
		sum += int(i)

		if i != num2 {
			fmt.Print(" + ")
		}

	}
	fmt.Println(" = ", sum)
}
