package main

import "fmt"

func main() {
	// var showName = "India's Got Latent"
	showName := "India's Got Latent" //syntax sugar but only with variables not possible with "const"
	const showTickets = 100
	var availableTickets = 100

	fmt.Print("Last Episode of the century", showName)
	fmt.Print("Tickets are available", showTickets)
	fmt.Print("Booking Your Seat Now", availableTickets, "Only few are left") // to refactor better you can just have a placeholder %v  whereever you want to use your variables

	var userName string

	fmt.Scan(userName) //------> in order to make this work we need a pointer to the address.
}
