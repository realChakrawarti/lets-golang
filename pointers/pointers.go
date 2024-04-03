package main

import "fmt"

func main() {

	age := 32
	agePointer := &age

	fmt.Println("Actual age: ", age)
	fmt.Println("Actual age in memory: ", agePointer)
	fmt.Println("De-referencing actual age Pointer: ", *agePointer)

	adultAge := getAdultAge(age)
	fmt.Println("Adult age: ", adultAge)

	adultAgePointer := getAdultAgePointer(&age)
	fmt.Println("Adult age using pointer: ", adultAgePointer)

	getAdultAgeMutate(agePointer)
	fmt.Println("Actual age after mutating: ", age)
}

func getAdultAge(age int) int {
	return age - 18
}

func getAdultAgePointer(age *int) int {
	return *age - 18
}

func getAdultAgeMutate(age *int) {
	*age = *age - 18
}
