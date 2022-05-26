package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var pointerToInitialVotes *int
	pointerToInitialVotes = &initialVotes
	return pointerToInitialVotes
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
	var electionResult ElectionResult
	electionResult = ElectionResult{Name: candidateName, Votes: votes}
	var pointerToElectionResult *ElectionResult
	pointerToElectionResult = &electionResult
	return pointerToElectionResult
}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	var message string
	message = result.Name + " (" + fmt.Sprint(result.Votes) + ")"
	return message
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	for key := range results {
		if key == candidate {
			results[key] = results[key] - 1
		}
	}
}
