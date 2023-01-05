package main

import (
	"fmt"
	"strings"
)

func Abbreviate(s string) string {
	temp := []rune{}

	replaceString := strings.Replace(s, "_", " ", -1)
	replaceString = strings.Replace(replaceString, "-", " ", -1)
	sliceString := strings.Fields(replaceString)
	for _, val := range sliceString {
		temp = append(temp, rune(val[0]))
	}
	return strings.ToUpper(string(temp))
}

func main() {
	fmt.Println(Abbreviate("The Road _not- Taken"))
}
