package main

import "fmt"

func main() {
	conferenceName := "Best Conference"
	const conferenceTickets = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("We have total off %v tickets and still aviable: %v \n", conferenceTickets, remainingTickets)
	fmt.Println("-- Get your tickets here --")

	var bookings [50]string

	var userName string

	//Get Username
	fmt.Println("What is you`re first name?")
	fmt.Scan(&userName)

	bookings[0] = userName

	fmt.Printf("Hello %v", userName)
}
