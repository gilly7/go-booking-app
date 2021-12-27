package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"

	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Println("Welcome to ", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	for {
		var bookings []string

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask user for their name
		fmt.Println("Enter your First Name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		//	bookings[0] = firstName + " " + lastName

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings:  %v\n", bookings)
	}
}
