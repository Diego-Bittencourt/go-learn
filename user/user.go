package user

import (
	"errors"
	"fmt"
	"time"
)

// to create a type, just use the type keyword, recommended to do outside any function
// when defining a struct, the struct name should have the first letter capitalize to be available in other files, and followed by the keyword struct
// the properties have their time after them with no colons (:) nether commas between the property name and its type
type User struct {
	//both the struct name and the properties can be exported if their name start with a capital letter.
	//You can export the struct but keep the properties not exported by leaving the properties names on lower case
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}


// creating a constructor function.
// this is not a built-in functionality, bot more of a common practice
func NewUser(firstName string, lastName string, birthDate string) (*User, error) {

	//you can also add validation to the constructor
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("input is empty")
	}

	//you can also return a pointer to the struct because returned value also make a copy of the returned element
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

// struct has no class neither inheritance, but we can simulate that by embedding
type Admin struct {
	email    string
	password string
	User     //you cna give a name to that but also cal leave it anonymously. By doing so, you can call the properties and methods without calling its name
}

func NewAdmin(email, password string) *Admin {

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}
}

// You can also attach a function to a struct (like a method to an object) with the following syntax
func (userStruct User) OutputUserDetailsMethod() {
	// The first pair of parenthesis, before the function name, is the struct to which the function is attached and
	//the function can access the struct properties by referencing the element inside the parenthesis, called Receiver Argument or just Receiver
	fmt.Println(userStruct.firstName, userStruct.lastName, userStruct.birthDate, userStruct.createdAt)

}
