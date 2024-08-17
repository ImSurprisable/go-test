package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{0, 5, 2}
	intArr[1] = 1
	fmt.Println(intArr[0])   // element 0
	fmt.Println(intArr[1:3]) // elements 1-2 in their own array

	// printing the memory addresses to each element
	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6} // create a slice (array w/ more functionality) by omitting the length value in the brackets []
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) // creates a new array with more capacity if necessary (goes from [4, 5, 6] to [4, 5, 6, 7 *, *])
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // append other slices/arrays with ... at the end
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8) // another way to create a slice; first parameter is length, second is capacity. capacity defaults to the length if not specified
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice3), cap(intSlice3))

	var myMap = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap)
	var age, ok = myMap["Adam"]                       // maps return the key as well as an "okay" boolean so you know if the key actually exists or not
	fmt.Printf("%v years old, this is %v\n", age, ok) // this will be fine since "Adam" is an actual key in the map
	age, ok = myMap["Jason"]
	fmt.Printf("%v years old, this is %v\n\n", age, ok) // this will not be fine because "Jason" isn't a key in the map, so it returns default int value 0.

	// some automatic foreach loops using 'range' keyword
	for name, age := range myMap {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// typical for loop structure
	for i := 0; i < 5; i++ {
		fmt.Printf("For %v\n", i)
	}

	// go's version of a while loop
	var i int = 0
	for i <= 3 {
		fmt.Printf("While %v\n", i)
		i++
	}
}
