package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println("Hello World!")
	// var myInt8 int8 = 12
	// var myInt16 int16 = 6
	// var myString string = "hello"
	// var myFloat32 float32 = 32.70
	// var myFloat64 float64 = 64.70
	// myNewString := `New String`
	// fmt.Println(myInt8)
	// fmt.Println(myInt8/int8(myInt16), "DIVIDE INT OPERATION")
	// fmt.Println(myFloat32/float32(myFloat64), "DIVIDE FLOAT OPERATION")
	// fmt.Println(int(myFloat32)%int(myFloat64), "REMAINDER INT OPERATION")
	// fmt.Println(myString)
	// fmt.Println(myNewString)

	// fmt.Println(utf8.RuneCountInString(myString), "RUNE COUNT OF LENGTH")

	// var myRune rune = 'A'
	// var aString string = "A"
	// fmt.Println(aString, "NORMAL A STRING")
	// fmt.Println(myRune, "RUNE VALUE")

	// var myBoolean bool = false
	// fmt.Println(myBoolean, "BOOLEAN VALUE")

	// var var1, var2, var3 int = 1, 2, 3
	// fmt.Println(var1, var2, var3, "ALL VARIABLES")

	// const pi float32 = 3.14
	// fmt.Println(pi, "PI VALUE")
	const numerator int = 198
	const denominator int = 0
	var result, remainder, err = divide(numerator, denominator)
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("Result of division is: %v", result)
	default:
		fmt.Printf("Result of division: %v, and remainder is: %v", result, remainder)
	}
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("Result of division is: %v", result)
	// } else {
	// 	fmt.Printf("Result of division: %v, and remainder is: %v", result, remainder)
	// }
}

func divide(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("Denominator cannot be zero")
		return 0, 0, err
	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder, err
}
