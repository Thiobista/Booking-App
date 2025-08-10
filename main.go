package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)


	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings = make([]UserData, 0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
	isOptedInForNewsletter bool
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()


		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {

           bookTicket(userTickets, firstName, lastName, email)

		   wg.Add(1)
           go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("conference booked out")
				// break
			}

		} else {
			if !isValidName {
				fmt.Println("Your first or last name is too short.")

			}
			if !isValidEmail {
				fmt.Println("Your email address is invalid.")

			}
			if !isValidTickets {
				fmt.Println("Your ticket number is invalid. Please enter a number between 1 and", remainingTickets)

			}
		}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v !\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")
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

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

// map for user

     var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	 }
   
	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n",bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation will be sent to %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v",userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("sending ticket %v \n to email address %v\n", ticket ,email)
    fmt.Println("##########")
	wg.Done()/
}
