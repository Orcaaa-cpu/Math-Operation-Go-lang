package main

import "fmt"

func main() {
	a, b := 30, 6
	fmt.Printf("The addition of %v + %v = %v \n", a, b, addition(a, b))
	fmt.Printf("The addition of %v - %v = %v \n", a, b, subtraction(a, b))
	fmt.Printf("The addition of %v * %v = %v \n", a, b, multiplication(a, b))
	fmt.Printf("The addition of %v / %v = %v \n", a, b, division(a, b))
}

func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) int {
	return a / b
}
