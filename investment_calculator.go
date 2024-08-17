package main

import (
	"fmt" // math is a built in package with math related functions
	"math"
)

func investmentCalculator() {
	const inflationRate = 2.5 //a constant can not be changed once declared
	var years float64
	var investmentAmount float64 //the variable is declared but it has no value. In this case, this variable will receive a null value that, in float64, it is 0.0
	expectedReturnRate := 5.5

	fmt.Print("Please, enter your investment: ")

	// the Scan() function gets input from the user. You need to pass a variable as first argument, but only the pointer, which you do by adding he & symbol in front of it
	fmt.Scan(&investmentAmount) // there is one limitation for fmt.Scan which is that it can only handle single words input. You can not use it to get phrases or complete sentences

	fmt.Print("Please, type the time, in years: ")
	fmt.Scan(&years)

	fmt.Print("Please, type the return rate: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue float64 = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
