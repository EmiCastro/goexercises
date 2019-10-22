package main

/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/
import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Enter some integer numbers")
	// Read input as string
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Parse to slice of ints
	ints, err := parseInts(line)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You entered this numbers: ", ints)

	// Divide slice between four goroutines
	q1, q2, q3, q4, err := divideSliceInQuarters(ints)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		fmt.Printf("Goroutine sorting: %+v\n", q1)
		sort.Ints(q1)
		wg.Done()
	}()

	go func() {
		fmt.Printf("Goroutine sorting: %+v\n", q2)
		sort.Ints(q2)
		wg.Done()
	}()

	go func() {
		fmt.Printf("Goroutine sorting: %+v\n", q3)
		sort.Ints(q3)
		wg.Done()
	}()

	go func() {
		fmt.Printf("Goroutine sorting: %+v\n", q4)
		sort.Ints(q4)
		wg.Done()
	}()

	wg.Wait()
	result := mergeQuarters(q1, q2, q3, q4)
	insertionsort(result)
	fmt.Println("Sorted numbers: ", result)
	fmt.Println("Bye!")
}

func mergeQuarters(q1, q2, q3, q4 []int) []int {
	result := make([]int, 0, len(q1)+len(q2)+len(q3)+len(q4))
	quarters := [][]int{q1, q2, q3, q4}
	for _, q := range quarters {
		result = append(result, q...)
	}
	return result
}

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

func parseInts(s string) ([]int, error) {

	fields := strings.Fields(s)
	ints := make([]int, len(fields))
	var err error

	for i, field := range fields {
		ints[i], err = strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
	}
	return ints, err

}

func divideSliceInQuarters(s []int) ([]int, []int, []int, []int, error) {
	if len(s) > 1 {
		middle := getMiddle(s)
		firstHalf := s[:middle]
		secondHalf := s[middle:]

		if len(s) > 3 {
			quarterMark1 := getMiddle(firstHalf)
			quarterMark2 := getMiddle(secondHalf)

			return firstHalf[:quarterMark1], firstHalf[quarterMark1:],
				secondHalf[:quarterMark2], secondHalf[quarterMark2:], nil
		}
		return firstHalf, secondHalf, nil, nil, errors.New("Not enough numbers, please enter at least four integers.")
	}
	return nil, nil, nil, nil, errors.New("Not enough numbers, please enter at least four integers.")
}

func getMiddle(s []int) int {
	if len(s) > 1 {
		return len(s) / 2
	} else {
		return len(s)
	}
}
