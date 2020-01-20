package main

import "fmt"

type Cell struct {
	Row        int
	Column     int
	Letter     map[string]*LetterDetail
	Multiplier int
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
}

// set default multiplier
// the minimum multiplier value for a cell is 1
// unless the case of a blank where we multiply by 0
func (c *Cell) init() {
	c.Multiplier = 1
}

// generate the 2D(1%X15) grid/board
func makeBoard() {
	count := 0
	for count < 15 {
		row := make([]Cell, 15)
		for i := 0; i < 15; i++ {
			cell := &Cell{}
			cell.init()
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
