package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main()  {
	
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		request, _ := reader.ReadString('\n')
		s := strings.Fields(request)
		animalSelected := strings.ToLower(s[0])
		action := strings.ToLower(s[1])
		
		var animal Animal

		switch animalSelected {
		case "cow":
			animal = Animal{"grass", "walk", "moo"}
		case "bird":
			animal = Animal{"worms", "fly", "peep"}
		case "snake":
			animal = Animal{"mice", "slither", "hsss"}
		default:
			fmt.Println("Sorry, we don't know that animal.")
			continue
		}

		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("A %s can't %s", animalSelected, action)
		}
	}
	
}

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}