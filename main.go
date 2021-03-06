package main

import (
	"fmt"
	"strings"
)

func main() {

	greetUsers()
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

		//user input validations
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			//	bookings[0] = firstName + " " + lastName

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			//fmt.Printf("These are all our bookings:  %v\n", bookings)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break

			}
		} else {

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")

			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @sign ")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}

}

func greetUsers() {
	fmt.Println("Welcome to our conference")
}
