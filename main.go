package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
}
