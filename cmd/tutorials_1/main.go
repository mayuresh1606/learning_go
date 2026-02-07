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

	// const numerator int = 198
	// const denominator int = 0
	// var result, remainder, err = divide(numerator, denominator)

	// switch {
	// case err != nil:
	// 	fmt.Printf(err.Error())
	// case remainder == 0:
	// 	fmt.Printf("Result of division is: %v", result)
	// default:
	// 	fmt.Printf("Result of division: %v, and remainder is: %v", result, remainder)
	// }

	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("Result of division is: %v", result)
	// } else {
	// 	fmt.Printf("Result of division: %v, and remainder is: %v", result, remainder)
	// }

	intArr := [5]int32{1, 2, 4, 5, 6}
	fmt.Printf(`Array: %v`, intArr)

	var intSlice []int32 = []int32{10, 22, 32, 45}
	fmt.Println("")
	fmt.Printf(`Slice: %v, length of slice %v, capacity %v`, intSlice, len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 5)
	// fmt.Println("")
	fmt.Printf("\nAfter append Slice: %v, length of slice %v, capacity %v \n", intSlice, len(intSlice), cap(intSlice))

	intSlicePredefined := make([]int32, 3, 5)
	fmt.Printf(`Predefined Slice: %v, length of slice %v, capacity %v`, intSlicePredefined, len(intSlicePredefined), cap(intSlicePredefined))

	var intMap map[string]int32 = map[string]int32{
		"Mayuresh": 23,
		"Lavanya":  22,
	}

	fmt.Printf("\nMy map: %v", intMap)
	fmt.Printf("\nMayuresh Age: %v", intMap["Mayuresh"])

	var mayureshAge, isPresent = intMap["Mayuresh"]
	if isPresent {
		fmt.Println("\nMayuresh Age: ", mayureshAge)
	} else {
		fmt.Println("\nNot present")
	}

	fmt.Printf("\nSomething which is not present: %v", intMap["Jason"])

	// will not run in order as maps are not ordered
	for name, age := range intMap {
		fmt.Printf("\n Name: %v, Age: %v", name, age)
	}

	for singleSlice := range intSlice {
		fmt.Printf("\nSlices: %v", intSlice[singleSlice])
	}

	for i, value := range intSlice {
		fmt.Printf("\nIndex and value of slice %v, %v", i, value)
	}

	// simple way to write while loop in go
	var whileCount int32 = 0
	for whileCount <= 5 {
		fmt.Printf("\n While loop with count: %v", whileCount)
		whileCount++
	}

	// simple way to write while loop in go 2
	for whileCount2 := 0; whileCount2 <= 5; whileCount2++ {
		fmt.Printf("\n While loop 2 with count: %v", whileCount2)
	}
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
