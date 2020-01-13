package main

var Board [][]string

func init() {
	makeBoard()
}

func makeBoard() {
	count := 15
	for count > 0 {
		row := []string{}
		for i := 0; i < 15; i++ {
			row = append(row, "")
		}
		Board = append(Board, row)
		count--
	}
}

// func main() {
// 	fmt.Println(len(Board[0]))
// }
