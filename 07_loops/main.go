package main

import "fmt"

func main() {

	// Long Method
	i := 1
	for i <= 5 {
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()

	// Short Method
	for j := 1; j <= 10; j++ {
		fmt.Print(j*10, " ")
	}
}
