package main

import (
	"fmt"
	"log"

	"github.com/RafaelGervasio/chess-go/board"
	"github.com/RafaelGervasio/chess-go/movement"
	"github.com/RafaelGervasio/chess-go/piece"
	"github.com/RafaelGervasio/chess-go/userinput"
)

func main() {
	var gameBoard board.Board
	gameBoard.InitializeBoard()
	currentTurn := piece.White // Set the initial turn to White

	isCheckmate, err := movement.Checkmate(gameBoard, currentTurn)
	if err != nil {
		log.Fatalf("Error in Checkmate function: %v", err)
	}

	for !isCheckmate {
		gameBoard = playTurn(gameBoard, currentTurn)
		currentTurn = changeTurn(currentTurn)

		isCheckmate, err = movement.Checkmate(gameBoard, currentTurn)
		if err != nil {
			log.Fatalf("Error in Checkmate function: %v", err)
		}
	}
	fmt.Println("Checkmate! Game over.")
}

// playTurn handles a player's turn.
func playTurn(gameBoard board.Board, turn piece.Color) board.Board {
	gameBoard.DisplayBoard()
	startSquare, endSquare, selectedPiece, err := userinput.GetUserInput(gameBoard, turn)

	for err != nil {
		fmt.Println("Invalid input.")
		startSquare, endSquare, selectedPiece, err = userinput.GetUserInput(gameBoard, turn)
	}

	if movement.ValidMove(gameBoard, startSquare, endSquare, selectedPiece, turn) &&
		!movement.LeavesPlayerInCheck(gameBoard, startSquare, endSquare, selectedPiece, turn) {

		gameBoard.DeleteFromBoard(endSquare)
		gameBoard.AddToBoard(endSquare, selectedPiece)
		gameBoard.DeleteFromBoard(startSquare)
		selectedPiece.HasMoved = true
		return gameBoard
	}

	fmt.Println("Invalid move. Try again.")
	return playTurn(gameBoard, turn)
}

// changeTurn switches the current turn to the opposite color.
func changeTurn(turn piece.Color) piece.Color {
	if turn == piece.White {
		return piece.Black
	}
	return piece.White
}
