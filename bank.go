package main

import (
	"fmt"

	"investment-calculator.com/fileops"
)

const balanceFileName = "balance.txt"

func bank() {
	var accountBalance, err = fileops.ReadFloatFromFile(balanceFileName)

	if err != nil {
		fmt.Println("ERROR OCURRED")
		fmt.Println(err)
		fmt.Println("-----------------------------------------")
		// panic("the application could't proceed. Please get in touch with the support")
		//The panic() function breaks the application and gives a messages that looks more like an error
	}
	var isItRunning bool = true
	fmt.Println("Welcome to the Go Bank Services")
	for isItRunning {
		//in Go, there is only the for loop, but it is very flexible
		// You can also declare a variable, create a conditional and change the variable at the end of each cycle
		//for i := 0; i < 2; i++
		fmt.Println("Select your desired option")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var userSelectedMenu int
		fmt.Print("Your choice: ")
		fmt.Scan(&userSelectedMenu)

		if userSelectedMenu == 1 {
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		} else if userSelectedMenu == 2 {
			var addedValue float64
			fmt.Print("Enter the deposited value: ")
			fmt.Scan(&addedValue)
			// accountBalance = accountBalance + addedValue
			accountBalance += addedValue
		} else if userSelectedMenu == 3 {
			var withdrawnValue float64
			fmt.Print("Enter withdrawn value: ")
			fmt.Scan(&withdrawnValue)
			if withdrawnValue > accountBalance {
				fmt.Println("Not enough balance")
				continue //the continue keyword tells Go to ignore the rest of the current's iteration code and skip to the next loop.
			}
			// accountBalance = accountBalance - withdrawnValue
			accountBalance -= withdrawnValue

		} else if userSelectedMenu == 4 {
			break
		} else {
			fmt.Println("Please, enter a valid choice.")
		}
		fmt.Println("----------------------------------")
	}
	fileops.WriteFloatToFile(balanceFileName, accountBalance)
	fmt.Println("Thank you for using our services")
}
