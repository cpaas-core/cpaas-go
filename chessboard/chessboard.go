package chessboard

type Rank []bool

type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (count int) {
	for i := range cb[rank]{
		if cb[rank][i]{count += 1}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (count int) {
	if file < 0 || file>len(cb){
		return 0
	}
	for i := range cb {
		if cb[i][file-1]{count += 1}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (count int) {
	for _,r := range cb{
		for range r{
			count +=1
		}
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (count int) {
	for _,r := range cb{
		for v := range r{
			if r[v] {
				count +=1
			}
		}
	}
	return
}
