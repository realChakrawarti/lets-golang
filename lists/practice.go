package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func product_update() {
	products := []Product{
		{
			"iphone-6", "iPhone 6", 499.99,
		},
		{
			"iphone-x", "iPhone-X", 599.99,
		},
		{
			"airpod-3", "AirPod 3 ANC", 299.99,
		},
	}

	fmt.Println("Products are: ", products)

	windowsSys := Product{
		"surface-pro", "Surface Pro", 899.99,
	}

	products = append(products, windowsSys)

	fmt.Println("Updated product list: ", products)
}

func practice() {
	hobbies := [3]string{"Singing", "Drumming", "Walking"}
	fmt.Println("My hobbies are: ", hobbies)

	fmt.Println("First hobby: ", hobbies[0])
	fmt.Println("Second and Third hobby: ", hobbies[1:])

	// hobbiesSlice := hobbies[0:2]
	hobbiesSlice := hobbies[:2]
	fmt.Println("Hobbies slice: ", hobbiesSlice)

	hobbiesSlice = hobbiesSlice[1:3]
	fmt.Println("Updated hobbies slice: ", hobbiesSlice)

	// Dynamic Array
	newHobbies := []string{"Coffee", "Chai", "Coding"}
	fmt.Println("New hobbies: ", newHobbies)
	newHobbies[0] = "Beer"

	newHobbies = append(newHobbies, "Coffee")

	fmt.Println("Updated new hobbies: ", newHobbies)

	product_update()
}
