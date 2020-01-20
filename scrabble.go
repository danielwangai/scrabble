package main

import "fmt"

type Cell struct {
	Row        int
	Column     int
	Letter     map[string]*LetterDetail
	Multiplier int
	Type       string // word or letter
}

var Board [][]Cell

type LetterDetail struct {
	Points int
	Count  int
}

// Letters and corresponding count and points
var Letters = map[string]LetterDetail{
	"A":     LetterDetail{Points: 1, Count: 9},
	"B":     LetterDetail{Points: 3, Count: 2},
	"C":     LetterDetail{Points: 3, Count: 2},
	"D":     LetterDetail{Points: 2, Count: 4},
	"E":     LetterDetail{Points: 1, Count: 12},
	"F":     LetterDetail{Points: 4, Count: 2},
	"G":     LetterDetail{Points: 2, Count: 3},
	"H":     LetterDetail{Points: 4, Count: 2},
	"I":     LetterDetail{Points: 1, Count: 9},
	"J":     LetterDetail{Points: 8, Count: 1},
	"K":     LetterDetail{Points: 5, Count: 1},
	"L":     LetterDetail{Points: 1, Count: 4},
	"M":     LetterDetail{Points: 3, Count: 2},
	"N":     LetterDetail{Points: 1, Count: 6},
	"O":     LetterDetail{Points: 1, Count: 8},
	"P":     LetterDetail{Points: 3, Count: 2},
	"Q":     LetterDetail{Points: 10, Count: 1},
	"R":     LetterDetail{Points: 1, Count: 6},
	"S":     LetterDetail{Points: 1, Count: 4},
	"T":     LetterDetail{Points: 1, Count: 6},
	"U":     LetterDetail{Points: 1, Count: 4},
	"V":     LetterDetail{Points: 4, Count: 2},
	"W":     LetterDetail{Points: 4, Count: 2},
	"X":     LetterDetail{Points: 8, Count: 1},
	"Y":     LetterDetail{Points: 4, Count: 2},
	"Z":     LetterDetail{Points: 10, Count: 1},
	"Blank": LetterDetail{Points: 0, Count: 2},
}

// assign point to cells
func init() {
	makeBoard()
	trippleWordCells()
	tripleLetterCells()
	doubleWordCells()
	doubleLetterCells()
}

func trippleWordCells() {
	// cell (0, 0)
	Board[0][0].Multiplier = 3
	Board[0][0].Type = "Word"

	// cell (0, 7)
	Board[0][7].Multiplier = 3
	Board[0][7].Type = "Word"

	// cell (0, 14)
	Board[0][14].Multiplier = 3
	Board[0][14].Type = "Word"

	// cell (7, 0)
	Board[7][0].Multiplier = 3
	Board[7][0].Type = "Word"

	// cell (14, 0)
	Board[14][0].Multiplier = 3
	Board[14][0].Type = "Word"

	// cell (14, 7)
	Board[14][7].Multiplier = 3
	Board[14][7].Type = "Word"

	// cell (14, 14)
	Board[14][14].Multiplier = 3
	Board[14][14].Type = "Word"

	// cell (7, 14)
	Board[7][14].Multiplier = 3
	Board[7][14].Type = "Word"
}

func tripleLetterCells() {
	// cell (1, 5)
	Board[1][5].Multiplier = 3
	Board[1][5].Type = "Letter"

	// cell (1, 9)
	Board[1][9].Multiplier = 3
	Board[1][9].Type = "Letter"

	// cell (5, 1)
	Board[5][13].Multiplier = 3
	Board[5][1].Type = "Letter"

	// cell (5, 5)
	Board[5][5].Multiplier = 3
	Board[5][5].Type = "Letter"

	// cell (5, 9)
	Board[5][9].Multiplier = 3
	Board[5][9].Type = "Letter"

	// cell (5, 13)
	Board[5][13].Multiplier = 3
	Board[5][13].Type = "Letter"

	// cell (9, 1)
	Board[9][1].Multiplier = 3
	Board[9][1].Type = "Letter"

	// cell (9, 5)
	Board[9][5].Multiplier = 3
	Board[9][5].Type = "Letter"

	// cell (9, 9)
	Board[9][9].Multiplier = 3
	Board[9][9].Type = "Letter"

	// cell (9, 13)
	Board[9][13].Multiplier = 3
	Board[9][13].Type = "Letter"

	// cell (1, 9)
	Board[1][9].Multiplier = 3
	Board[1][9].Type = "Letter"

	// cell (13, 5)
	Board[13][5].Multiplier = 3
	Board[13][5].Type = "Letter"

	// cell (13, 9)
	Board[13][9].Multiplier = 3
	Board[13][9].Type = "Letter"
}

func doubleWordCells() {
	// cell (1, 1)
	Board[1][1].Multiplier = 2
	Board[1][1].Type = "Word"

	// cell (1, 13)
	Board[1][13].Multiplier = 2
	Board[1][13].Type = "Word"

	// cell (2, 2)
	Board[2][2].Multiplier = 2
	Board[2][2].Type = "Word"

	// cell (2, 12)
	Board[2][12].Multiplier = 2
	Board[2][12].Type = "Word"

	// cell (3, 3)
	Board[3][3].Multiplier = 2
	Board[3][3].Type = "Word"

	// cell (3, 11)
	Board[3][11].Multiplier = 2
	Board[3][11].Type = "Word"

	// cell (4, 4)
	Board[4][4].Multiplier = 2
	Board[4][4].Type = "Word"

	// cell (4, 10)
	Board[4][10].Multiplier = 2
	Board[4][10].Type = "Word"

	// cell (10, 4)
	Board[10][4].Multiplier = 2
	Board[10][4].Type = "Word"

	// cell (10, 10)
	Board[10][10].Multiplier = 2
	Board[10][10].Type = "Word"

	// cell (11, 3)
	Board[11][3].Multiplier = 2
	Board[11][3].Type = "Word"

	// cell (11, 11)
	Board[11][11].Multiplier = 2
	Board[11][11].Type = "Word"

	// cell (12, 2)
	Board[12][2].Multiplier = 2
	Board[12][2].Type = "Word"

	// cell (12, 12)
	Board[12][12].Multiplier = 2
	Board[12][12].Type = "Word"

	// cell (13, 1)
	Board[13][1].Multiplier = 2
	Board[13][1].Type = "Word"

	// cell (13, 13)
	Board[13][13].Multiplier = 2
	Board[13][13].Type = "Word"
}

func doubleLetterCells() {
	// cell (0, 3)
	Board[0][3].Multiplier = 2
	Board[0][3].Type = "Letter"

	// cell (0, 11)
	Board[0][11].Multiplier = 2
	Board[0][11].Type = "Letter"

	// cell (2, 6)
	Board[2][6].Multiplier = 2
	Board[2][6].Type = "Letter"

	// cell (2, 8)
	Board[2][8].Multiplier = 2
	Board[2][8].Type = "Letter"

	// cell (3, 0)
	Board[3][0].Multiplier = 2
	Board[3][0].Type = "Letter"

	// cell (3, 7)
	Board[3][7].Multiplier = 2
	Board[3][7].Type = "Letter"

	// cell (3, 14)
	Board[3][14].Multiplier = 2
	Board[3][14].Type = "Letter"

	// cell (6, 2)
	Board[6][2].Multiplier = 2
	Board[6][2].Type = "Letter"

	// cell (6, 6)
	Board[6][6].Multiplier = 2
	Board[6][6].Type = "Letter"

	// cell (6, 8)
	Board[6][8].Multiplier = 2
	Board[6][8].Type = "Letter"

	// cell (6, 12)
	Board[6][12].Multiplier = 2
	Board[6][12].Type = "Letter"

	// cell (7, 3)
	Board[7][3].Multiplier = 2
	Board[7][3].Type = "Letter"

	// cell (7, 11)
	Board[7][11].Multiplier = 2
	Board[7][11].Type = "Letter"

	// cell (8, 2)
	Board[8][2].Multiplier = 2
	Board[8][2].Type = "Letter"

	// cell (8, 6)
	Board[8][6].Multiplier = 2
	Board[8][6].Type = "Letter"

	// cell (8, 8)
	Board[8][8].Multiplier = 2
	Board[8][8].Type = "Letter"

	// cell (8, 12)
	Board[8][12].Multiplier = 2
	Board[8][12].Type = "Letter"

	// cell (11, 0)
	Board[11][0].Multiplier = 2
	Board[11][0].Type = "Letter"

	// cell (11, 7)
	Board[11][7].Multiplier = 2
	Board[11][7].Type = "Letter"

	// cell (11, 14)
	Board[11][14].Multiplier = 2
	Board[11][14].Type = "Letter"

	// cell (12, 6)
	Board[12][6].Multiplier = 2
	Board[12][6].Type = "Letter"

	// cell (12, 8)
	Board[12][8].Multiplier = 2
	Board[12][8].Type = "Letter"

	// cell (14, 3)
	Board[14][3].Multiplier = 2
	Board[14][3].Type = "Letter"

	// cell (14, 11)
	Board[14][11].Multiplier = 2
	Board[14][11].Type = "Letter"
}

// generate the 2D(1%X15) grid/board
func makeBoard() {
	count := 0
	for count < 15 {
		row := make([]Cell, 15)
		for i := 0; i < 15; i++ {
			cell := &Cell{}
			// default multiplier for each cell is 1
			cell.Multiplier = 1
			cell.Row = count
			cell.Column = i
			row[i] = *cell
		}
		Board = append(Board, row)
		count++
	}
}

func main() {
	// fmt.Println(len(Board[0]))
	fmt.Println(Letters)
}
