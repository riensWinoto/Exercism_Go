package main

import "fmt"

var myRhyme []string = []string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"}

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var myProverb []string
	lowestLimit, highestLimit := 0, len(rhyme)-1
	for ; lowestLimit <= highestLimit; lowestLimit++ {
		if lowestLimit == highestLimit {
			proverbFinalLineWords := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
			myProverb = append(myProverb, proverbFinalLineWords)
		} else {
			proverbWords := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[lowestLimit], rhyme[lowestLimit+1])
			myProverb = append(myProverb, proverbWords)
		}
	}
	return myProverb
}

func main() {
	fmt.Println(Proverb(myRhyme))
}
