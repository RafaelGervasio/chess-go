package main

import (
	"fmt"
	"github.com/RafaelGervasio/chess-go/board"
	"github.com/RafaelGervasio/chess-go/piece"
)


type Game struct {
	Board Board
	Turn  Turn
}


func main() {
	var game Game
	var board Board

	game{Board: board, Turn: White}

	board.initializeBoard()
	board.displayBoard()
}
