package main

import "fmt"

func arraysSlicesMapsLoops() {
	intArr := [5]int32{1, 2, 4, 5, 6}
	fmt.Printf("\n\nArray: %v", intArr)

	var intSlice []int32 = []int32{10, 22, 32, 45}
	fmt.Println("")
	fmt.Printf(`Slice: %v, length of slice %v, capacity %v`, intSlice, len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 5)
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
