//Declare a main package (a package is a way to group functions).
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console.
//This package is one of the standard library packages you got when you installed Go.
import "fmt"
//You can use the pkg.go.dev site to find published modules whose packages have functions you can use in your own code.
import "rsc.io/quote"

//Implement a main function to print a message to the console.
//A main function executes by default when you run code in the file.
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}