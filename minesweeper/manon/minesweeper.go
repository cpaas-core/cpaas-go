package minesweeper

import "strconv"

// Annotate returns an annotated board
func Annotate(board []string) []string {
	var annotated []string
	// for each row
	for rowIdx, row := range board {
		var annotatedRow string
		// for each col
		for colIdx, char := range row {
			var mines int
			// if there is a space
			if char == ' ' {
				// count the surrounding mines
				mines = countMines(board, rowIdx, colIdx)
			}
			if mines > 0 {
				// if there are surrounding mines, the position must be replaced by the number
				annotatedRow += strconv.Itoa(mines)
			} else {
				// if not, we should keep the mine or the space
				annotatedRow += string(char)
			}
		}
		annotated = append(annotated, annotatedRow)
	}
	return annotated
}

// countMines count the number of mines `*` surrounding the position rowIdx,colIdx in the given board
func countMines(board []string, rowIdx, colIdx int) int {
	var mines int
	// check the rows from rowIdx-1 to rowIdx+1 without getting out of the range of the board
	for rowPosition := max(rowIdx-1, 0); rowPosition <= min(rowIdx+1, len(board)-1); rowPosition++ {
		// check the cols from colIdx-1 to colIdx+1 without getting out of the range of the board
		for colPosition := max(colIdx-1, 0); colPosition <= min(colIdx+1, len(board[0])-1); colPosition++ {
			if rowPosition == rowIdx && colPosition == colIdx {
				continue
			}
			element := board[rowPosition][colPosition]
			// if there is a mine
			if element == '*' {
				// increment the counter
				mines++
			}
		}
	}
	return mines
}
