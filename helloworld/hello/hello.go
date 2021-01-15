package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	//Set properties of predefined Logger
	// Including the log entry prefix and flag to disable printing time, source file, line number
	log.SetPrefix("Error greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	//Get a greeting message and print it
	message, err := greetings.Hellos(names)
	//If error was returned, print to console and exit
	if err != nil {
		log.Fatal(err)
	}

	//If no error was returned, print returned message
	fmt.Println(message)
}