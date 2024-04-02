package main

import "fmt"

func displayOptions() {
	fmt.Print("\n-------------------------------------------------------\n")
	fmt.Println("Welcome to Go bank!")
	fmt.Print("-------------------------------------------------------\n")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Amount")
	fmt.Println("3. Withdrawal Amount")
	fmt.Println("4. Exit")
}
