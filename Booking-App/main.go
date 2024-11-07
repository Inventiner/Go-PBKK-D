package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceName string = "Go Conference"

const conferenceTicketsLondon int = 125
const conferenceTicketsNewYorkCity int = 150
const conferenceTicketsAmsterdam int = 75

var remainingTicketsLondon uint = uint(conferenceTicketsLondon)
var remainingTicketsNewYorkCity uint = uint(conferenceTicketsNewYorkCity)
var remainingTicketsAmsterdam uint = uint(conferenceTicketsAmsterdam)

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	venue           string
	numberOfTickets uint
}

func main() {
	greetUsers()

	for {
		var venue int

		fmt.Printf("Please select Venue Location:\n [1] London [2] New York City [3] Amsterdam [4] Ticket Status [5] See All Bookings [-1] Exit\n")
		fmt.Scan(&venue)

		switch venue {
		case 1:
			remainingTicketsLondon = processBooking(remainingTicketsLondon, "London")
		case 2:
			remainingTicketsNewYorkCity = processBooking(remainingTicketsNewYorkCity, "New York City")
		case 3:
			remainingTicketsAmsterdam = processBooking(remainingTicketsAmsterdam, "Amsterdam")
		case 4:
			fmt.Printf("We have %v tickets for London, %v tickets for New York City, and %v tickets for Amsterdam\n", remainingTicketsLondon, remainingTicketsNewYorkCity, remainingTicketsAmsterdam)
		case 5:
			getAllBookings()
		case -1:
			fmt.Println("Thank you for using our booking application. Goodbye!")
		default:
			fmt.Println("Invalid Venue Location. Please try again.")
			continue
		}

		if venue == -1 {
			break
		} else {
			fmt.Printf("\n============================\n")
		}
	}

	wg.Wait()
}

var wg = sync.WaitGroup{}

func processBooking(remainingTickets uint, location string) uint {
	UserData := getUserInputs(location)
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(UserData, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		remainingTickets = bookTickets(UserData, remainingTickets)
		wg.Add(1)
		go sendTicket(UserData)

		firstNames := getFirstNames()
		fmt.Printf("These first names of our bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("Our Conventions for %v is booked out. Try checking other venue or check again next year :).", location)
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

	return remainingTickets
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and shared across all three venue\n", conferenceTicketsLondon+conferenceTicketsNewYorkCity+conferenceTicketsAmsterdam)
	fmt.Printf("We have %v tickets for London, %v tickets for New York City, and %v tickets for Amsterdam\n", remainingTicketsLondon, remainingTicketsNewYorkCity, remainingTicketsAmsterdam)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for index, booking := range bookings {
		if index == (len(bookings) - 1) {
			firstNames = append(firstNames, booking.firstName)
		} else {
			firstNames = append(firstNames, booking.firstName+",")
		}
	}
	return firstNames
}

func getUserInputs(location string) UserData {
	var data UserData

	data.venue = location

	fmt.Println("Enter your firstname")
	fmt.Scan(&data.firstName)

	fmt.Println("Enter your lastname")
	fmt.Scan(&data.lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&data.email)

	fmt.Println("Enter numbers of tickets: ")
	fmt.Scan(&data.numberOfTickets)

	return data
}

func bookTickets(data UserData, remainingTickets uint) uint {
	remainingTickets = remainingTickets - uint(data.numberOfTickets)
	bookings = append(bookings, data)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", data.firstName, data.lastName, data.numberOfTickets, data.email)
	fmt.Printf("%v Remaining tickets for %v %v\n", remainingTickets, conferenceName, data.venue)

	return remainingTickets
}

func validateUserInput(data UserData, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(data.firstName) >= 2 && len(data.lastName) >= 2
	isValidEmail := strings.Contains(data.email, "@")
	isValidTicketNumber := data.numberOfTickets > 0 && data.numberOfTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getAllBookings() {
	var password string

	fmt.Println("Enter Administrator password to view all bookings: ")
	fmt.Scan(&password)

	if password != "ADMIN" {
		fmt.Println("Invalid password. Please try again.")
	} else {
		fmt.Println("All Bookings:")
		for index, booking := range bookings {
			fmt.Printf("%v. %v %v - %v %v tickets\n Email Address: %v\n", index+1, booking.firstName, booking.lastName, booking.venue, booking.numberOfTickets, booking.email)
		}
	}
}

func sendTicket(data UserData) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", data.numberOfTickets, data.firstName, data.lastName)
	fmt.Println("############################")
	fmt.Printf("Sending ticket:\n%v\nto email address: %v\n", ticket, data.email)
	fmt.Println("############################")
	wg.Done()
}
