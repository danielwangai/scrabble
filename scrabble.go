package main

type Cell struct {
	Row        int
	Column     int
	Letter     string
	Multiplier int
}

var Board [][]Cell

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

// func main() {
// 	fmt.Println(len(Board[0]))
// }
