package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output(label string) {
	fmt.Printf("%v: %v", label, m)
}

func main() {
	// Create a slice (not array, since length is not specificed), dynamic size
	// Go re-allocates a memory whenever a new item is added, costs performance
	userNames := []string{}

	// Index out of range
	// userNames[0] = "Jodhu"
	userNames = append(userNames, "Ram")
	userNames = append(userNames, "Shyam")

	fmt.Println("Usernames: ", userNames)

	// Pre allocates a slice, performance optimization, beyond that it re-allocates (default)
	// Params - type, preallocate, capacity
	preAllocateUserNames := make([]string, 2, 5)
	preAllocateUserNames[0] = "Ram"
	preAllocateUserNames[1] = "Shayam"

	fmt.Println("preAllocateUserNames: ", preAllocateUserNames)

	// Using map

	courseRatings := map[string]float64{}

	courseRatings["go"] = 4.9
	courseRatings["python"] = 4.6

	fmt.Println("Course Ratings: ", courseRatings)

	// Params - type, max capacity after which it re-allocates
	preAllocateCourseRatings := make(floatMap, 3)

	preAllocateCourseRatings["go"] = 4.9
	preAllocateCourseRatings["python"] = 4.6
	preAllocateCourseRatings["ocaml"] = 4.4

	preAllocateCourseRatings.output("preAllocateCourseRatings")

}
