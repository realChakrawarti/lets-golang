package other

func Factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * Factorial(number-1)
}

func Sumup(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}

func VariadicSumup(startingNumber int, numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum + startingNumber
}
