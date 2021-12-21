package chessboard

const rankLength = 8

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank [rankLength]bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (numOccupied int) {
	row, exists := cb[rank]
	if exists {
		for _, square := range row {
			if square {
				numOccupied++
			}
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (numOccupied int) {
	if file >= 1 && file <= rankLength {
		for _, rank := range cb {
			if rank[file-1] {
				numOccupied++
			}
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return len(cb) * rankLength
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (numOccupied int) {
	for file := 1; file <= rankLength; file++ {
		numOccupied += CountInFile(cb, file)
	}
	return
}
