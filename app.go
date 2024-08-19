package main

import "fmt"

func main() {
	fmt.Println("Select the service you want to use: ")
	fmt.Println("1. Go Bank")
	fmt.Println("2. Go Profit")
	fmt.Println("3. Go Investment")
	fmt.Print("Your choice: ")
	var chosenService int
	fmt.Scan(&chosenService)
	switch chosenService {
	// in Go, the switch statement only allows one case, so no need for the break statement after each case
	case 1:
		bank()
	case 2:
		profitCalculator()
	case 3:
		investmentCalculator()
	default:
		fmt.Println("Select a valid choice")
	}
}
