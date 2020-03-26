package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func getMultiValue() (int, int) {
	return 3, 5
}

func main() {
	num1, num2 := getMultiValue()

	fmt.Println(num1, num2)
	fmt.Println(greeting("Agung Sepruloh"))
	fmt.Println(getSum(num1, num2))
}
