package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, email, userTickets)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The firstnames of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Sorry, we are sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Print("Sorry, your name is invalid. Please try again!\n")
			}
			if !isValidEmail {
				fmt.Print("Sorry, your email is invalid. Please try again!\n")
			}
			if !isValidTicketNumber {
				fmt.Print("Sorry, your ticket number is invalid. Please try again!\n")
			}
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your", conferenceName, "tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name:")
	fmt.Scanln(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scanln(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scanln(&email)

	fmt.Println("Please enter number of tickets you want to book")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets! You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets are now remaining: %v\n", remainingTickets)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")
}
