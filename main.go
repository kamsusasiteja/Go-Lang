// The very first progran for Go Lang
// Program and structure of Go File
//go mod init booking-app     gives a go.mod file that consists of module booking-app & go 1.19
// variables and constants in Go
//Formatted output- Printf
// data types in Go
//Getting user input in Go
//a small logic for book-ticket




package main
import "fmt" //required packages

//program starts here
func main(){

	var conferenceName= "Go Conference"
	// alternate declaration like above, 	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 

	fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName) // %T indicates type of value

	fmt.Println("Welcome to",conferenceName,"booking application!")
	fmt.Println("We have a total of",conferenceTickets,"tickets and",remainingTickets,"are still available")
	fmt.Println("Get your conference tickets here!")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets that you want to book ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email on %v\n",firstName,lastName,userTickets,email )
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)


	// userName ="Tom"
	// userTickets = 2  
	// fmt.Printf("User %v booked %v ticket(s)\n",userName,userTickets)
	

	
	// fmt.Println(conferenceName)
}
