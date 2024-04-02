package main

import (
	"fmt"
	"realChakrawarti/go-bank/fileops"
)

const BANK_ACCOUNT_FILE = "balance.txt"
const DEFAULT_ACCOUNT_BALANCE = 500

func main() {

	accountBalance, err := fileops.ReadValueFromFile(BANK_ACCOUNT_FILE, DEFAULT_ACCOUNT_BALANCE)

	if err != nil {
		fmt.Printf("\nError occured while trying to connect to the server: %v", err)
	}

	for {
		displayOptions()
		choice := takeInput("Enter option: ")

		if choice == 1 {
			fmt.Printf("Account balance: %0.2f\n", accountBalance)
		} else if choice == 2 {
			depositAmount := takeInput("Amount to be deposited: ")
			if depositAmount <= 0 {
				fmt.Println("Invalid amount provided")
				continue
			} else {
				accountBalance += depositAmount
				fileops.WriteValueToFile(accountBalance, BANK_ACCOUNT_FILE)
				fmt.Printf("Account balance (updated): %0.2f\n", accountBalance)
			}

		} else if choice == 3 {
			withdrawalAmount := takeInput("Amount to be withdrawn: ")

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount provided")
				continue
			} else if withdrawalAmount > accountBalance {
				fmt.Println("Withdrawal amount cannot be more than account balance!")
				continue
			} else {
				accountBalance -= withdrawalAmount
				fileops.WriteValueToFile(accountBalance, BANK_ACCOUNT_FILE)
				fmt.Printf("Account balance (updated): %0.2f\n", accountBalance)
			}

		} else if choice == 4 {
			fmt.Print("Thank you for banking with us!\n")
			break
		} else {
			fmt.Print("Invalid choice provided.")
			continue
		}
	}

}

func takeInput(label string) float64 {
	var inputValue float64
	fmt.Print(label)
	fmt.Scan(&inputValue)

	return inputValue
}
