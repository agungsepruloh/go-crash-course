package main

import (
	"fmt"
	"strconv" // string converter
)

// Person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m" {
		return
	}
	// change the lastname if the gender is f and married
	person.lastName = spouseLastName
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 23}

	// Alternative
	person2 := Person{"Samantha", "Smith", "Boston", "f", 23}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person2)
	fmt.Println()

	// hasBirthday
	person1.hasBirthday()

	// is getMarried and female
	person1.getMarried("Williams")

	// greeting
	fmt.Println(person1.greet())
}
