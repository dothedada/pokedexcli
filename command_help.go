package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println("=======================")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range getCommand() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	fmt.Println("=======================")
	fmt.Println()
	return nil
}
