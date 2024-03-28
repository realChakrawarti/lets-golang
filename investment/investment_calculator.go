package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount := takeInput("Investment Amount: ")
	expectedReturnRate := takeInput("Expected Return Rate: ")
	years := takeInput("Years: ")
	inflationRate := takeInput("Inflation Rate (enter 0 to ignore inflation): ")

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	inflationAdjustedFutureValue := futureValue * math.Pow(1+inflationRate/100, years)

	fmt.Printf("Amount on maturity: %0.2f\n", futureValue)
	fmt.Printf("Inflation adjusted: %0.2f\n", inflationAdjustedFutureValue)
}

func takeInput(label string) float64 {
	var inputValue float64
	fmt.Print(label)
	fmt.Scan(&inputValue)

	return inputValue
}
