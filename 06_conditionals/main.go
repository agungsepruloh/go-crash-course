package main

import "fmt"

func main() {
	// For AND use &&
	// For OR use ||

	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d \n", x, y)
	} else if x > y {
		fmt.Printf("%d is greater than %d \n", x, y)
	} else {
		fmt.Printf("%d and %d is equal \n", x, y)
	}

	color := "Blue"

	switch color {
	case "Red":
		fmt.Println("Color is Red")
	case "Blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is Not Blue and Red")
	}
}
