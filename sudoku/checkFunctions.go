package piscine

import (
	"github.com/01-edu/z01"
)

const (
	GridSize       = 9
	MinFilledCells = 17
)

func PrintError() {
	err := "Error"
	for _, c := range err {
		z01.PrintRune(rune(c))
	}
	z01.PrintRune('\n')
}

func ParseInput(args []string) ([][]int, bool) {
	board := make([][]int, GridSize)
	filledCells := 0

	for i, arg := range args {
		if len(arg) != GridSize {
			return nil, true
		}

		board[i] = make([]int, GridSize)
		for j, char := range arg {
			if char == '.' {
				board[i][j] = 0
			} else {
				num := int(char - '0')
				if num < 1 || num > 9 {
					return nil, true
				}
				board[i][j] = num
				filledCells++
			}
		}
	}

	if filledCells < MinFilledCells {
		return nil, true
	}

	if checkHorizontally(board) || checkVertically(board) {
		return nil, true
	}

	return board, false
}

func checkHorizontally(board [][]int) bool {
	var firstCandidate int
	var secondCandidate int
	var row int

	for i := 0; i < GridSize; i++ {
		for j := i + 1; j < GridSize; j++ {
			firstCandidate = board[row][i]
			secondCandidate = board[row][j]

			if firstCandidate == secondCandidate && (firstCandidate != 0 && secondCandidate != 0) {
				return true
			}
		}
		row++
	}

	return false
}

func checkVertically(board [][]int) bool {
	var firstCandidate int
	var secondCandidate int

	for row := 0; row < GridSize; row++ {
		for column := 0; column < GridSize; column++ {
			for column2 := column + 1; column2 < GridSize; column2++ {
				firstCandidate = board[column][row]
				secondCandidate = board[column2][row]

				if firstCandidate == secondCandidate && (firstCandidate != 0 && secondCandidate != 0) {
					return true
				}
			}
		}
	}
	return false
}
