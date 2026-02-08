package main

import (
	"fmt"
	"unicode/utf8"
)

func variablesBasicDataTypes() {
	fmt.Println("Hello World!")
	var myInt8 int8 = 12
	var myInt16 int16 = 6
	var myString string = "hello"
	var myFloat32 float32 = 32.70
	var myFloat64 float64 = 64.70
	myNewString := `New String`
	fmt.Println(myInt8)
	fmt.Println(myInt8/int8(myInt16), "DIVIDE INT OPERATION")
	fmt.Println(myFloat32/float32(myFloat64), "DIVIDE FLOAT OPERATION")
	fmt.Println(int(myFloat32)%int(myFloat64), "REMAINDER INT OPERATION")
	fmt.Println(myString)
	fmt.Println(myNewString)

	fmt.Println(utf8.RuneCountInString(myString), "RUNE COUNT OF LENGTH")

	var myRune rune = 'A'
	var aString string = "A"
	fmt.Println(aString, "NORMAL A STRING")
	fmt.Println(myRune, "RUNE VALUE")

	var myBoolean bool = false
	fmt.Println(myBoolean, "BOOLEAN VALUE")

	var var1, var2, var3 int = 1, 2, 3
	fmt.Println(var1, var2, var3, "ALL VARIABLES")

	const pi float32 = 3.14
	fmt.Println(pi, "PI VALUE")
}
