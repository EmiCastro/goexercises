package main
/*
Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”. 
*/
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputText string

	fmt.Println("Please, enter some string.")

	if _, err := fmt.Scan(&inputText); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	if conditionsAreMet(inputText) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func conditionsAreMet(inputText string) bool {
	text := strings.ToLower(inputText)
	return strings.HasPrefix(text, "i") &&
		strings.HasSuffix(text, "n") &&
		strings.Contains(text, "a")
}
