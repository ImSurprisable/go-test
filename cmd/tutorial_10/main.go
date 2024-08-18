package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))

	var float32Slice = []float32{1.5, 2.5, 3.5}
	fmt.Println(sumSlice(float32Slice))
	
	var float64Slice = []float64{}
	fmt.Println(sumSlice(float64Slice))

	fmt.Println(isEmpty(float64Slice))
	fmt.Println(isEmpty(intSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, value := range slice {
		sum += value
	}
	return sum
}


func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}