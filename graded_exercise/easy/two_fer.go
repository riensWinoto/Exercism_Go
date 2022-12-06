package main

import "fmt"

var coolName string = "Riens"

// ShareWith function will return "You" if its name parameter blank string.
// Otherwise it will return based on name parameter value
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
}

func main() {
	fmt.Print(ShareWith(coolName))
}
