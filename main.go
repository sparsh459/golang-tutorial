package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceTopic string = "GO Conference"

const conferenceTicket int = 50

var conferenceTicketleft uint = 50

// slice
// {} used to create an empty list of strings
// var bookin []string{}

// slice of empty maps
// var bookin = make([]map[string]string, 0)

// slice of empty struct
var bookin = make([]UserData, 0)

// struct
type UserData struct {
	FirstName    string
	LastName     string
	Email        string
	Age          int
	TicketBought uint
}

// short cut for slice bookin := []string{}

var wg = sync.WaitGroup{}

func main() {
	// Welcoming users
	greetUsers()

	// for conferenceTicketleft > 0 && len(bookin) < 50 {

	// getting user input
	FirstName, LastName, email, age, ticketbought := getUserInput()

	// validating user input
	isvalidName, isvalidEmail, isvalidTicketNumber := helper.ValidateUserInput(FirstName, LastName, email, ticketbought, conferenceTicketleft)

	if isvalidName && isvalidEmail && isvalidTicketNumber {

		// booking tickets
		bookingTickets(FirstName, LastName, ticketbought, age, email)

		wg.Add(1)
		go sendTicket(ticketbought, FirstName, LastName, email)

		// getting list(slice) of first names
		firstNames := getsFirstNames()

		fmt.Printf("The array of the first names: %v\n", firstNames)

		if conferenceTicketleft == 0 {
			// end program
			fmt.Println("No Ticket remaining. Please come next year")
			// break
		}

	} else {
		if !isvalidName {
			fmt.Println("First Name or Last Name you entered is incorrect")
		} else if !isvalidEmail {
			fmt.Println("Email you entered is Incorrect")
		} else if !isvalidTicketNumber {
			fmt.Println("Ticket Number you entered is incorrect")
		}
	}
	// }
	wg.Wait()
}

func greetUsers() {

	fmt.Printf("Types of the conferenceTopic: %T, conferenceTicket: %T, conferenceTicketleft: %T variables assigned.\n", conferenceTopic, conferenceTicket, conferenceTicketleft)
	fmt.Printf("Welcome to %v booking application.\n", conferenceTopic)
	fmt.Printf("Out of total of %v tickets %v are left.\n", conferenceTicket, conferenceTicketleft)
	fmt.Println("Get Your Ticket here.")

}

func getUserInput() (string, string, string, int, uint) {

	// Taking input from user

	var FirstName string
	var LastName string
	var email string
	var age int
	var ticketbought uint

	fmt.Println("Enter First Name:")
	fmt.Scan(&FirstName)

	fmt.Println("Enter Last Name:")
	fmt.Scan(&LastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter Your Age:")
	fmt.Scan(&age)

	fmt.Println("Enter the amount of ticket bought")
	fmt.Scan(&ticketbought)

	return FirstName, LastName, email, age, ticketbought

}

func getsFirstNames() []string {
	// for each loop

	firstNames := []string{}
	// for index, booking := bookin {
	// _ for ignoring variables
	for _, booking := range bookin {

		// used for empty string slices
		// var name = strings.Fields(booking)
		// firstNames = append(firstNames, name[0])

		firstNames = append(firstNames, booking.FirstName)
	}
	return firstNames
}

func bookingTickets(FirstName string, LastName string, ticketbought uint, age int, email string) {
	conferenceTicketleft = conferenceTicketleft - ticketbought

	// fmt.Printf("The Whole slice: %v\n", bookin)
	// fmt.Printf("The len of the slice: %v\n", len(bookin))
	// fmt.Printf("The first value of the slice: %v\n", bookin[0])
	// fmt.Printf("The type of the bookings: %T\n", bookin)

	// fmt.Println("The entered values are::")
	// fmt.Printf("Name: %v, Age: %v, Email: %v, Ticket Bought: %v \n", FirstName+" "+LastName, age, email, ticketbought)

	// for converting a uint value to string
	// strconv.FormatUint(uint64(age), 10)

	// maps addition of data
	// var userData = make(map[string]string)
	// userData["FirstName"] = FirstName
	// userData["LastName"] = LastName
	// userData["Age"] = strconv.FormatUint(uint64(age), 10)
	// userData["Email"] = email
	// userData["Ticket Bought"] = strconv.FormatUint(uint64(ticketbought), 10)

	var userData = UserData{
		FirstName:    FirstName,
		LastName:     LastName,
		Email:        email,
		Age:          age,
		TicketBought: ticketbought,
	}

	bookin = append(bookin, userData)

	fmt.Printf("The Whole slice: %v\n", bookin)

	fmt.Printf("Thank You %v %v for booking the tickets. You will receive a confirmation email at %v\n", FirstName, LastName, email)
	fmt.Printf("Ticket Left: %v\n", conferenceTicketleft)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}

/*

===>> code examples

switch code example

city := "london"

switch city {
case "london":
	// execute code here
case "New York":
	// execute code here
case "Berlin", "Mongolia":
default:
	fmt.Println("No required city selected")
}

array
var bookings [50]string
bookings[0] = FirstName + " " + LastName
bookings[1] = email

fmt.Printf("The Whole Array: %v\n", bookings)
fmt.Printf("The len of the array: %v\n", len(bookings))
fmt.Printf("The first value of the array: %v\n", bookings[0])
fmt.Printf("The type of the bookings: %T\n", bookings)

pointer of a variable
fmt.Println("\nPointer of conferenceTopic")
fmt.Println(&conferenceTopic)

                       "INPUT"  "OUTPUT"
                          |         |
                          |         |
                         \|/       \|/
func printFirstNames([] string) []string

assigning a bool value
var noticket bool = conferenceTicketleft == 0
noticket := conferenceTicketleft == 0

"""break""" if you want to end the program here
"""continue""" if you want user to try again

*/
