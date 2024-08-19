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
    - To type a pointer, just add * in front of the type of the data the pointer points to. So if the pointer is pointing to an Int, just type the pointer *Int.

Variable's scope
    - The variable is only available inside the function it was declared, including the main() function.
    - Therefore, if you have a variable declared inside main() and you try to use it in a function declared outside main(), it won't be possible.
    - You can declare variable on a higher level in your code, outside the main() function. In this way, this function will be available in all functions declared inside your file.
    - When declaring a variable outside all functions, you can't use the := shortcut.

Constants
    - A constant is a type of variable that can not be changed later and is declared using the `const` keyword

Strings
    - Strings are marked by a pair of double-quotes or backticks. Go doesn't allow single quotes.
    - When using strings, you can insert a value in the string using %v, %i and other keywords in the string. Therefore, to write a literal percentage, you need to type %%, so the second % is read as string.

Functions 
    - To declare a function, use the `func` keyword followed by parenthesis, the return types and, lastly the curly braces.
    - In the parenthesis, insert the arguments (separated by commas) followed by their types and inside the curly braces, the function body.
    - The return types of the function are declared between the arguments parenthesis and the curly braces. If there is only return element, just the type name is enough. If your function returns more than one element, you need to declare each return type wrapped by parenthesis and separate by comma.
    - Alternatively, you can declare a variable for each type, in the return type declaration, and use the variables in the function. Using this approach, you can just type the return type and the function will return the variable mentioned before in the order they were declared. All these features are optional.
    - If all parameters have the same type, you can cast the type at the end of the parameter's list.
    - Declare the function outside main()
    - Unlike many programming languages, you can return more than one value from a function, by separating them by a comma

If statements
    - The if statements is started by the keyword if followed by the conditional (the conditional is not wrapped in parenthesis). The conditionals can be stacked by using && (and) or || (or) operators. Then, curly braces holding the code to be run in case the condition is true.
    - If statements can be connected by if else and else keywords

Loops
    - In Go, there is only the for loop, which is very flexible.
    - The loop starts with the `for` keyword following by the validation, without parenthesis. The validation can follow the for or loop, like javascript (i := 0; i < 2; i++) or a conditional to behave like a while loop. 

Switch
    - Switch statements, in Go, receives a variable (without parenthesis) and add the cases, including the default, inside curly braces.
    - There is no need to add the `break` keyword after each case because Switches in Go only accept one case.
    - After each case value, add : and the code after that.
    - The `break` keywords jumps out from te switch code. 

Errors
    - Differently from other programming languages, errors in Go don't usually break the application. Usually, the error generates a error object but the application continues to run
    - Errors can be identified if they are different from nil
    - panic() function breaks the application and shows a error message

fmt package
    - The fmt package is a built-in package in GO that has a lot useful functions including the following
        -> fmt.Print() - prints a value in the console the way it is
        -> fmt.Println() - similarly to the above but adds a line break after the text
        -> fmt.Printf() - prints a value in the console (the first argument) but allows to format the content in several ways. See oficial documentation for more details
        -> Each of the functions above have a correspondent Saving function which are: fmt.Sprint(), fmt.Sprintf() and fmt.Sprintln() which, instead of printing the content to the console, they return the content so it can be saved in a variable.
        -> Another important function is the fmt.Scan() which gets input from the user in the console and saves the input in a variable. To identify what variable will hold that input, you should pass the variable pointer as the variable's name preceded by the & keyword. So something like fmt.Scan(&myVariable)

