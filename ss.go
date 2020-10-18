//Declare a main package (a package is a way to group functions).
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console.
//This package is one of the standard library packages you got when you installed Go.
import "fmt"
//Implement a main function to print a message to the console.
//A main function executes by default when you run code in the file.

func main() {
    arrayOfNumbers := []int{2 ,6, 9, 8, 20, 15, 3, 5, 6, 9}
    min, max := findMinAndMax(arrayOfNumbers)
    average := findAverage(arrayOfNumbers)
    fmt.Println("Min = ", min)
    fmt.Println("Max = ", max)
    fmt.Println("Average = ", average)

    testPackages()
}

func findMinAndMax(a []int) (min int, max int) {
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

func findAverage(xs []int) int {
    total := int(0)
    for _, x := range xs {
      total += x
    }
    return total / int(len(xs))
}