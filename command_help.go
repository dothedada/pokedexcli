package main

import (
	"fmt"
)

func commandHelp(conf *config, param ...string) error {
	fmt.Println()
	fmt.Println("=======================")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, value := range getCommand() {
		fmt.Printf("%s -> %s\n", value.name, value.description)
	}
	fmt.Println("=======================")
	fmt.Println()
	return nil
}
