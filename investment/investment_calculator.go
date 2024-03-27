package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years, inflationRate float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Inflation Rate (enter 0 to ignore inflation): ")
	fmt.Scan(&inflationRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	inflationAdjustedFutureValue := futureValue * math.Pow(1+inflationRate/100, years)

	fmt.Println("Amount on maturity: ", futureValue)
	fmt.Println("Inflation adjusted: ", inflationAdjustedFutureValue)
}
