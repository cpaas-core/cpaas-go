package minesweeper

import "fmt"

func GetColumnNumber(board []string) int {
	return len(board[0])
}

func GetRowNumber(board []string) int {
	return len(board)
}

// Check if a position is cell in the board
func IsACell(row int, col int, board []string) bool {
	rowSize := GetRowNumber(board)
	colSize := GetColumnNumber(board)
	if (row > 0) && (row <= rowSize) && (col > 0) && (col <= colSize) {
		return true
	}
	return false
}

// Check whether given cell (string) has a mine or not.
func IsAMine(row int, col int, board []string) bool {
	stringRow := board[row-1]
	possibleMine := []rune(stringRow)[col-1]
	// rune 42 == *
	return possibleMine == 42
}

// Count the number of mines in the adjacent cells
func CountAdjacentMines(row int, col int, board []string) int {
	/*
	   Count all the mines in the 8 adjacent cells

	   (1,1) (1,2) (1,3)
	   	\   	|   /
	   (2,1)-(2,2)-(2,3)
	   	/  		|   \
	   (3,1) (3,2) (3,3)

	   (1,1) (row-1, col-1)
	   (1,2) (row-1, col)
	   (1,3) (row-1, col+1)
	   (2,1) (row,   col-1)
	   (2,2) (row,   col) -> do not check
	   (2,3) (row,   col+1)
	   (3,1) (row+1, col-1)
	   (3,2) (row+1, col)
	   (3,3) (row+1, col+1)
	*/
	adjacentMines := 0
	if IsACell(row-1, col-1, board) {
		if IsAMine(row-1, col-1, board) {
			adjacentMines++
		}
	}
	if IsACell(row-1, col, board) {
		if IsAMine(row-1, col, board) {
			adjacentMines++
		}
	}
	if IsACell(row-1, col+1, board) {
		if IsAMine(row-1, col+1, board) {
			adjacentMines++
		}
	}
	if IsACell(row, col-1, board) {
		if IsAMine(row, col-1, board) {
			adjacentMines++
		}
	}
	if IsACell(row, col+1, board) {
		if IsAMine(row, col+1, board) {
			adjacentMines++
		}
	}
	if IsACell(row+1, col-1, board) {
		if IsAMine(row+1, col-1, board) {
			adjacentMines++
		}
	}
	if IsACell(row+1, col, board) {
		if IsAMine(row+1, col, board) {
			adjacentMines++
		}
	}
	if IsACell(row+1, col+1, board) {
		if IsAMine(row+1, col+1, board) {
			adjacentMines++
		}
	}
	return adjacentMines
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	row := 0
	col := 0
	mines := 0
	var minesInBoard []string

	if len(board) == 0 {
		return board
	}

	if board[0] == "" {
		return board
	}

	for _, rowBoard := range board {
		row++
		rowString := ""
		for indexCol, cell := range rowBoard {
			col++
			if cell == 42 {
				rowString = rowString + "*"
			} else {
				mines = CountAdjacentMines(row, col, board)
				if mines == 0 {
					rowString = rowString + " "
				} else {
					rowString = rowString + fmt.Sprint(mines)
				}
			}
			if indexCol+1 == GetColumnNumber(board) {
				minesInBoard = append(minesInBoard, rowString)
				col = 0
			}
		}
	}
	return minesInBoard
}
