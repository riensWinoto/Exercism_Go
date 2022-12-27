package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var (
	myLog                string = "🔍 recommended search product 🔍"
	myOldRune, myNewRune rune   = '🔍', '❗'
	logLimit             int    = 10
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var (
		tempHolderSlice  []rune
		tempHolderString string = "default"
	)

	for _, val := range log {
		if string(val) == "❗" || string(val) == "🔍" || string(val) == "☀" {
			tempHolderSlice = append(tempHolderSlice, val)
		}
	}

	if len(tempHolderSlice) != 0 {
		switch string(tempHolderSlice[0]) {
		case "❗":
			tempHolderString = "recommendation"
			break
		case "🔍":
			tempHolderString = "search"
			break
		case "☀":
			tempHolderString = "weather"
			break
		}
	}
	return tempHolderString
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(Application(myLog))
	fmt.Println(Replace(myLog, myOldRune, myNewRune))
	fmt.Println(WithinLimit(myLog, logLimit))
}
