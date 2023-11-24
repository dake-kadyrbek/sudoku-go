package main

import (
	"os"
	"piscine"
)

const GridSize = 9

func main() {
	args := os.Args[1:]
	if len(args) != GridSize {
		piscine.PrintError()
		return
	}

	board, err := piscine.ParseInput(args)
	if err {
		piscine.PrintError()
		return
	}

	if piscine.SolveSudoku(board) {
		piscine.PrintBoard(board)
	} else {
		piscine.PrintError()
	}
}
