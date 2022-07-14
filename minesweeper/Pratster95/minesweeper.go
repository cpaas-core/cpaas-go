package minesweeper

import "strconv"

// Annotate returns an annotated board
func Annotate(board []string) []string {

	minesweeper := make([]string, len(board))

	for r := 0; r < len(board); r++ {
		minesweeper[r] = ""
		for c := 0; c < len(board[r]); c++ {
			if board[r][c] == '*' {
				minesweeper[r] += "*"
			} else {
				totalMines := mines_count(board, r, c)
				if totalMines > 0 {
					minesweeper[r] += strconv.Itoa(totalMines)
				} else {
					minesweeper[r] += " "
				}
			}
		}
	}
	return minesweeper
}

func mines_count(board []string, r, c int) int {
	count := 0
	for new_row := r - 1; new_row <= r+1; new_row++ {
		for new_col := c - 1; new_col <= c+1; new_col++ {

			if new_row >= 0 && new_row < len(board) && new_col >= 0 && new_col < len(board[new_row]) && board[new_row][new_col] == '*' {

				count += 1
			}
		}
	}
	return count
}
