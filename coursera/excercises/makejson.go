package main

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/
import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

func main() {
	var name string
	var address string
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Hello, what's your name?")
	scanner.Scan()
	name = scanner.Text()
	
	fmt.Println("Please, enter your address.")
	scanner.Scan()
	address = scanner.Text()
	
	
	userMap := make(map[string]string)
	userMap["name"] = name
	userMap["address"] = address

	userJSON, _ := json.Marshal(userMap)

	fmt.Println(string(userJSON))
}