package main

import "fmt"

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
}

func transformNumbers(numberList *[]int, transform transformFn) []int {
	tNumbers := make([]int, len(*numberList))
	for idx, value := range *numberList {
		tNumbers[idx] = transform(value)
	}

	return tNumbers

}

// Return function as value
func getTransformer(multipler int) transformFn {
	if multipler == 2 {
		return double
	} else {
		return triple
	}
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}
