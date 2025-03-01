package main

import (
	"fmt"
	"strings"
)

func main() {
	// Show details
	showName := "India's Got Latent" // Syntax sugar, but not possible with "const"
	const totalTickets = 100
	var availableTickets uint = 100
	var bookings []string

	fmt.Printf("🎭 Welcome to the last episode of the century: %v!\n", showName)
	fmt.Printf("🎟️ Total tickets available: %v\n", totalTickets)
	fmt.Printf("📢 Book your seats now! Only %v tickets left.\n", availableTickets)

	// Booking loop
	for {
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

		// Validation checks
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNum := userTickets > 0 && userTickets <= availableTickets

		if isValidName && isValidEmail && isValidTicketNum {
			// Process booking
			availableTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			// Confirmation message
			fmt.Printf("\n✅ Thank you, %v, for booking %v ticket(s)! A confirmation email will be sent to %v.\n", firstName, userTickets, email)
			fmt.Printf("⚡ Hurry up! Only %v tickets left.\n", availableTickets)

			// Display list of attendees (only first names)
			var firstNames []string
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
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
