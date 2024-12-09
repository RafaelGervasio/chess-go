package main

import (
	"github.com/RafaelGervasio/chess-go/board"
)


// type Game struct {
// 	Board Board
// 	Turn  Turn
// }


func main() {
	var myBoard board.Board

	myBoard.InitializeBoard()
	myBoard.DisplayBoard()
}
