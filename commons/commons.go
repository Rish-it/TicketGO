package commons

import "strings"

// Validates the user input
func Validator(firstName, lastName, email string, userTickets uint, availableTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNum := userTickets > 0 && userTickets <= availableTickets

	return isValidName, isValidEmail, isValidTicketNum
}
