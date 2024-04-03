package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	createdAt time.Time
}

func takeInput(label string) string {
	var inputValue string
	fmt.Print(label)
	fmt.Scan(&inputValue)

	return inputValue
}

func createAccount(guest user) {
	fmt.Printf("%v %v | %v", guest.firstName, guest.lastName, guest.createdAt)
}

func createAccountPointer(g *user) {
		fmt.Printf("%v %v | %v", (*g).firstName, (*g).lastName, (*g).createdAt)
}

func main() {
	firstName := takeInput("Enter your first name: ")
	lastName := takeInput("Enter your last name: ")

	guest := user{
		firstName: firstName,
		lastName:  lastName,
		createdAt: time.Now(),
	}

	createAccount(guest)
	createAccountPointer(&guest)
}
