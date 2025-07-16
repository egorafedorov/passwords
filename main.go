package main

import (
	"fmt"
	"passwords/account"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("*** Password manager app ***")
	vault := account.NewVault()

Menu:
	for {
		userChoice := getMenu()
		switch userChoice {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		case 4:
			break Menu
		}
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func getMenu() int {
	var userChoice int
	fmt.Println("1. Create an account")
	fmt.Println("2. Find an account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")
	fmt.Print("Select the menu item: ")
	fmt.Scanln(&userChoice)
	return userChoice
}

func createAccount(vault *account.Vault) {
	login := promptData("Enter your login")
	password := promptData("Enter your password")
	url := promptData("Enter URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		if err.Error() == "INVALID_LOGIN" {
			color.Red("Error! Login cannot be empty")
		} else if err.Error() == "INVALID_URL" {
			color.Red("Error! Invalid URL format")
		}
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.Vault) {
	url := promptData("Enter URL to search for account")
	accounts := vault.FindAccount(url)
	if len(accounts) == 0 {
		color.Red("No accounts found")
	}
	for _, account := range accounts {
		account.OutputData()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Enter the URL to delete your account")
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		color.Green("Account deleted")
	} else {
		color.Red("Account not found")
	}
}
