package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var numInt int = 25
	numInt += 5
	fmt.Println(numInt)

	var floatNum float64 = 3.1415926535
	floatNum *= 2
	fmt.Println(floatNum)

	var intA int = 3
	var intB int = 2
	var intC int = intA / intB // remains as an int throughout the operation, so it floors the result
	fmt.Println(intC)

	var boringString string = "Im such a boring" + " string :("
	fmt.Println(boringString)

	// backquote lets you format strings in the code itself
	var funString string = `
	hi this string
	is very funky but
	its pretty cool
	ok bye
	`
	fmt.Println(funString)

	fmt.Println(utf8.RuneCountInString(boringString)) // accurate counting of strings outside of normal ASCII characters

	var myRune rune = 'a' // runes are characters? weird name but sure
	fmt.Println(myRune)   // not sure why this prints out 97, i dont think thats the ascii value of 'a'?? ill probably find out why this happens later

	var intNumZ int
	fmt.Println(intNumZ) // default values (in the case of an int, its 0)

	// inferred types are possible basically
	// var var1 = "inferred type"
	// var2 := "also inferred"
	// var var3, var4, var5 = 2, 's', true
	// var6, var7 := "lol", 1.234

	const name string = "IM_SURPRISABLE"
	fmt.Println(name)

}
