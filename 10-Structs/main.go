package main

import (
	"fmt"
	"strings"
)

var bookings = make([]UserData, 0)

type UserData struct {
	firstname string
	lastname string
	email string
	tickets int
}

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	// Greeting User
	greetUser(conferenceName, conferenceTickets, remainingTickets)

	// for {
		var userFirstName string
		var userLastName string
		var userEmail string
		var userTickets int
		var selectCity int

		// Validating User Name
		userFirstName, userLastName = validateName(userFirstName, userLastName)
	
		// Validating User Email
		userEmail = validateEmail(userEmail)

		// Validate Tickets
		userTickets = validateTickets(userTickets, remainingTickets)
	
		// Validate City
		validateCity(selectCity)

		// Remaining Tickets
		remainingTickets -= userTickets
		
		var userData = UserData {
			firstname: userFirstName,
			lastname: userLastName, 
			email: userEmail, 
			tickets: userTickets,
		}
		
		bookings = append(bookings, userData)
		fmt.Printf("List of bookings is %v.\n", bookings)
		
		// Booking Confirmation
		bookTicket(userFirstName, userLastName, userEmail, userTickets, remainingTickets, conferenceName)

		// Sold Tickets
		soldTickets(remainingTickets)
	// }
}

func greetUser(conferenceName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("\nWelcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend...\n")
}

func validateName(userFirstName string, userLastName string) (string, string) {
	for {
			fmt.Print("Enter your first name: ")
			fmt.Scan(&userFirstName)
		
			fmt.Print("Enter your last name: ")
			fmt.Scan(&userLastName)
	
			isValidName := len(userFirstName) > 2 && len(userLastName) > 2	

			if !isValidName {
				fmt.Printf("\n\nfirst name or last name you entered is too short!!!\n")
				continue
			} else {
				break
			}
	}
	return userFirstName, userLastName
}

func validateEmail(userEmail string) (string){
	for {
			fmt.Print("Enter your email: ")
			fmt.Scan(&userEmail)
	
			isValidEmail := strings.Contains(userEmail, "@")

			if !isValidEmail {
				fmt.Printf("\n\nemail address you entered doesn't contain @ sign!!!\n")
				continue
			} else {
				break
			}
	}
	return userEmail
}

func validateTickets(userTickets int, remainingTickets int) (int) {
	for {
			fmt.Print("How many tickets You want to buy? \n")
			fmt.Scan(&userTickets)

			isValidTicket := userTickets > remainingTickets || userTickets < 1
			
			if isValidTicket { 
				fmt.Printf("\n\nnumber of tickets you entered is invalid!!!\n")
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
				continue
			} else {
				break
			}
	}
	return userTickets
}

func validateCity(selectCity int) {
	for {
			fmt.Println("Select City: 1-London, 2-Mumbai, 3-New York: ")
			fmt.Scan(&selectCity)

			if (selectCity >= 1 && selectCity <= 3) {
				switch selectCity {
					case 1: 
								fmt.Println("You have chosen London City.")
								break
					case 2: 
								fmt.Println("You have chosen Mumbai City.")
								break
					case 3: 
								fmt.Println("You have chosen New York City.")
								break
					default: 
								fmt.Println("Invalid City!!!")
				}
				break
			} else {
				fmt.Println("Please enter valid no between 1 and 3.")
				continue
			}
	}
}

func bookTicket(
	userFirstName string,
	userLastName string,
	userEmail string,
	userTickets int,
	remainingTickets int,
	conferenceName string,
) {
	fmt.Printf("Thank you %v %v for booking %v tickets.\n", userFirstName, userLastName, userTickets)
	fmt.Printf("You will receive a confirmation email at %v\n", userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func soldTickets(remainingTickets int) {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstname)
	}
	fmt.Println("These are all bookings:", firstNames)

	if remainingTickets == 0 {
		fmt.Println("Our conference is booked out. Come back next year.")
	}
}