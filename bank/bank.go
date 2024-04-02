package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const BANK_ACCOUNT_FILE = "balance.txt"
const DEFAULT_ACCOUNT_BALANCE = 500

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(BANK_ACCOUNT_FILE, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(BANK_ACCOUNT_FILE)

	if err != nil {
		return DEFAULT_ACCOUNT_BALANCE, errors.New("server down")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return DEFAULT_ACCOUNT_BALANCE, errors.New("server is busy. please try again after sometime")
	}

	return balance, nil
}

func main() {

	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Printf("\nError occured while trying to connect to the server: %v", err)
	}

	for {
		fmt.Print("\n-------------------------------------------------------\n")
		fmt.Println("Welcome to Go bank!")
		fmt.Print("-------------------------------------------------------\n")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Amount")
		fmt.Println("3. Withdrawal Amount")
		fmt.Println("4. Exit")

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
				writeBalanceToFile(accountBalance)
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
				writeBalanceToFile(accountBalance)
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
