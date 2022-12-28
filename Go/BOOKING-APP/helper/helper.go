package helper

import (
	"errors"
	"strings"
)

func ValidateUserInput(userFirstName string, userLastName string, userEmail string, userTickets uint, remainingTickets uint) error {
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
