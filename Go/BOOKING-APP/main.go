package main

import (
	"errors"
	"fmt"
	"strings"
)

var conferenceName = "GoGo Conference"

const totalTickets = 200

var remainingTickets uint = 200

func greetUsers() {
	fmt.Printf("Welcome to %v 2022 booking application", conferenceName)
}

func printFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, item := range bookings {
		var names = strings.Fields(item)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(userFirstName string, userLastName string, userEmail string, userTickets uint) error {
	isInvalidName := len(userFirstName) < 2 || len(userLastName) < 2
	if isInvalidName {
		return errors.New("you have entered invalid name")
	}
	isValidEmail := strings.Contains(userEmail, "@")
	if !isValidEmail {
		return errors.New("you have entered invalid email")
	}
	isValidTicketNo := userTickets <= remainingTickets && userTickets > 0
	if !isValidTicketNo {
		return errors.New("you have entered invalid number of tickets")
	}
	return nil
}

func getUserInput() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint
	fmt.Println("Enter first name: ")
	fmt.Scan(&userFirstName)
	fmt.Println("Enter last name: ")
	fmt.Scan(&userLastName)
	fmt.Println("Enter email: ")
	fmt.Scan(&userEmail)
	fmt.Println("Enter no. of tickets: ")
	fmt.Scan(&userTickets)
	return userFirstName, userLastName, userEmail, userTickets
}

func main() {
	var bookings []string // bookings := []string{}
	fmt.Println("\nTotal tickets:", totalTickets)

	greetUsers()
	for remainingTickets > 0 {
		fmt.Printf("\nTickets available now:%v\n", remainingTickets)
		userFirstName, userLastName, userEmail, userTickets := getUserInput()

		if err := validateUserInput(userFirstName, userLastName, userEmail, userTickets); err != nil {
			fmt.Println(err)
		} else {
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("\nThank you %v %v for booking %v tickets! You will receive a confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)

			fmt.Printf("\nRemaining tickets: %v", remainingTickets)
			bookings = append(bookings, userFirstName+" "+userLastName)

			fmt.Printf("\nAll bookings first names: %v\n", printFirstNames(bookings))
		}
		if remainingTickets == 0 {
			//end prgm
			fmt.Printf("\nConference tickets sold out!")
			break
		}
	}
}
