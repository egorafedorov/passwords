package main

import (
	"fmt"
	"passwords/account"

	"github.com/fatih/color"
)

func main() {
	login := promptData("Enter your login")
	password := promptData("Enter your password")
	url := promptData("Enter URL")
	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		if err.Error() == "INVALID_LOGIN" {
			color.Red("Error! Login cannot be empty")
		} else if err.Error() == "INVALID_URL" {
			color.Red("Error! Invalid URL format")
		}
	}
	myAccount.OutputData()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
