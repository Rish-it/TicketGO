package main

import "fmt"

func main() {
	// var showName = "India's Got Latent"
	showName := "India's Got Latent" //syntax sugar but only with variables not possible with "const"
	const showTickets = 100
	var availableTickets uint = 100
	var bookings []string

	fmt.Println("Last Episode of the century", showName)
	fmt.Println("Tickets are available", showTickets)
	fmt.Println("Booking Your Seat Now", availableTickets, "Only few are left") // to refactor better you can just have a placeholder %v  whereever you want to use your variables

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your Email:")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets:")
	fmt.Scan(&userTickets) //------> in order to make this work we need a pointer to the address.

	availableTickets = availableTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName) //--- using slice for dynamic list mng

	fmt.Printf("Thankyou %v for choosing us and booking %v tickets. You will receive a confirmation email at %v\n", firstName, userTickets, email)

	fmt.Printf("Book fast only few are left %v\n", availableTickets)

	fmt.Printf("List of Attendees %v", bookings)
}
