package main
/*
Write a program which prompts the user to enter a floating point number and prints the integer 
which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
*/
import (
	"fmt"
	"os"
)

func main() {
	var floatNum float32
	fmt.Printf("Please, enter a floating point number :)\n")

	if _, err := fmt.Scan(&floatNum); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(int(floatNum))
}
