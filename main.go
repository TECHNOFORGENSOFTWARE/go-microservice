package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// var bookings []string
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T,conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Helo to our ", conferenceName, "Booking application")

	fmt.Printf("Welcome to  %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v Tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get yooour tickets to attend")

	// var bookings = [50]string{"nithin","prakash","peter"}
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// var userTickets int
		//ask user for thier name
		// fmt.Println(remainingTickets)
		// fmt.Println(&remainingTickets)
		fmt.Println("Enter your first Name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last Name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		// isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// isValideEmail := strings.Contains(email, "@")
		if userTickets > remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("First username is  %v, Last user name is  %v booked the tickets with email with %v\n", firstName, lastName, email)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// noticketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				//end program

				fmt.Println("Our confernece is booked out and come back next year")
				break
			}
		}
		fmt.Printf("we have only %v tickets remaning. soyou cant bookmore than %v tickets", remainingTickets, userTickets)
		continue

		// fmt.Printf("The whole array : %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Slice type:%T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

	}
}
