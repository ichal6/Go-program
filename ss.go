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
    fmt.Print(bubbleSort(arrayOfNumbers), "\n")
    testPackages()
}

func findMinAndMax(array []int) (min int, max int) {
	min = array[0]
    max = array[0]
	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
        }
	}
	return min, max
}

func findAverage(array []int) int {
    total := int(0)
    for _, value := range array {
      total += value
    }
    return total / int(len(array))
}

func bubbleSort(array []int) []int {
    arrayLength := len(array)

    for i := 0; i < arrayLength-1; i++ {
        for j := 0; j < arrayLength-i-1; j++ {
            if array[j] > array[j+1]{
                array[j], array[j+1] = array[j+1], array[j]
            }
        }
    }
    return array
}

