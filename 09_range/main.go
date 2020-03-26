package main

import "fmt"

func main() {
	ids := []int{33, 26, 36, 78, 90}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}
	fmt.Println()

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}
	fmt.Println()

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)
	fmt.Println()

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s \n", k, v)
	}
	fmt.Println()

	for k := range emails {
		fmt.Println("Name:", k)
	}
	fmt.Println()

	for _, v := range emails {
		fmt.Println("Email:", v)
	}
}
