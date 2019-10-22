package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	var animals = make(map[string]Animal)
	fmt.Println("******   Welcome! Enter a new animal as: newanimal <name> <typeOfAnimal>		    ******")
	fmt.Println("******   3 animals are supported 'cow', 'bird' and 'snake'   				    ******")
	fmt.Println("******   Once you've created at least one animal you can query and requets actions from it. ******")
	fmt.Println("******   Request an action as: query <name> <action>    				    ******")
	fmt.Println("******   3 requestes are supported: 'eat', 'move' and 'speak'				    ******")
	fmt.Println("******   Press q\\Q to quit 								    ******")

	for {
		fmt.Printf("> ")
		command, name, info, err := readInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if command == "quit" {
			fmt.Println("Bye!")
			break
		}

		switch command {
		case "newanimal":
			if err := createAnimal(name, info, animals); err != nil {
				fmt.Println(err)
				continue
			} else {
				fmt.Println("Created it!")
			}
		case "query":
			if err := queryAction(name, info, animals); err != nil {
				fmt.Println(err)
				continue
			}
		default:
			fmt.Printf("\n\"%s\" request is not supported\n", command)
			continue
		}
	}
}

func readInput() (string, string, string, error) {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	line = strings.ReplaceAll(line, "\n", "")
	args := strings.Split(line, " ")

	if 3 != len(args) {
		if (1 == len(args)) && ("q" == strings.ToLower(args[0])) {
			return "quit", "", "", nil
		}
		return "", "", "", errors.New("Error: Invalid input. You should enter 3 strings separated by spaces.")
	}

	return strings.ToLower(args[0]), strings.ToLower(args[1]), strings.ToLower(args[2]), nil
}

func createAnimal(name, specie string, animals map[string]Animal) error {
	switch specie {
	case "cow":
		animals[name] = Cow{"grass", "walk", "moo"}
	case "snake":
		animals[name] = Snake{"mice", "slither", "hsss"}
	case "bird":
		animals[name] = Bird{"worms", "fly", "peep"}
	default:
		return errors.New("Animal type not found")
	}
	return nil
}

func queryAction(name, action string, animals map[string]Animal) error {

	animal := animals[name]
	if animal == nil {
		return errors.New("Animal not found")
	}
	switch action {
	case "move":
		animal.Move()
	case "eat":
		animal.Eat()
	case "speak":
		animal.Speak()
	default:
		return errors.New("Unavailable action.")
	}

	return nil
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func ExcecuteAction(action string, a Animal) error {
	switch strings.ToLower(action) {
	case "move":
		a.Move()
	case "eat":
		a.Eat()
	case "speak":
		a.Speak()
	default:
		return errors.New("Unavailable action.")
	}
	return nil
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}
