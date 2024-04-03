package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	createdAt time.Time
}

type Admin struct {
	email   string
	country string
	User
}

func NewAdmin(email, country string) Admin {
	return Admin{
		email:   email,
		country: country,
		User: User{
			firstName: "Admin",
			lastName:  "Admini",
			createdAt: time.Now(),
		},
	}
}

func CreateAccount(guest User) {
	fmt.Println("Creating account using struct reference")
	fmt.Printf("%v %v | %v\n\n", guest.firstName, guest.lastName, guest.createdAt)
}

func CreateAccountPointer(g *User) {
	fmt.Println("Creating account using struct pointer")
	fmt.Printf("%v %v | %v\n\n", (*g).firstName, (*g).lastName, (*g).createdAt)
}

func (u User) AttachCreateAccountToStruct() {
	fmt.Println("Creating acoount using struct function attachment")
	fmt.Printf("%v %v | %v\n\n", u.firstName, u.lastName, u.createdAt)
}

func (u *User) ClearNames() {
	fmt.Println("Clearing first and last name")
	u.firstName = "-"
	u.lastName = "-"
}

func NewUserConstructor(firstName, lastName string) (User, error) {
	fmt.Println("Creating the struct using a constructor")

	if firstName == "" || lastName == "" {
		return User{}, errors.New("firstname and lastname are required")
	}

	// This creates a copy of the firstName and lastName variables,
	// We could instead return the pointer the struct
	return User{
		firstName: firstName,
		lastName:  lastName,
		createdAt: time.Now(),
	}, nil
}

func NewUserConstructorPointer(firstName, lastName string) (*User, error) {
	fmt.Println("Creating the struct pointer using a constructor")

	if firstName == "" || lastName == "" {
		return nil, errors.New("firstname and lastname are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		createdAt: time.Now(),
	}, nil
}
