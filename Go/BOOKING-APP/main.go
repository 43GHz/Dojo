package main

import "fmt"

func main() {
	var conferenceName = "GoGo"
	const totalTickets = 200
	var remainingTickets = 200
	fmt.Printf("Welcome to %v 2022 booking application", conferenceName)
	fmt.Println("Total tickets:",totalTickets,"\nTickets available now:",remainingTickets)
}
