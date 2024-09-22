package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("hello"))

	names := []string{"Glady", "Baldy", "Moldy"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range messages {
		log.Println(msg)
	}
}
