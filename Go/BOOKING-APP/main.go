package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "GoGo"
	const totalTickets = 200
	var remainingTickets uint = 200
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint
	var bookings []string // bookings := []string{}

	fmt.Printf("Welcome to %v 2022 booking application", conferenceName)
	for {
		fmt.Println("Total tickets:", totalTickets, "\nTickets available now:", remainingTickets)
		fmt.Println("Enter first name: ")
		fmt.Scan(&userFirstName)
		fmt.Println("Enter last name: ")
		fmt.Scan(&userLastName)
		fmt.Println("Enter email: ")
		fmt.Scan(&userEmail)
		fmt.Println("Enter no. of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("\nThank you %v %v for booking %v tickets! You will receive a confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)

		fmt.Printf("\nRemaining tickets: %v", remainingTickets)
		bookings = append(bookings, userFirstName+" "+userLastName)
		firstNames := []string{}
		for _, item := range bookings {
			var names = strings.Fields(item)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("\nAll bookings first names: %v\n", firstNames)
	}
}
