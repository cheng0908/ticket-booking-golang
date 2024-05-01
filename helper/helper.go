package helper // belong to main package

import (
	"fmt"
	"sync"
	"time"
)

// we have to explicitly export (public, private or protect) that function so it can be import in another package
// Capital first character means export it. Change greetUsers to GreetUsers to make it available to other packages.

func GreetUsers(confName string) {
	if confName == "" {
		confName = "conference"
	}
	fmt.Printf("Welcome to our \"%v\" booking system\n\n", confName)
}

// Struct object, stands for structure object

type UserDataS struct {
	UserName    string // you need to capitalize the first letter of each field name. This makes them exported (publicly accessible).
	Email       string
	UserTickets uint
}

func SendTicket(userTickets uint, userName string, email string) {
	// simulate sending email
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v\n", userTickets, userName)
	fmt.Println("############")
	fmt.Printf("Sending ticket:\n%v\nto email%v\n\n", ticket, email)
	fmt.Println("############")
	WaitGroup.Done() // routine is finished
}

var WaitGroup sync.WaitGroup
