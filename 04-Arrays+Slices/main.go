package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}
	
	fmt.Printf("\nWelcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend...\n")

	var userFirstName string = "Tom"
	var userLastName string = "Bom"
	var userEmail string = "tom@gmail.com"
	var userTickets uint = 4

	// fmt.Print("Enter your first name: ")
	// fmt.Scan(&userFirstName)

	// fmt.Print("Enter your last name: ")
	// fmt.Scan(&userLastName)

	// fmt.Print("Enter your email: ")
	// fmt.Scan(&userEmail)

	// fmt.Print("How many tickets You want to buy? ")
	// fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	bookings = append(bookings, userFirstName + " " + userLastName)
	
	fmt.Printf("\nThank you %v %v for booking %v tickets.\n", userFirstName, userLastName, userTickets)
	fmt.Printf("You will receive a confirmation email at %v\n", userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Println("These are all bookings:", bookings)
}