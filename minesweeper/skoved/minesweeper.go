package minesweeper

import (
	"fmt"
	"strings"
)

const mine = '*'

type tile struct {
	mine      bool
	mineCount uint
}

func (t tile) String() string {
	if t.mine {
		return string(mine)
	}
	if t.mineCount == 0 {
		return " "
	}
	return fmt.Sprintf("%d", t.mineCount)
}

// increments the mineCount of the tile if the tile does not have a mine
func (t *tile) incMineCount() {
	if !t.mine {
		t.mineCount++
	}
}

// return a sequence of ints from start to end. Sequence will always start at zero and end at length - 1
func makeRange(start, end, length int) []int {
	if start < 0 {
		start = 0
	}
	if end >= length {
		end = length - 1
	}
	if start > end {
		panic("Start cannot be greater than end")
	}

	newRange := make([]int, end-start+1)
	for i := range newRange {
		newRange[i] = start + i
	}
	return newRange
}

func incTiles(board [][]tile, rowRange, colRange []int) {
	for _, rowNum := range rowRange {
		for _, colNum := range colRange {
			board[rowNum][colNum].incMineCount()
		}
	}
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	length := len(board)

	// short circuit for empty boards
	if length == 0 {
		return []string{}
	}
	if board[0] == "" {
		return []string{""}
	}

	// create board with tiles
	transform := make([][]tile, length)
	for i, row := range board {
		transform[i] = make([]tile, len(row))
	}

	// add mines to new board
	for rowNum, row := range board {
		for colNum, space := range row {
			if space == mine {
				transform[rowNum][colNum].mine = true
			}
		}
	}

	for rowNum, row := range board {
		for colNum, space := range row {
			if space == mine {
				rowRange := makeRange(rowNum-1, rowNum+1, length)
				colRange := makeRange(colNum-1, colNum+1, len(row))
				incTiles(transform, rowRange, colRange)
			}
		}
	}

	// convert transform to []string
	markedBoard := make([]string, length)
	for rowNum, row := range transform {
		var builder strings.Builder
		builder.Grow(len(row))
		for _, space := range row {
			builder.WriteString(fmt.Sprintf("%v", space))
		}
		markedBoard[rowNum] = builder.String()
	}
	return markedBoard
}
