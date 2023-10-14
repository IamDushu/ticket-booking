package main

import (
	"fmt"
	"strings"
)

// Package Level Variables
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

		} else {
			if !isValidName {
				fmt.Println("Your first name or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Your entered email is invalid, doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("No.of Tickets entered is invalid")
			}
		}
	}
	fmt.Println("Our Conference is sold out!. Come back next year.")
}

func greetUsers() {
	fmt.Printf("Welcome to our %v Booking Application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking) //"Dushyanth G" -> ["Dushyanth","G"]
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}
