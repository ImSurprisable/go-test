package main

import (
	"errors"
	"fmt"
)

func main() {
	var myWords = "hello hi hello"
	printMe(myWords)

	var numA int = 10
	var numB int = 3
	var result, remainder, err = intDivision(numA, numB)

	switch { // switch statements dont require an initial argument & infer all of the 'break's after each case
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of %v divided by %v is %v.\n", numA, numB, result)
	default:
		fmt.Printf("The result of %v divided by %v is %v with remainder %v.\n", numA, numB, result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("The divison was exact.")
	case 1, 2:
		fmt.Println("The division was close.")
	default:
		fmt.Println("The division was not close.")
	}
}

func printMe(words string) {
	fmt.Println(words)
}

func intDivision(numerator int, denominator int) (int, int, error) { // return type goes AFTER parameters
	var err error // default value is 'nil', basically meaning there isnt an error yet
	if denominator == 0 {
		err = errors.New("cannot divide integers by zero") // set the error to what the issue is
		return 0, 0, err                                   // need to modify the return value to also include an error at the end
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
