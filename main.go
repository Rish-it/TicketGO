package main

import "fmt"

func main() {
	var showName = "India's Got Latent"
	const showTickets = 100
	var availableTickets = 100
	fmt.Print("Last Episode of the century", showName)
	fmt.Print("Tickets are available", showTickets)
	fmt.Print("Booking Your Seat Now", availableTickets, "Only few are left") // to refactor better you can just have a placeholder %v  whereever you want to use your variables

}
