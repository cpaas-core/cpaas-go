package minesweeper

import (
	"container/list"
	"strconv"
)

type point struct {
	x, y int
}

func newPoint(x, y int) *point {
	p := point{x: x, y: y}
	return &p
}

func (p point) getBorders(thresholdX int, thresholdY int) (point, point) {
	leftX := 0
	topY := 0
	rightX := thresholdX
	bottomY := thresholdY

	if p.x-1 > 0 {
		leftX = p.x - 1
	}
	if p.x+1 < thresholdX {
		rightX = p.x + 1
	}
	if p.y-1 > 0 {
		topY = p.y - 1
	}
	if p.y+1 < thresholdY {
		bottomY = p.y + 1
	}

	return *newPoint(leftX, topY),
		*newPoint(rightX, bottomY)
}

func arrayBackToBoard(array [][]rune) []string {
	board := make([]string, len(array))
	for i, r := range array {
		for _, v := range r {
			board[i] += string(v)
		}
	}
	return board
}

func boardToArrayAndList(board []string) ([][]rune, *list.List) {
	var boardArr = make([][]rune, len(board))
	pointList := list.New()
	for i, s := range board {
		boardArr[i] = make([]rune, len(s))
		for j, r := range s {
			if string(r) == "*" {
				p := newPoint(i, j)
				pointList.PushFront(p)
			}
			boardArr[i][j] = r
		}
	}
	return boardArr, pointList
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	if len(board) == 0 {
		return board
	}
	if len(board[0]) == 0 {
		return board
	}

	boardArr, pointList := boardToArrayAndList(board)
	for e := pointList.Front(); e != nil; e = e.Next() {
		pp := e.Value.(*point)
		topLeft, bottomRight := pp.getBorders(len(board)-1, len(board[0])-1)
		for i := topLeft.x; i <= bottomRight.x; i++ {
			for j := topLeft.y; j <= bottomRight.y; j++ {
				if boardArr[i][j] == ' ' {
					boardArr[i][j] = '1'
				} else {
					if boardArr[i][j] == '*' {
						continue
					} else {
						v, _ := strconv.Atoi(string(boardArr[i][j]))
						s := strconv.Itoa(v + 1)
						boardArr[i][j] = rune(s[0])
					}
				}
			}
		}
	}
	resultBoard := arrayBackToBoard(boardArr)

	return resultBoard
}
