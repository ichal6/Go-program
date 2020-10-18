//Declare a main package (a package is a way to group functions).
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console.
//This package is one of the standard library packages you got when you installed Go.
import "fmt"
import "sort"
//Implement a main function to print a message to the console.
//A main function executes by default when you run code in the file.

func testPackages() {
	arrayOfNumbers := []int{2 ,6, 9, 8, 20, 15, 3, 5, 6, 9}
	sort.Ints(arrayOfNumbers)
    min := arrayOfNumbers[0]
    max := arrayOfNumbers[len(arrayOfNumbers)-1]
    fmt.Println("Min = ", min)
    fmt.Println("Max = ", max)
}

