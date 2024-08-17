package main

import "fmt"

func main() {
	var p *int32 = new(int32) // new pointer with an address to a new int32 value
	var i int32 = 18
	fmt.Printf("The value p points to is %v\n", *p)

	*p = 10 // the int32 at the pointer's address has its value changed to 10
	fmt.Printf("The value p points to is %v\n", *p)

	p = &i // the pointer now points to the memory address of i
	fmt.Printf("The value p points to is %v\n", *p)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Println(thing1)
	square(&thing1)
	fmt.Println(thing1)

}

func square(thing2 *[5]float64) {
	for i := range thing2 {
		thing2[i] *= thing2[i]
	}
}