package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Desplays commands",
			callback:    commandhelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next twenty locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous twenty locations",
			callback:    commandMapb,
		},
	}
	for {

		fmt.Print("pokedex >")
		scanner.Scan()
		input := scanner.Text()
		ci := cleanInput(input)
		comand, exists := commands[ci[0]]
		if exists {
			err := comand.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandhelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: \n")

	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap() error {
	return nil
}

func commandMapb() error {
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func cleanInput(text string) []string {
	var cleaned []string
	t := strings.TrimSpace(text)
	t = strings.ToLower(t)
	cleaned = strings.Fields(t)
	return cleaned
}
