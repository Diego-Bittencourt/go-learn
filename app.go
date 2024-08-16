package main

// The above package is needed in Go.
// In Go, the code is organized in packages. One Go application can have several packages and each package can be composed of several files.
// The main package needs to be called main to sinalize Go which is the main package of the code
// You can use the functionality of one package to n another package, like is being done with the import statement below.

import "fmt"

// the package above is builted in Go so no need to download

func main() {
	fmt.Print("Hello World")
	// In Go, strings are defined by double quotes or backticks. Single quotes are not used for strings in Go.
}

// To create an executable file, the Go needs to compile, therefore it needs a module.
// The module can be created by command line, and it needs to the main package to start compiling.
// In windows, the compilation created a .exe file.

// Go code will not run from top to bottom. It will start running inside the main function, therefore it is necessary to have the main function inside the main package
