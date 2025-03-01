package main

import (
	"fmt"
	"strings"
)

const totalTickets uint = 100

var showName string = "India's Got Latent"
var availableTickets uint = 100
var bookings = []string{}
var firstName, lastName, email string
var userTickets uint

func main() {

	greet()

	// Booking loop
	for {

		firstName, lastName, email, userTickets = userDetail()

		// Validation checks
		isValidName, isValidEmail, isValidTicketNum := validator(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNum {
			// Process booking
			booky() // ✅ Fixed (no arguments needed)

			// Display list of attendees (only first names)
			firstNames := attendees()
			fmt.Printf("📜 List of attendees: %v\n", firstNames)

			// If tickets are sold out, end the booking process
			if availableTickets == 0 {
				fmt.Println("\n🚨 The show is housefull! Thanks for your interest.")
				break
			}
		} else {
			// Handle invalid input
			fmt.Printf("\n❌ Invalid input. Please check:\n")
			if !isValidName {
				fmt.Println("- Your first and last names must have at least 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("- Please enter a valid email address.")
			}
			if !isValidTicketNum {
				fmt.Printf("- Only %v seats are available. Enter a valid number.\n", availableTickets)
			}
		}
	}
}

// ✅ FIXED: No parameters needed, using global variables directly
func greet() {
	fmt.Printf("🎭 Welcome to the last episode of the century: %v!\n", showName)
	fmt.Printf("📢 Book your seats now! Only %v tickets left.\n", totalTickets)
}

// Returns the first names of attendees
func attendees() []string {
	var firstNames []string
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func userDetail() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	// User input
	fmt.Print("\nEnter Your First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter Your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter Your Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter Number of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// ✅ FIXED: No parameters needed, using global variables directly
func booky() {
	availableTickets = availableTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	// Confirmation message
	fmt.Printf("\n✅ Thank you, %v, for booking %v ticket(s)! for %v. A confirmation email will be sent to %v.\n", firstName, userTickets, showName, email)
	fmt.Printf("⚡ Hurry up! Only %v tickets left.\n", availableTickets)
}
