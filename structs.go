package main

import (
	"fmt"

	"example/structs/user"
)

// the type keyword can also be used to create an alias for another type.
// with that functionality, you can methods and assign the methods to other type, even the primitive types.
type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	// firstName := getUserData("Please enter your first name: ")
	// lastName := getUserData("Please enter your last name: ")
	// birthDate := getUserData("Please enter your birthDate (MM/DD/YYYY): ")

	//creating a struct using the explicit method
	var appUser *user.User
	appUser = &user.User{
		//the property names can be omitted, but the order of the properties needs to be the same the order in the struct type declaration

	}

	var myString str = "Testing if the type alias functionality"

	myString.log()

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthDate (MM/DD/YYYY): ")
	//creating another struct with the utility function
	anotherUser, err := user.NewUser(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println("Error Happened", err)
		return
	}
	fmt.Println(anotherUser)
	//  ... do something awesome with that gathered data!
	outPutUserDetails(*appUser)
	//To call the method, just call it by referencing the struct it belongs like below

}

func outPutUserDetails(user user.User) {
	//this approach the appUser struct is being passed, therefore copied in memory
	// fmt.Println(user.firstName, user.lastName, user.birthDate, user.createdAt)
	user.OutputUserDetailsMethod()
}

//to use the pointer instead of the struct itself the following should be applied.

/*
func outPutUserDetails(user *User) {
	fmt.Println((*user).firstName, (*user).lastName, (*user).birthDate, (*user).createdAt)

	but Go also allows this shortcut
	fmt.Println(user.firstName, user.lastName, user.birthDate, user.createdAt)
}
*/

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
