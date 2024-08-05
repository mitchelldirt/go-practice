package main

import (
	"errors"
	"fmt"
)

func functionsControls(){
	// printMe("Hello")
	numerator := 10
	denominator := 6

	var result, remainder, err = intDivision(numerator, denominator)

		if err != nil {
		fmt.Print(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v.", result, remainder)
	}

	switch {
	case err!=nil:
		fmt.Println(err.Error())
	case remainder==0:
		fmt.Printf("The result of the integer division is %v\n", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v.\n", result, remainder)
	}

	switch remainder{
	case 0:
		fmt.Println("The remainder was 0.")
	case 1,2:
		fmt.Println("The remainder was close to 0.")
	default:
		fmt.Println("The remainder was not close to 0.")
	}
}

// func printMe(input string) {
// 	fmt.Println(input)
// }

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator==0{
		err = errors.New("CANNOT DIVIDE BY ZERO")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
