package main
/*Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
*/
import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please, introduce up to ten numbers:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	ints, err := parseInts(input)
	if err != nil {
		log.Fatal(err)
	}
	
	bubbleSort(ints)
	fmt.Println("Sorted integers:", ints)
}

func bubbleSort(slice []int) {
	end := len(slice)-1
	for range slice {
		for i := 0; i < end; i++ {
			if slice[i] > slice[i + 1]{
				swap(slice, i)
			}
		}
	}
}

func swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func parseInts(s string) ([]int, error) {
	var (
		fields = strings.Fields(s)
		ints   = make([]int, len(fields))
		err    error
	)
	for i, v := range fields {
		ints[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

