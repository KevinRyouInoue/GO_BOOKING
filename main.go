package main

import "fmt"

func main() {
	conferenceName := "GO CONFERENCE"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	fmt.Println("Get Your Ticket Here to Attend!")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask use for input
	fmt.Println("Enter Your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your email:")
	fmt.Scan(&email)

	fmt.Println("Enter Your number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("THANK YOU %v %v booked %v tickets. Sent email at %v\n", firstName, lastName, userTickets, email)
}
