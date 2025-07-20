package main

import (
	"fmt"
	"passwords/account"
	"passwords/files"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("*** Password manager app ***")
	vault := account.NewVault(files.NewJsonDb("data.json"))

Menu:
	for {
		vars := promptData([]string{
			"1. Create an account",
			"2. Find an account",
			"3. Delete account",
			"4. Exit",
			"Select the menu item",
		})
		switch vars {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		case "4":
			break Menu
		}
	}
}

func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Enter your login"})
	password := promptData([]string{"Enter your password"})
	url := promptData([]string{"Enter URL"})
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

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Enter URL to search for account"})
	accounts := vault.FindAccount(url)
	if len(accounts) == 0 {
		color.Red("No accounts found")
	}
	for _, account := range accounts {
		account.OutputData()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Enter the URL to delete your account"})
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		color.Green("Account deleted")
	} else {
		color.Red("Account not found")
	}
}
