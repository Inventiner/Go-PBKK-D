package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	// ask users for their personal data
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your firstname")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastname")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter numbers of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v Remaining tickets for %v\n", remainingTickets, conferenceName)

			// show first name of bookings
			firstNames := []string{}
			for index, booking := range bookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				if index == (len(bookings) - 1) {
					firstNames = append(firstNames, firstName)
				} else {
					firstNames = append(firstNames, firstName+",")
				}
			}
			fmt.Printf("These first names of our bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our Conventions is booked out. Come back next year :).")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your entered first name or last name is too short!")
			}
			
			if !isValidEmail {
			fmt.Println("Your entered email address is invalid!")
			} 
			
			if !isValidTicketNumber {
				fmt.Println("Your entered number of ticket is invalid!")
			}
		}

	}

}
