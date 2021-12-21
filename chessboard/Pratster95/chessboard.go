package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank [8]bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (count int) {
	_, exists := cb[rank]
	if exists {
		for _, square := range cb[rank] {
			if square {
				count += 1
			}
		}

	}
	return

}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (count int) {
	if file >= 1 && file <= 8 {
		for _, rank := range cb {
			if rank[file-1] {
				count += 1
			}
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return len(cb) * 8

}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (count int) {
	for rank := range cb {
		count += CountInRank(cb, rank)
	}
	return
}
