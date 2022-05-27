package electionday

import "fmt"

var votePointer *int
var voteCounter int


// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	votePointer = &voteCounter
	*votePointer = initialVotes
	return votePointer
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	var votes int
	if counter == nil {
		votes = 0
	} else {
		votes = *counter
	}
	return votes
}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {

	var result ElectionResult
	var newResult *ElectionResult

	result.Name = candidateName
	result.Votes = votes
	newResult = &result
	
	return newResult
}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	message := fmt.Sprintf("%s (%d)", result.Name, result.Votes)
	return message
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] = results[candidate]-1
}
