package main

import (
	"github.com/RafaelGervasio/chess-go/board"
	"github.com/RafaelGervasio/chess-go/piece"
	"github.com/RafaelGervasio/chess-go/movement"
	"github.com/RafaelGervasio/chess-go/square"
	"os"
	"errors"
	"fmt"
	"strings"
)


// type Game struct {
// 	Board board.Board
// 	Turn  piece.Color
// }




func main() {
	var myBoard board.Board

	myBoard.InitializeBoard()
	myBoard.DisplayBoard()

	// get user input
	// transalte that into start end piece color 
		// 
}



func getUserInput(board board.Board, turn piece.Color) (startSquare, endSquare square.Square, piece piece.Piece, err error) {
	inputString := getUserInputString()
	startRow, startCol, endRow, endCol, err := transformInputIntoInts(inputString)
	if err != nil {
		return nil, nil, nil, fmt.Errof("getUserInput: %w", err)
	}
	startSquare, endSquare, piece, err := getSquarePieceFromInts(board, startRow, startCol, endRow, endCol)	
	if err != nil {
		return nil, nil, nil, fmt.Errof("getUserInput: %w", err)
	}

	if !validInput(startSquare, endSquare, piece) {
		return nil, nil, nil, fmt.Errof("getUserInput: invalid input!")
	}

}

func getUserInputString() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a start-end pair: (e2-e4 eg): ")
	input, _ := reader.ReadString('\n')
	// Trim any trailing newline or whitespace
	return strings.TrimSpace(input)
}

func parseSquare(square string) (col, row int, err error) {
	if len(square) != 2 {
		return 0, 0, errors.New("square must be 2 characters (e.g., e2)")
	}

	col = int(square[0] - 'a' + 1) // Convert 'a'-'h' to 1-8
	row = int(square[1] - '0')     // Convert '1'-'8' to integers

	if col < 1 || col > 8 || row < 1 || row > 8 {
		return 0, 0, errors.New("row and column must be between 1 and 8")
	}

	return col, row, nil
}


func transformInputIntoInts(inputString string) (startRow, startCol, endRow, endCol int, err error) {
	if len(inputString) != 5 || inputString[2] != '-' {
		return 0, 0, 0, 0, fmt.Errorf("TransformInputIntoInts: invalid format, expected start-end pair (e.g., e2-e4)")
	}

	start := inputString[:2]
	end := inputString[3:]

	startCol, startRow, err = parseSquare(start)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("invalid start square: %v", err)
	}

	endCol, endRow, err = parseSquare(end)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("invalid end square: %v", err)
	}

	return startRow, startCol, endRow, endCol, nil
}


func getSquarePieceFromInts(board board.Board, startRow, startCol, endRow, endCol int) (startSquare, endSquare square.Square, piece piece.Piece, err error) {
	startSquare, piece, err := board.GetSquareAndPiece(startRow, startCol)
	if err != nil {
		return false, fmt.Errof("getSquarePieceColorFromInput: %w", err)
	}

	endSquare, _, err := board.GetSquareAndPiece(endRow, endCol)
	if err != nil {
		return false, fmt.Errof("getSquarePieceColorFromInput: %w", err)
	}

	return startSquare, endSquare, piece, nil

}

func validInput(startSquare, endSquare square.Square, piece piece.Piece, turn piece.Color) bool {
	if piece == nil || piece.Color != turn || (startSquare.Row == endSquare.Row && startSquare.Col == endSquare.Col) {
		return false
	}
	return true
}

	



