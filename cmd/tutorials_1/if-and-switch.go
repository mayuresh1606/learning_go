package main

import (
	"errors"
	"fmt"
)

func ifSwitchAndError() {
	const numerator int = 198
	const denominator int = 10
	var result, remainder, err = divide(numerator, denominator)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("\nResult of division is: %v", result)
	default:
		fmt.Printf("\nResult of division: %v, and remainder is: %v", result, remainder)
	}

	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("\nResult of division is: %v", result)
	} else {
		fmt.Printf("\nResult of division: %v, and remainder is: %v", result, remainder)
	}
}

func divide(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("\nDenominator cannot be zero\n")
		return 0, 0, err
	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder, err
}
