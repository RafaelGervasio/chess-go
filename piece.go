package piece

type Turn int

const (
	White Turn = iota
	Black
)

type Piece struct {
	Name    string
	Color   Turn
	Display string
}


// Piece interface defines valid movement and display for pieces
type PieceMovement interface {
	ValidMove(from, to Square, board Board) bool
}


func (p Piece) landedOnFriendlyPiece(from, to Square, board Board) bool {
	if board[to] == nil {
		return false
		// landing square is empty
	}
	return board[to].Color == p.Color
}

// func (p Piece) validMovementPattern(from, to Square, board Board) bool {
// 	switch p.Name {
// 		case "Pawn":
// 			/* code */
// 		case "Rook":
// 			// code
// 		default:
// 			/* code */
// 			return
// 		}
// }


// func (p Piece) jumpedOverPiece (from, to Square, board Board) bool {
// 		switch p.Name {
// 		case "Pawn":
// 			/* code */
// 		case "Rook":
// 			// code
// 		default:
// 			/* code */
// 			return
// 		}
// }




// input
// from ....
	// get square
	// board[square] 
		// got the piece
	// piece.validmove?
	// delete current piece from stating square
	// delete piece from ending square
	// add current piece from ending square

