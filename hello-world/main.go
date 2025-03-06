package main

import (
	"fmt"
	"log"

	"example/hello/car"
)

func getPointerOfInt(number *int) {
	fmt.Printf("The address is %v and the value is %v", number, *number)
}

func main() {

	car := car.New("Toyota", "Corolla", 2020)
	someNUmber := 1
	getPointerOfInt(&someNUmber)
	car.Honk()

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Pam", "Jim", "Dwight", "Michael",
	}

	messages, err := my_hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
