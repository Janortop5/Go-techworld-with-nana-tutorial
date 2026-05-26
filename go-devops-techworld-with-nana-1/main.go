package main 

import (
	"fmt"
	"time"
	"sync"
)

	// var conferenceName1 string = "Go Conference"
var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

greetUsers()

//	fmt.Printf("conferenceTickets is %T, remainiingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

//	fmt.Printf("Welcome to %v booking application\nWe have total of %v tickets and %v are still available\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	for { // remainingTickets > 0
			firstName, lastName, email, userTickets := getUserInput()
			isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
			// isValidCity := city == "singapore" || city == "London"
			// isInvalidCity := city != "Singapore" && city != "London"

			if  isValidName && isValidEmail && isValidTicketNumber { // userTickets < remainingTickets
					
				  bookTicket(userTickets, firstName, lastName, email)
					
					wg.Add(1)
					go sendTicket(userTickets, firstName, lastName, email)
			

					// call function print first name
					firstNames := getFirstNames()
					fmt.Printf("The first name of bookings are %v\n", firstNames)



					if  remainingTickets == 0 {
						// end program
						fmt.Println("Our conference is booked out. Come back next year.")
						break
					}
			} else {
				  if !isValidName { 
							fmt.Println("first name or last name you entered is too short")
					} 
					if !isValidEmail { // else if
						fmt.Println("email address you entered doesn't contain @ sign")
					} 
					if !isValidTicketNumber { // else if
						fmt.Println("number of tickets you entered is invalid")
					}
					// fmt.Println("Your input data is invalid, try again")
					// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			} 
		 wg.Wait()
	}

/*	city := "London"

	switch city {
			case "New York":
					// execute code for booking New York conference tickets
			case "Singapore", "Hong Kong":
					// execute code for booking Singapore conference tickets
			// case "London":
					// some code here
			case "London", "Berlin":
					// some code here for London & Berlin
			// case "Hong Kong":
					// some code here
			case "Mexico City":
					// some code here
			default:
					fmt.Print("No valid city selected")
	}
}*/
}

func greetUsers() {
		fmt.Printf("Welcome to %v booking application\n", conferenceName)
		fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
		firstNames := []string{}
		for _, booking := range bookings {
				// firstNames = append(firstNames, booking["firstName"])
				firstNames = append(firstNames, booking.firstName)
		}
		return firstNames 
}

func getUserInput() (string, string, string, uint) {
			var firstName string
			var userTickets uint
			var lastName string
			var email string
			// ask user for ther name

			fmt.Println("Enter your first name: ")
			fmt.Scan(&firstName)

			fmt.Println("Enter your last name")
			fmt.Scan(&lastName)

			fmt.Println("Enter your email address")
			fmt.Scan(&email)

			fmt.Println("Enter number of tickets")
			fmt.Scan(&userTickets)
			return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
			remainingTickets = remainingTickets - userTickets
			
			// create a map for a user
			// var myslice []string
			// var mymap = map[string]string

			// var userData = make(map[string]string)
			// userData["a"] = firstName
			// userData["b"] = lastName
			//userData["c"] = email

			var userData = UserData {
				firstName: firstName,
				lastName: lastName,
				email: email,
				numberOfTickets: userTickets,
			}
			
			//userData["firstName"] = firstName
			// userData["lastName"] = lastName
			// userData["email"] = email
      // userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
			
			bookings = append(bookings, userData)

			fmt.Printf("List of bookings is %v\n", bookings)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
		// time.Sleep(10 * time.Second)
		time.Sleep(50 * time.Second)
		var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
		fmt.Println("#####################")
		fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
		fmt.Println("#####################")
		wg.Done()
}