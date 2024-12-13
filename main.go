package main

import (
	"github.com/RafaelGervasio/chess-go/board"
	"github.com/RafaelGervasio/chess-go/piece"
	"github.com/RafaelGervasio/chess-go/movement"
	"github.com/RafaelGervasio/chess-go/square"
	"github.com/RafaelGervasio/chess-go/userinput"
	"os"
	"errors"
	"fmt"
	"log"
	"strings"
)


func main() {
	gameBoard := board.InitializeBoard()
	currentTurn := piece.Color
	
	isCheckmate, err := movement.Checkmate(gameBoard, currentTurn)
	if err != nil {
		log.Fatal("Checkmate function returned an error")
	}

	for !isCheckmate {
		gameBoard = playTurn(gameBoard, currentTurn)
		currentTurn = changeTurn(currentTurn)

		isCheckmate, err := movement.Checkmate(gameBoard, currentTurn)
		if err != nil {
			log.Fatal("Checkmate function returned an error")
		}
	}

}

func playTurn(board board.Board, turn piece.Color) board board.Board {
	board.DisplayBoard()
	startSquare, endSquare, piece, err := userinput.GetUserInput(board, turn)

	for err != nil {
		fmt.Println("Invalid input.")
		startSquare, endSquare, piece, err = userinput.GetUserInput(board, turn)
	}

	if movement.ValidMove(board, startSquare, endSquare, piece, turn) && !LeavesPlayerInCheck(board, startSquare, endSquare, piece, turn) {
	    boardToMutate := GetBoardCopy()
	    boardToMutate.DeleteFromBoard(endSquare)
	    boardToMutate.AddToBoard(endSquare, piece)
	    boardToMutate.DeleteFromBoard(startSquare)
	    piece.HasMoved = false
	    return boardToMutate
	} else {
		fmt.Println("Invalid input.")
		playTurn(board, turn)
	}
}


func changeTurn(turn piece.Color) turn piece.Color {
	if turn == piece.White {
		return piece.Black
	}
	return piece.White
}
	


