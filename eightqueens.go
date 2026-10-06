package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var positions [8]int
	solveEightQueens(0, &positions)
}

func solveEightQueens(column int, positions *[8]int) {
	if column == len(*positions) {
		for i := 0; i < len(*positions); i++ {
			z01.PrintRune(rune(positions[i] + 49))
		}
		z01.PrintRune(10)
		return
	}

	for row := 0; row < len(*positions); row++ {
		valid := 1
		for previous := 0; previous < column; previous++ {
			if positions[previous] == row || positions[previous]-previous == row-column || positions[previous]+previous == row+column {
				valid = 0
				break
			}
		}
		if valid == 1 {
			positions[column] = row
			solveEightQueens(column+1, positions)
		}
	}
}
