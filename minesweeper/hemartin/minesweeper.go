package minesweeper

import (
	"fmt"
	"strings"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {
	if len(board) == 0 {
		return []string{}
	}

	if board[0] == "" {
		return []string{""}
	}

	intermediate := make([][]int, len(board))
	for i := range intermediate {
		intermediate[i] = make([]int, len(board[0]))
	}

	for rowPos, row := range board {
		for colPos, col := range row {
			if string(col) == "*" {
				intermediate[rowPos][colPos] = -9

				prevRowPos, nextRowPos := rowPos-1, rowPos+1
				prevColPos, nextColPos := colPos-1, colPos+1

				hasPrevRow := prevRowPos >= 0
				hasNextRow := nextRowPos <= len(board)-1
				hasPrevCol := prevColPos >= 0
				hasNextCol := nextColPos <= len(row)-1

				if hasPrevCol {
					intermediate[rowPos][prevColPos] += 1
				}

				if hasNextCol {
					intermediate[rowPos][nextColPos] += 1
				}

				if hasPrevRow {
					intermediate[prevRowPos][colPos] += 1

					if hasPrevCol {
						intermediate[prevRowPos][prevColPos] += 1
					}

					if hasNextCol {
						intermediate[prevRowPos][nextColPos] += 1
					}

				}

				if hasNextRow {
					intermediate[nextRowPos][colPos] += 1

					if hasPrevCol {
						intermediate[nextRowPos][prevColPos] += 1
					}

					if hasNextCol {
						intermediate[nextRowPos][nextColPos] += 1
					}
				}
			}

		}
	}

	annotated := make([]string, len(board))
	for rowPos, row := range intermediate {
		var b strings.Builder
		b.Grow(len(row))
		for _, col := range row {
			if col < 0 {
				b.WriteString("*")
			} else if col == 0 {
				b.WriteString(" ")
			} else {
				b.WriteString(fmt.Sprint(col))
			}
			annotated[rowPos] = b.String()
		}
	}

	return annotated
}
