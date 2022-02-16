package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {

	countInRank := 0

	rankInChessboard, exists := cb[rank]
	if exists {
		for _, occupied := range rankInChessboard {
			if occupied {
				countInRank = countInRank + 1
			}
		}
	}
	return countInRank
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {

	countInFile := 0

	if (9 > file) && (file > 0) {
		for _, rank := range cb {
			for index, squareOccupied := range rank {
				if (index + 1) == file {
					if squareOccupied {
						countInFile = countInFile + 1
					}
				}
			}
		}
	}
	return countInFile
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {

	countInBoard := 0

	for _, rank := range cb {
		for _, square := range rank {
			if square || !square {
				countInBoard = countInBoard + 1
			}
		}
	}
	return countInBoard
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {

	countInBoard := 0

	for _, rank := range cb {
		for _, square := range rank {
			if square {
				countInBoard = countInBoard + 1
			}
		}
	}
	return countInBoard
}
