package main

import "fmt"

var (
	originWord string = "Hello 世界, Awesome Go Language"
)

func Reverse(input string) string {
	var tempHolder, finalHolder []rune
	tempHolder = []rune(input)

	for i := len(tempHolder) - 1; i >= 0; i-- {
		finalHolder = append(finalHolder, tempHolder[i])
	}
	return string(finalHolder)
}

func main() {
	fmt.Println(Reverse(originWord))
}
