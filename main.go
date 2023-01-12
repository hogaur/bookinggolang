package main

import "fmt"

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

		for userTickets > remainingTickets {
			fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
			fmt.Println("Enter number of tickets: ")
			fmt.Scan(&userTickets)
		}

		if userTickets != 0 {
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v, \n", firstName, lastName, userTickets, email)
			fmt.Printf("The whoel slice : %v\n", bookings)
			fmt.Printf("The first value : %v\n", bookings[0])
			fmt.Printf("The slice type : %T\n", bookings)
			fmt.Printf("The slice length : %v\n", len(bookings))
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
