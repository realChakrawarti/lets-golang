package main

import "fmt"

type transformFn func(int) int

// Functions are first-class values in Go

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println("Doubled: ", doubled)
	fmt.Println("Tripled: ", tripled)
}

func transformNumbers(numberList *[]int, transform transformFn) []int {
	tNumbers := make([]int, len(*numberList))
	for idx, value := range *numberList {
		tNumbers[idx] = transform(value)
	}

	return tNumbers

}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}
