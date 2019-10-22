package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type aminal struct {
	food       string
	locomotion string
	noise      string
}

func (a *aminal) eat() {
	fmt.Printf("%s\n", a.food)
}

func (a *aminal) move() {
	fmt.Printf("%s\n", a.locomotion)
}

func (a *aminal) speak() {
	fmt.Printf("%s\n", a.noise)
}

func main() {

	// Reading string with numbers from command line
	fmt.Println("******   Enter the request as: <Animal> <request>             ******")
	fmt.Println("******   3 animals are supported 'cow', 'bird' and 'snake'    ******")
	fmt.Println("******   3 requestes are supported: 'eat', 'move' and 'speak' ******")
	fmt.Println("******   Press q\\Q to quit ******")

	for {
		fmt.Printf("> ")
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')

		// Converting string to interger slice
		line = strings.ReplaceAll(line, "\n", "")
		args := strings.Split(line, " ")
		if 2 != len(args) {
			if (1 == len(args)) && ("q" == strings.ToLower(args[0])) {
				break
			}
			fmt.Println("--- Error: Invalid input ---")
		}

		var animalObj aminal
		switch strings.ToLower(args[0]) {
		case "cow":
			animalObj = aminal{"grass", "walk", "moo"}
		case "bird":
			animalObj = aminal{"worms", "fly", "peep"}
		case "snake":
			animalObj = aminal{"mice", "slither", "hsss"}
		default:
			fmt.Printf("\n\"%s\" animal is not supported\n", args[0])
			return
		}

		switch strings.ToLower(args[1]) {
		case "eat":
			fmt.Printf("%s eats ", args[0])
			animalObj.eat()
		case "move":
			fmt.Printf("%s ", args[0])
			animalObj.move()
		case "speak":
			fmt.Printf("%s speaks ", args[0])
			animalObj.speak()
		default:
			fmt.Printf("\n\"%s\" request is not supported\n", args[1])
			return
		}
	}

}
