This repo is a learning repo about the Golang

Topics:

Packages
    - The main package, along the main function, must be present in the application to compile appropriatedly
    - The code is organized in packages and each package can have many files

Variables
    - You can declare a variable using the `var` keyword, the variable's name and the type. Then followed by an equal sign and the content of the variable.
    - GO can guess the variable's type.
    - You can declare more than one variable in the same line of code, but either all variables have the same type (casted) or let GO type each variable.
    - There is a shortcut to declare a variable using the := operator, without the need of the `var` keyword
    - You can declare the variable first and then assign a value to it later. In this case, you need to cast the type and the variable will have a "null value" assigned automatically respective to the data type (0, "", 0.0 and so on)

Constants
    - A constant is a type of variable that can not be changed later and is declared using the `const` keyword

Strings
    - Strings are marked by a pair of double-quotes or backticks. Go doesn't allow single quotes.
    - When using strings, you can insert a value in the string using %v, %i and other keywords in the string. Therefore, to write a literal percentage, you need to type %%, so the second % is read as string.

fmt package
    - The fmt package is a built-in package in GO that has a lot useful functions including the following
        -> fmt.Print() - prints a value in the console the way it is
        -> fmt.Println() - similarly to the above but adds a line break after the text
        -> fmt.Printf() - prints a value in the console (the first argument) but allows to format the content in several ways. See oficial documentation for more details
        -> Each of the functions above have a correspondent Saving function which are: fmt.Sprint(), fmt.Sprintf() and fmt.Sprintln() which, instead of printing the content to the console, they return the content so it can be saved in a variable.

