package main

import (
	"errors"
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
	fmt.Scanln(&inputValue)

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

func (u *user) clearNames() {
	fmt.Println("Clearing first and last name")
	u.firstName = "-"
	u.lastName = "-"
}

func newUserConstructor(firstName, lastName string) (user, error) {
	fmt.Println("Creating the struct using a constructor")

	if firstName == "" || lastName == "" {
		return user{}, errors.New("firstname and lastname are required")
	}

	// This creates a copy of the firstName and lastName variables,
	// We could instead return the pointer the struct
	return user{
		firstName: firstName,
		lastName:  lastName,
		createdAt: time.Now(),
	}, nil
}

func newUserConstructorPointer(firstName, lastName string) (*user, error) {
	fmt.Println("Creating the struct pointer using a constructor")

	if firstName == "" || lastName == "" {
		return nil, errors.New("firstname and lastname are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		createdAt: time.Now(),
	}, nil
}

func main() {
	firstName := takeInput("Enter your first name: ")
	lastName := takeInput("Enter your last name: ")

	guest, err := newUserConstructor(firstName, lastName)

	if err != nil {
		fmt.Println(err)
		return
	}

	guestUsingPointer, err := newUserConstructorPointer(firstName, lastName)

	if err != nil {
		fmt.Println(err)
		return
	}

	createAccount(guest)                // Struct value reference
	createAccountPointer(&guest)        // Struct pointer reference
	guest.attachCreateAccountToStruct() // Struct attached method using receiver arguments

	guest.clearNames()   // Mutating the struct value reference
	createAccount(guest) // Log after mutating the value

	createAccountPointer(guestUsingPointer) // Logging the values passed as a reference
}
