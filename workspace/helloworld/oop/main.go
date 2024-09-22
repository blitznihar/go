package main

import (
	"fmt"
	"log"
	"oop/employee"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World")
	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
	fmt.Println(quote.Go())
	message, err := greetings.Hello("adasd ")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)

	}
	fmt.Println(message)
	log.Println(message)

	names := []string{"Glady", "Baldy", "Moldy"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(messages)

	for _, msg := range messages {
		log.Println(msg)
	}

}
