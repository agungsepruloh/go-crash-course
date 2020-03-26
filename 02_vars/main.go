package main

import "fmt"

func main() {
	// MAIN DATA TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr (unsign - no negative number)
	// rune - alias for int32
	// byte - alias for uint8
	// float32 float64
	// complex64 complex128

	// Using var keyword
	var age int = 19
	const isCool = true

	// Shorthand
	size := 3.2
	name := "Agung Sepruloh"
	email, address := "agungsepruloh1996@gmail.com", "27, Kolonel Masturi"

	fmt.Println(age, isCool)
	fmt.Println(name)
	fmt.Println(email, address)
	fmt.Printf("%T\n", name) // print the data type
	fmt.Printf("%T\n", size) // print the data type
}
