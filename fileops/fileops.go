package fileops

import (
	"errors" //package responsible for creating errors
	"fmt"
	"os"      //package that can interact with the system
	"strconv" //string conversion
)

func WriteFloatToFile(filename string, balance float64) {
	balanceString := fmt.Sprint(balance)                //saving the float into string
	os.WriteFile(filename, []byte(balanceString), 0644) //[]byte means a collection of bytes.
}

func ReadFloatFromFile(filename string) (float64, error) {
	valueFromFile, err := os.ReadFile(filename)
	if err != nil {
		//nil is a value that the error receives if there are no errors
		fmt.Println(err)
		return 1000, errors.New("the file was not found")
	}
	valueFloat := string(valueFromFile) //converting to string because float64 does not work with byte
	balanceFloat, err := strconv.ParseFloat(valueFloat, 64)

	if err != nil {
		return 1000, errors.New("failed to read the file content")
	}
	return balanceFloat, nil
}
