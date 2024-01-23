package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Go converence"
var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers(conferenceName, conferenceTickets)

	for remainingTickets > 0 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("Booked first names %v \n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name of last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered is invalid.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
			continue
		}
	}
	wg.Wait()
}

func greetUsers(conferenceName string, conferenceTickets uint) {
	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thnak you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v. \n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n%vto email address: %v\n", ticket, email)
	fmt.Println("######################")
	wg.Done()
}
