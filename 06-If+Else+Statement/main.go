package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	bookings := []string{}
	
	fmt.Printf("\nWelcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend...\n")

	for {
		var userFirstName string
		var userLastName string
		var userEmail string
		var userTickets int

		// Validating User Name
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
	
		// Validating User Email
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

		// Validate Tickets
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
	
		remainingTickets -= userTickets
		bookings = append(bookings, userFirstName + " " + userLastName)
		
		fmt.Printf("Thank you %v %v for booking %v tickets.\n", userFirstName, userLastName, userTickets)
		fmt.Printf("You will receive a confirmation email at %v\n", userEmail)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Println("These are all bookings:", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}
}