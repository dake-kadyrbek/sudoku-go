package piscine

import "github.com/01-edu/z01"

func isValid(board [][]int, row, col, num int) bool { // Check if the number is already in the row or column
	for i := 0; i < GridSize; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// Check if the number is in the 3x3 grid
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

func SolveSudoku(board [][]int) bool {
	empty := findEmptyLocation(board)
	if empty[0] == -1 && empty[1] == -1 {
		// No empty location, puzzle solved
		return true
	}

	row, col := empty[0], empty[1]

	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num

			if SolveSudoku(board) {
				return true
			}

			// If placing the current number doesn't lead to a solution, backtrack
			board[row][col] = 0
		}
	}

	// No valid number found, need to backtrack
	return false
}

func findEmptyLocation(board [][]int) [2]int {
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if board[i][j] == 0 {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1} // No empty location
}

func PrintBoard(board [][]int) {
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			z01.PrintRune(rune(board[i][j] + 48))
			if j == GridSize-1 {
				z01.PrintRune('\n')
				continue
			}
			z01.PrintRune(' ')
		}
	}
}
