package main

import (
	"booking-app/helper"
	"fmt"
	//"strings"
	//"strconv"
)

var conferenceName = "GoGo Conference"

const totalTickets = 200

var remainingTickets uint = 200
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func main() {
	fmt.Println("\nTotal tickets:", totalTickets)

	greetUsers()
	for remainingTickets > 0 {
		fmt.Printf("\nTickets available now:%v\n", remainingTickets)
		userFirstName, userLastName, userEmail, userTickets := getUserInput()

		if err := helper.ValidateUserInput(userFirstName, userLastName, userEmail, userTickets, remainingTickets); err != nil {
			fmt.Println(err)
		} else {
			remainingTickets = bookTicket(userTickets, remainingTickets, userFirstName, userLastName, userEmail)
		}
		if remainingTickets == 0 {
			//end prgm
			fmt.Printf("\nConference tickets sold out!")
			break
		}
	}
	fmt.Printf("\nAll bookings first names: %v\n", printFirstNames())
}

func greetUsers() {
	fmt.Printf("Welcome to %v 2022 booking application", conferenceName)
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, item := range bookings {
		firstNames = append(firstNames, item.firstName)
	}
	return firstNames
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

func bookTicket(userTickets uint, remainingTickets uint, userFirstName string, userLastName string, userEmail string) uint {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName: userFirstName,
		lastName:  userLastName,
		email:     userEmail,
		tickets:   userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("\nThank you %v %v for booking %v tickets! You will receive a confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("\nRemaining tickets: %v", remainingTickets)
	fmt.Printf("\nList of bookings: %v\n", bookings)
	return remainingTickets
}
