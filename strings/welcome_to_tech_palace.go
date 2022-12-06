package main

import (
	"fmt"
	"strings"
)

var (
	customer     string = "Riens"
	wlcMessage   string = "Welcome!"
	starsPerLine int    = 10
	oldMessage   string = `**************************
*    BUY NOW, SAVE 10%   *
**************************
`
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customer = strings.ToUpper(customer)
	welcome_word := "Welcome to the Tech Palace, " + customer
	return welcome_word
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	composedMessage := stars + "\n" + welcomeMsg + "\n" + stars
	return composedMessage
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanedMessage := strings.Replace(oldMsg, "*", "", -1)
	finalCleanedMessage := strings.TrimSpace(cleanedMessage)
	return finalCleanedMessage
}

func main() {
	fmt.Println(WelcomeMessage(customer))
	fmt.Println(AddBorder(wlcMessage, starsPerLine))
	fmt.Println(CleanupMessage(oldMessage))
}
