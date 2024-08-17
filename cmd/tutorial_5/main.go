package main

import (
	"fmt"
	"strings"
)

func main() {
	// https://youtu.be/8uiZC0l4Ajw?si=3jEXBgq2RNMxvHeZ&t=1693
	var myString = "résumé"
	for i, v := range myString {
		fmt.Printf("%v: %v\n", i, v)
	}
	fmt.Println()
	var runeVersion = []rune(myString) // this is better if you want the character length instead of byte length from a string
	for i, v := range runeVersion {
		fmt.Printf("%v: %v\n", i, v)
	}

	var strSlice = []string{"h", "e", "l", "l", "o", " ", "world"}
	var strBuilder strings.Builder // much more efficient than constantly concatenating strings
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v\n", catStr)
}
