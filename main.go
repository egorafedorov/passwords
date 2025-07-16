package main

import (
	"fmt"
	"passwords/account"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("*** Password manager app ***")

Menu:
	for {
		userChoice := getMenu()
		switch userChoice {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
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

func createAccount() {
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
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func findAccount() {

}

func deleteAccount() {

}
