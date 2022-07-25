package minesweeper

import (
	"strconv"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {
	boardHeight := len(board)
	if (boardHeight < 1) {
		return nil
	}

	boardWidth := len(board[0])
	var countingBoard [][]rune
	var finalBoard = make([]string, boardHeight)
	
	for _, s := range board {
		var countingRow []rune
		for _, char := range s {
			countingRow = append(countingRow, char)
		}
		countingBoard = append(countingBoard, countingRow)
	}

	for row := 0; row < boardHeight; row++ {
		var newRow string
		for column := 0; column < boardWidth; column++ {
			if (countingBoard[row][column] == '*') {
				newRow += "*"
			} else {
				adjacentMines := 0

				for i := row-1; i <= row+1; i++ {
					for j := column-1; j <= column+1; j++ {
						if (i>=0) && (i<boardHeight) && (j>=0) && (j<boardWidth) {
							if (countingBoard[i][j] == '*') {
								adjacentMines++
							}
						}
					}
				}

				if (adjacentMines == 0) {
					newRow += " "
				} else {
					newRow += strconv.Itoa(adjacentMines)
				}
			}
		}
		finalBoard[row] = newRow
	}

	return finalBoard
}
