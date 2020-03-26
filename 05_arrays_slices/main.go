package main

import "fmt"

func main() {
	// Declare an array
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign values
	fruitArr := [2]string{"Apple", "Orange"}
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1]) // get a specific value use index
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice)) // get length of the array
	fmt.Println(fruitSlice[1:3]) // get in range
}
