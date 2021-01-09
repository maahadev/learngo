// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

var (
	phone_number         map[string]uint64
	product_availability map[int]bool
	phone_numbers        map[string]*int64
	shopping_basket      map[string]struct {
		quantity  int
		item_name string
	}
)

func main() {
	fmt.Println(phone_number)
	fmt.Println(product_availability)
	fmt.Println(phone_numbers)
	fmt.Println(shopping_basket)
}
