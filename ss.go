//Declare a main package (a package is a way to group functions).
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console.
//This package is one of the standard library packages you got when you installed Go.
import "fmt"

//Implement a main function to print a message to the console.
//A main function executes by default when you run code in the file.

var sizeOfArray int = 10

func main() {
    arrayOfNumber := [10]int{2 ,6, 9, 8, 20, 15, 3, 5, 6, 9}
    min, max := findMinAndMax(arrayOfNumber)
    fmt.Print("Min = ")
    fmt.Println(min)
    fmt.Print("Max = ")
    fmt.Println(max)
}

func findMinAndMax(a [10]int) (min int, max int) {
	min = a[0]
    max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
        }
	}
	return min, max
}
