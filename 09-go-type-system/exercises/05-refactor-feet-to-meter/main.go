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
// EXERCISE: Refactor Feet to Meter
//
//  Remember the feet to meters program?
//  Now, it's time to refactor it.
//  Define your own Feet and Meters types.
//
//  Follow the steps inside the code.
// ---------------------------------------------------------
type Feet float64
type Meters float64

func main() {
	var (
		feet   Feet
		meters Meters
	)

	feet, _ = Feet(strconv.ParseFloat(os.Args[1], 64))

	meters = Meters(feet * 0.3048)

	// ----------------------------
	// 8. UNCOMMENT THE CODE BELOW
	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
