package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	cfg := &config{}

	// initsalize the cli
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
	// sets up an infinate loop so program wont exit unless we exit
	for {

		fmt.Print("pokedex >")
		scanner.Scan()
		input := scanner.Text() // checks for input
		ci := cleanInput(input)
		comand, exists := commands[ci[0]]
		if exists {
			err := comand.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// comand functions
func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandhelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	if cfg.Next == nil {
		locationAreaURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		locationAreaURL = *cfg.Next
	}
	res, err := http.Get(locationAreaURL)
	if err != nil {
		return fmt.Errorf("failed to get response from client:%w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll failed: %w", err)
	}
	var locationArea LocationArea
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed: %w", err)
	}

	for _, result := range locationArea.Results {
		fmt.Println(result.Name)

	}

	cfg.Next = &locationArea.Next
	cfg.Previous = &locationArea.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
	} else {
		locationAreaURL = *cfg.Previous
	}

	res, err := http.Get(locationAreaURL)
	if err != nil {
		return fmt.Errorf("failed to get response from client: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll failed: %w", err)
	}
	var locationArea LocationArea
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed: %w", err)
	}

	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}

	cfg.Next = &locationArea.Next
	cfg.Previous = &locationArea.Previous

	return nil
}

type config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func cleanInput(text string) []string {
	var cleaned []string
	t := strings.TrimSpace(text)
	t = strings.ToLower(t)
	cleaned = strings.Fields(t)
	return cleaned
}

// struct and variable setup
var commands map[string]cliCommand

var locationAreaURL string

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
