package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
}

// creating a method for the "gasEngine" struct
func (e gasEngine) milesLeft() uint16 {
	return uint16(e.gallons) * uint16(e.mpg)
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e electricEngine) milesLeft() uint16 {
	return uint16(e.mpkwh) * uint16(e.kwh)
}


type engine interface {
	milesLeft() uint16
}


func canMakeIt(e engine, miles uint16) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it! :D")
	} else {
		fmt.Println("Sorry, you can't make that. :(")
	}
}


func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 200)

	var myElectriceEngine electricEngine = electricEngine{10, 25}
	canMakeIt(myElectriceEngine, 400)

	var myOtherEngine = struct {
		mpg uint8
		gallons uint8
	} {10, 30}
	fmt.Println(myOtherEngine.mpg, myOtherEngine.gallons)
}
