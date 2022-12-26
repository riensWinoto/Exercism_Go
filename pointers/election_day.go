package main

import "fmt"

type ElectionResult struct {
	Name  string
	Votes int
}

var (
	initVote, incrementVote int            = 2, 3
	voteCounter             *int           = &initVote
	candidateName           string         = "Daendels"
	voteResult              map[string]int = map[string]int{candidateName: initVote}
)

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	candidateVotes := &ElectionResult{Name: candidateName, Votes: votes}
	return candidateVotes
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%v (%v)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}

func main() {
	fmt.Println(NewVoteCounter(initVote))
	fmt.Println(VoteCount(voteCounter))
	IncrementVoteCount(voteCounter, incrementVote)
	fmt.Println(NewElectionResult(candidateName, initVote))
	electRes := NewElectionResult(candidateName, initVote)
	fmt.Println(DisplayResult(electRes))
	DecrementVotesOfCandidate(voteResult, candidateName)
}
