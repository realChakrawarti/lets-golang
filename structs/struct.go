package main

import (
	"fmt"
	"realChakrawarti/structs/user"
)

func takeInput(label string) string {
	var inputValue string
	fmt.Print(label)
	fmt.Scanln(&inputValue)

	return inputValue
}

func main() {
	firstName := takeInput("Enter your first name: ")
	lastName := takeInput("Enter your last name: ")

	guest, err := user.NewUserConstructor(firstName, lastName)

	if err != nil {
		fmt.Println(err)
		return
	}

	guestUsingPointer, err := user.NewUserConstructorPointer(firstName, lastName)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("john@example.com", "Bulgaria")
	admin.AttachCreateAccountToStruct() // Logs the name of Admin created, hard-coded

	user.CreateAccount(guest)                // Struct value reference
	user.CreateAccountPointer(&guest)        // Struct pointer reference
	guest.AttachCreateAccountToStruct() // Struct attached method using receiver arguments

	guest.ClearNames()   // Mutating the struct value reference
	user.CreateAccount(guest) // Log after mutating the value

	user.CreateAccountPointer(guestUsingPointer) // Logging the values passed as a reference
}
