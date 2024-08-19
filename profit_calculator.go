package main

import "fmt"

func profitCalculator() {
	var revenue = getUserInput("Revenue: ")
	var expenses = getUserInput("Expenses: ")
	var taxRate = getUserInput("Tax Rate: ")

	var ebt, profit, ratio = calculateProfitData(revenue, expenses, taxRate)

	fmt.Printf("The Profit before taxa is %v\n", ebt)
	fmt.Printf("The Profit after taxa is %v\n", profit)
	fmt.Printf("The ratio is %0.2f\n", ratio)
}

func calculateProfitData(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return
}

func getUserInput(text string) (userInput float64) {
	fmt.Print(text)
	fmt.Scan(&userInput)
	return
}
