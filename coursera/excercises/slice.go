package main
/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop.
Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/
import (
	"fmt"
	"sort"
	"os"
	"bufio"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	slice := make([]int, 3)
	regex, _ := regexp.Compile("^\\d+$")

	i := 0
	for {
		scanner := bufio.NewScanner(os.Stdin)
		var num int
		var str string

		fmt.Println("Please enter a number to continue or 'X' to quit")

		scanner.Scan()
		str = scanner.Text()

		if regex.MatchString(str) {
			num, _ = strconv.Atoi(str)
		} else {
			str = strings.ToUpper(str)

			if str == "X" {
				fmt.Println("Bye!")
				break
			} else {
				fmt.Println("Sorry, you entered an incorrect value.")
				continue
			}
		}

		if i < len(slice) {
			slice[0] = num
			sortSlice(slice)
		} else {
			slice = append(slice, num)
			sortSlice(slice)
		}

		i++
	}
}

func sortSlice(s []int) {
	sort.Ints(s)
	fmt.Println(s)
}