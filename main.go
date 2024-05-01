// encapsulate go files into a go package, which is "main". A packages is a collection of Go files.
package main

// import packages
import (
	"booking-app/helper"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Global variables, only available in "main" package
var remainingTickets uint = 50

// when the function name is same as the go package, go compiler will execute this function first. In other hand, in golang, we have to claim the first code or function that go compiler run at first.
func main() {
	fmt.Println("Welcome to our ticket booking system")

	// Define the parameters and variables

	// Method one
	// const conferenceTickets = 50 //constant variable
	// var remainingTickets = 50
	// var conferenceName = "Go Conference"

	// Method Two
	const conferenceTickets int = 50 //constant variable
	// var remainingTickets uint = 50
	var conferenceName string = "Go Conference"

	// define dynamic size of container, as known as "slice"
	var slice_booking = make([]map[string]string, 0) //empty list of map items. '0' is initial size

	// function of fmt package
	fmt.Printf("conferenceName: %T, conferenceTickets: %T, remainingTickets: %T\n", conferenceName, conferenceTickets, remainingTickets) //T stand for Type
	fmt.Println("This is", conferenceName, "booking application")
	fmt.Printf("This is %v booking application", conferenceName) //with specific variable format, f stand for format.
	fmt.Println("we have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	// Storing users information

	var bookings = [50]string{"Alex Lau", "David Ting"} // arrays with fixed size 50; var variable_name [size] variable_type initial_values
	//// adding and accessing value with index
	// bookings[0] = "Florence"

	helper.GreetUsers(conferenceName)

	for {

		userName, email, userTickets := getUserInput(remainingTickets)

		// checking still have enough tickets
		if userTickets > remainingTickets {
			fmt.Printf("\nSorry %v, we only have %v tickets, so you can't book %v tickets.\n", userName, remainingTickets, userTickets)
			continue
		}

		// calculating the remaining tickets
		remainingTickets = remainingTickets - userTickets

		// saving user information to container(array)
		bookings[0] = userName //fixed the sized array, we have to define idx to insert the data
		fmt.Printf("The whole array: %v \nType: %T\n", bookings, bookings)
		fmt.Printf("The first value in array: %v\n", bookings[0])

		// booking tickets
		var userData = make(map[string]string) //only support same type map items. "make" is initial the map item
		userData["userName"] = userName
		userData["email"] = email
		userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		var userDataS = helper.UserDataS{
			UserName:    userName,
			Email:       email,
			UserTickets: userTickets,
		}

		fmt.Printf("Here is %v's userDataS: %v\n", userDataS.UserName, userDataS)

		// updating slice, a dynamic container
		slice_booking = append(slice_booking, userData)
		fmt.Printf("whole slice_booking:%v\n", slice_booking)

		// name mask
		firstNames := getFirstName(slice_booking)

		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		fmt.Printf("User %v has booked %v tickets.\n", userName, userTickets)

		helper.WaitGroup.Add(1) // waits for the launched goroutine to finish
		go helper.SendTicket(userTickets, userName, email)

		fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		// end for loop
		if remainingTickets == 0 {
			fmt.Println("\nOur conference is booked out. Come out next year.")
			break
		}

	}
	helper.WaitGroup.Wait()
	// Switch statement usage
	city := "London"
	switch city {

	case "New York":
		// execute code for booking New York conference
	case "Singapore":
		// execute code for booking New Your conference
	case "London", "Berlin":
		// some code here
	}

}

func getFirstName(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookings { // underscore means non-using variable
		var names = strings.Fields(booking["userName"]) // split the name by space
		firstNames = append(firstNames, names[0])       // get the first name from full name which has been split by Fields.
	}
	return firstNames
}

func getUserInput(remainingTickets uint) (string, string, uint) {
	// var userName // no defined yet, we have to define it at beginning.
	var userName string = "None"
	var email string = "No@email.com"
	var userTickets uint = 0

	// ask user for their name
	// fmt.Print("Enter your name: ")
	// fmt.Scan(&userName) //set username's pointer to new user value. In another way, we assign the memory location to the variable.
	// fmt.Scan() reads input until the first whitespace

	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)    //This creates a new reader that reads from the standard input.
	userName, _ = reader.ReadString('\n')  //This method reads input from the buffer until it encounters a newline character ('\n')
	userName = strings.TrimSpace(userName) // This trims the newline character at the end

	// Format checker for userName
	var isValidName bool = len(userName) >= 2
	if !isValidName {
		fmt.Println("Name format is invalid, please try again")
		return userName, email, userTickets
	}

	fmt.Print("Your email: ")
	fmt.Scan(&email)

	// Format checker for email
	isValidEmail := strings.Contains(email, "@")
	if !isValidEmail {
		fmt.Println("Email format is invalid, please try again")
		return userName, email, userTickets
	}

	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	fmt.Print("How many tickets you want: ")
	fmt.Scan(&userTickets)
	// Convert string to integer
	// number, err := strconv.Atoi(userTickets) // strconv (string conversion) "Atoi" stands for "ASCII to integer."
	// if err != nil {
	// fmt.Println("Invalid input. Please enter a valid integer.")
	// return
	// }
	// Here, 64 specifies the maximum bit size of the uint
	// number, err := strconv.ParseUint(userTickets, 10, 64)
	// if err != nil {
	// 	fmt.Println("Invalid input. Please enter a valid unsigned integer.")
	// 	return
	// }

	return userName, email, userTickets
}

// run command: go run main.go helper.go / go run .
