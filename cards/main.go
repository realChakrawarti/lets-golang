package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades" // Type is inferred and := can only be used for initialization
	card = "King of Hearts" // Can be re-declared
	anotherCard := testInferType() // A return type must be provided from the function to be inferred
	sliceData := deck{"Hello", testInferType()}
	sliceData = append(sliceData, "World")
	fmt.Println(sliceData)
	fmt.Println(card)
	fmt.Println(anotherCard)

	for i, cardInfo := range sliceData {
		fmt.Println(i, cardInfo)
	}
}

func testInferType() string {
	return "Jack of Club"
}