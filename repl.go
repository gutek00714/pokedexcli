package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

var commands map[string]cliCommand

// var commands = map[string]cliCommand{
// 	"exit": {
// 		name:        "exit",
// 		description: "Exit the Pokedex",
// 		callback:    commandExit,
// 	},
// 	"help": {
// 		name: "help",
// 		description: "Displays a help message",
// 		callback: commandHelp,
// 	},
// }

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	return strings.Fields(lowered)
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	// CODE FROM REPL LESSON
	// for {
	// 	fmt.Print("Pokedex > ")
	// 	scanner.Scan()
	// 	line := cleanInput(scanner.Text())
	// 	if len(line) == 0 {
	// 		continue
	// 	} else {
	// 		first_word := line[0]
	// 		fmt.Println("Your command was:", first_word)
	// 	}
	// }

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := cleanInput(scanner.Text())
		if len(line) == 0 {
			continue
		} else {
			// get first word
			first_word := line[0]

			// check if the key(first_word) is in map
			cmd, exists := commands[first_word]
			if exists {
				err := cmd.callback()
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	// fmt.Printf("%s: %s\n", commands["help"].name, commands["help"].description)
	// fmt.Printf("%s: %s\n", commands["exit"].name, commands["exit"].description)
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}