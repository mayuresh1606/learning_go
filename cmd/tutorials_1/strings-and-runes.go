package main

import (
	"fmt"
	"strings"
)

func stringsAndRunes() {
	// var myString string = "résumé"

	// // you can see the output as length: 8 since go doesn't count the length of the stirng but bytes that it represents.
	// fmt.Println(len(myString))
	// // whole string has type string
	// fmt.Printf("Type of string: %T\n", myString)
	// for i, v := range myString {
	// 	// note that single value of that string has value of int32 since it represents the bits value
	// 	fmt.Printf("Index: %v, Value: %v, Type of string: %T\n", i, v, v)
	// 	// println(i, v)
	// 	// println("Value, ", v)
	// }

	// rune behaviour
	var myRune = []rune("résumé")
	// we can also declare runes with single quotes like so:
	// var myRune2 = 'a'

	// indexed value of rune
	fmt.Printf("\nIndexed value %v", myRune[1])
	// you can see the output as length: 8 since go doesn't count the length of the stirng but bytes that it represents.
	fmt.Println("\nLength of rune: ", len(myRune))
	// whole rune has type int32
	fmt.Printf("Type of rune: %T\n", myRune)
	for i, v := range myRune {
		// note that single value of that string has value of int32 since it represents the bits value
		fmt.Printf("Index: %v, Value: %v, Type of string: %T\n", i, v, v)
		// println(i, v)
		// println("Value, ", v)
	}

	var myString = []string{"M", "a", "y", "u"}
	catString := ""
	for i := range myString {
		// fmt.Printf("\n Index: %v, Value: %v", i, v)
		catString += myString[i]
	}

	fmt.Printf("\nConcatinated String: %v", catString)

	// Since strings are immutable in go += creates new string each time while concatenating which is inefficient we can use string builder for that
	var strBuilder strings.Builder
	for i := range myString {
		// an array is allocated internally and values are appended in that array
		strBuilder.WriteString(myString[i])
	}

	// only when String() method is called then that string is created instead of our previous approach
	catString2 := strBuilder.String()

	fmt.Printf("\nString Builder concatinated string: %v", catString2)
}
