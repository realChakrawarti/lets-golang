package main

import "fmt"

func main() {
	products := [4]string{"A", "B", "C", "D"}
	prices := [4]float64{9.99, 12.99, 13.99, 19.99}
	fmt.Println(prices, products)
	
	// Update an item
	products[2] = "New Product"
	fmt.Println("Products: ", products)

	// Slicing array (creating new list from old ones), select 1st and go till 3rd index while excluding it
	heroProductPrices := prices[1:3] // 12.99 13.99
	fmt.Println(heroProductPrices)
	
	// Excluding parts of an array
	heroProductPrices = prices[:3]
	fmt.Println("Till 3rd index while excluding it ",heroProductPrices)
	heroProductPrices = prices[1:]
	fmt.Println("Till last index while excluding the 0th index ", heroProductPrices)

	// len(): Length of a slice or an array
	fmt.Println("Length",len(heroProductPrices)) 
	
	// cap(): Capacity of the array
	fmt.Println("Capacity",len(heroProductPrices)) 

	// Dynamic slices

	dynamicSlice := []float64{0.99, 1.99}

	fmt.Println("Dynamice Slice initialized", dynamicSlice)
	
	// Add items to the existing slice
	dynamicSlice = append(dynamicSlice, 3.99)
	
	fmt.Println("Dynamic Slice modified", dynamicSlice)


	// Unpacking list values
	newItems := []float64{2.00, 3.14, 6.28}
	unpackedValues := append(dynamicSlice, newItems... )

	fmt.Println("Unpacking and appending using js rest-like operator", unpackedValues)

	practice()

}