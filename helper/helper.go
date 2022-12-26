package helper

import "strings"

func ValidateUserInput(FirstName string, LastName string, email string, ticketbought uint, conferenceTicketleft uint) (bool, bool, bool) {

	isvalidName := len(FirstName) >= 2 && len(LastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTicketNumber := ticketbought > 0 && ticketbought <= conferenceTicketleft

	return isvalidName, isvalidEmail, isvalidTicketNumber

}
