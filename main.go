package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	conferenceTickets := 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var bookings []string

	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 1 && len(lastName) >= 1
		isValidEmail := IsValidEmail(email)
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v, \n", firstName, lastName, userTickets, email)
		} else {
			if !isValidName {
				fmt.Println("Invalid First name of last name. Both need to be longer than 1 character.")
				fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
				fmt.Println("Enter your first name: ")
				fmt.Scan(&firstName)

				fmt.Println("Enter your last name: ")
				fmt.Scan(&lastName)
			}
			if !isValidEmail {
				fmt.Println("Invalid Email. Email needs a @.")
				fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
				fmt.Println("Enter your email: ")
				fmt.Scan(&email)
				if !IsValidEmail(email) {
					continue
				}
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid Ticket Number. Please enter a valid ticket number.")
				fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
				fmt.Println("Enter number of tickets: ")
				fmt.Scan(&userTickets)
			}
		}
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
		fmt.Printf("These are all our bookings: %v\n", bookings)

		if remainingTickets == 0 {
			fmt.Printf("All tickets are booked for %v. Thank you!", conferenceName)
			break
		}
	}
}

func IsValidEmail(email string) bool {
	isValidEmail := strings.Contains(email, "@")
	return isValidEmail
}
