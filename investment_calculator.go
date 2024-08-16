package main

import (
	"fmt" // math is a built in package with math related functions
	"math"
)

func main() {
	//using the variable declaration recommendations
	// var investmentAmount float64 = 1000
	// var years float64 = 10
	// var expectedReturnRate float64 = 5.5

	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5

	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	var futureValue float64 = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println(futureValue)
}

// About naming conventions:
//   - The file names, it is usual to use snake_case
//   - The module address it is usual to use kebab-case (like an url)
//   - variable names are usual camelCase

// Typing
// You can let Go guess the data type or set the type by writing the type after the variable

// Variables
// If you are not going to set the variable's type and want Go to set the variable's type, you should (recommended) use the following variable declaration shortcut
// instead of
// 		var variable = 1.5
// You should use
// 		variable := 1.5
