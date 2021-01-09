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
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

var students = map[string][]string{
	"gryffinder": {"weasley", "hagrid", "dumbledore", "lupin"},
	"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
	"ravenclaw":  {"bagnold", "wildsmith", "montmorency", "horace"},
	"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
	"bobo":       {"wizardry", "unwanted"},
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide a house name.")
		return
	}

	item, found := students[os.Args[1]]
	if !found {
		fmt.Println("Not Found")
		return
	}
	for _, name := range item {

		fmt.Printf("+ %s\n", name)

	}
}
