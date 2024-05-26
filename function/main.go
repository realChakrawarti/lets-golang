package main

import (
	"fmt"
	"realChakrawarti/function/other"
)

type transformFn func(int) int

// Functions are first-class values in Go

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, getTransformer(2))
	tripled := transformNumbers(&numbers, getTransformer(3))

	// Anonymous function - One of use, not named, code colocation
	squared := transformNumbers(&numbers, func(num int) int {
		return num * num
	})

	fmt.Println("Doubled: ", doubled)
	fmt.Println("Tripled: ", tripled)
	fmt.Println("Squared: ", squared)

	fmt.Println("Factorial of 5 | Recursion: ", other.Factorial(5))
	sumNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Sum of numbers: ", other.Sumup(sumNumbers))
	fmt.Println("Sum of numbers | Variadic: ", other.VariadicSumup(1, 2, 3, 4, 5, 6, 7, 8, 9))

	// Slice into standalone parameter values

	fmt.Println("Sum of numbers | Split slice into params: ", other.VariadicSumup(1, sumNumbers...))
}

func transformNumbers(numberList *[]int, transform transformFn) []int {
	tNumbers := make([]int, len(*numberList))
	for idx, value := range *numberList {
		tNumbers[idx] = transform(value)
	}

	return tNumbers

}

// Factory pattern + Closure
func createTransformer(factor int) transformFn {
	return func(number int) int {
		return number * factor
	}
}

// Return function as value
func getTransformer(multipler int) transformFn {
	return createTransformer(multipler)
}
