package main

import "fmt"

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	resultInt := add(4, 4)
	resultString := add("Hello ", "World!")
	resultFloat := add(2.14, 3.14)

	fmt.Println("--- Generics ---")
	fmt.Printf("Result Interger: %v\nResult String: %v\nResult Float: %v\n", resultInt, resultString, resultFloat)
}