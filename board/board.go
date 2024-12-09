package board

import (
	"fmt"
	"github.com/RafaelGervasio/chess-go/piece" // import the Piece package
	"github.com/RafaelGervasio/chess-go/square" // import the Square package
)

type Board struct {
	Positions map[square.Square]*piece.Piece
}


// Method to initialize the board
func (b *Board) InitializeBoard() {
	b.Positions = make(map[square.Square]*piece.Piece)

	for row := 1; row <= 8; row++ {
		for col := 1; col <= 8; col++ {
			square := square.Square{Row: row, Col: col}

			switch row {
			case 1:
				b.Positions[square] = createPiece(col, piece.White)
			case 2:
				b.Positions[square] = &piece.Piece{Name: "Pawn", Color: piece.White, Display: "♙", Moved: false}
			case 7:
				b.Positions[square] = &piece.Piece{Name: "Pawn", Color: piece.Black, Display: "♟", Moved: false}
			case 8:
				b.Positions[square] = createPiece(col, piece.Black)
			default:
				b.Positions[square] = nil
			}
		}
	}
}

// Helper function to create pieces for the back row based on column
func createPiece(col int, color piece.Turn) *piece.Piece {
	switch col {
	case 1, 8:
		return &piece.Piece{Name: "Rook", Color: color, Display: "♖", Moved: false}
	case 2, 7:
		return &piece.Piece{Name: "Knight", Color: color, Display: "♘", Moved: false}
	case 3, 6:
		return &piece.Piece{Name: "Bishop", Color: color, Display: "♗", Moved: false}
	case 4:
		return &piece.Piece{Name: "Queen", Color: color, Display: "♕", Moved: false}
	case 5:
		return &piece.Piece{Name: "King", Color: color, Display: "♔", Moved: false}
	default:
		return nil
	}
}

// Method to display the board
func (b *Board) DisplayBoard() {
	for row := 1; row <= 8; row++ {
		for col := 1; col <= 8; col++ {
			square := square.Square{Row: row, Col: col}
			piece := b.Positions[square]

			if piece != nil {
				fmt.Printf("[%d,%d]: %s  ", square.Row, square.Col, piece.Display)
			} else {
				fmt.Printf("[%d,%d]: Empty  ", square.Row, square.Col)
			}
		}
		fmt.Println() // Newline after each row
	}
}

// Add a piece to the board at a specific square
func (b *Board) AddToBoard(square square.Square, piece *piece.Piece) {
	b.Positions[square] = piece
}

// Remove a piece from the board at a specific square
func (b *Board) DeleteFromBoard(square square.Square) {
	b.Positions[square] = nil
}

func (b Board) GetSquareAndPiece(row, col int) (square.Square, piece.Piece, err) {
	for square, piece := range b.Positions {
		if square.Row == row && square.Col == col {
			return square, piece, nil
		}
	}
	return nil, nil, fmt.Errof("Square not found for row: %d and col: %d", row, col)
}


// GetBoardCopy returns a copy of the board with all its pieces.
func (b Board) GetBoardCopy() Board {
	// Create a new map to hold the copied positions
	copyPositions := make(map[square.Square]*piece.Piece)

	// Iterate over the original board's positions and copy the pieces
	for square, pieceInSquare := range b.Positions {
		if pieceInSquare != nil {
			// Create a new piece copy
			copyPositions[square] = &piece.Piece{
				Name:    pieceInSquare.Name,
				Color:   pieceInSquare.Color,
				Display: pieceInSquare.Display,
			}
		} else {
			// If the piece is nil, just set the value to nil in the copy
			copyPositions[square] = nil
		}
	}

	// Return a new Board with the copied positions
	return Board{Positions: copyPositions}
}



// I like this. Seems like a pretty clear and complete package.