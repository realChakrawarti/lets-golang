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
	fmt.Println("Creating account using struct reference")
	fmt.Printf("%v %v | %v\n\n", guest.firstName, guest.lastName, guest.createdAt)
}

func createAccountPointer(g *user) {
	fmt.Println("Creating account using struct pointer")
	fmt.Printf("%v %v | %v\n\n", (*g).firstName, (*g).lastName, (*g).createdAt)
}

func (u user) attachCreateAccountToStruct() {
	fmt.Println("Creating acoount using struct function attachment")
	fmt.Printf("%v %v | %v\n\n", u.firstName, u.lastName, u.createdAt)
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
	guest.attachCreateAccountToStruct()
}
