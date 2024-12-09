package movement

import (
    "github.com/RafaelGervasio/chess-go/board"
    "github.com/RafaelGervasio/chess-go/piece"
    "github.com/RafaelGervasio/chess-go/square"
    "errors"
)

// ValidMove checks if the move from 'start' to 'end' is valid for the given piece.
func ValidMove(board board.Board, start, end square.Square) bool {
    piece := board.Positions[start]
    if piece == nil {
        return false
    }
    color := piece.Color

    return !(PieceStayedStill(start, end) ||
        !LandedOnFriendlyPiece(board, *piece, end) ||
        ValidMovePattern(board, start, end) ||
        !JumpingOverPiece(board, start, end) ||
        !LeavesPlayerInCheck(board, start, end, *piece, color))
}


// LeavesPlayerInCheck checks if the player will be in check after moving the piece.
func LeavesPlayerInCheck(board board.Board, start, end square.Square, piece piece.Piece, color piece.Turn) bool {
    // Backup piece's state
    hadMoved := piece.Name
    newBoard := GetUpdatedBoard(board, start, end, piece)
    piece.Name = hadMoved

    return Check(newBoard, color)
}

// Check determines if the player of the given color is in check.
func Check(board board.Board, color piece.Turn) bool {
    kingSquare := GetKingSquare(board, color)
    if kingSquare == nil {
        return false
    }

    piecesOfOppositeColor := GetAllPiecesOfColor(board, OppositeColor(color))
    for _, p := range piecesOfOppositeColor {
        if ValidMove(board, p.Position, kingSquare.Position) {
            return true
        }
    }
    return false
}

// OppositeColor returns the opposite color of the given color.
func OppositeColor(color piece.Turn) piece.Turn {
    if color == piece.White {
        return piece.Black
    }
    return piece.White
}

// PieceStayedStill checks if the piece stayed at the same position.
func PieceStayedStill(start, end square.Square) bool {
    return start.Row == end.Row && start.Col == end.Col
}

// LandedOnFriendlyPiece checks if the piece lands on a square occupied by a piece of the same color.
func LandedOnFriendlyPiece(board board.Board, piece piece.Piece, end square.Square) bool {
    square := board.Positions[end]
    return square != nil && square.Color == piece.Color
}

// ValidMovePattern checks if the move pattern is valid based on the piece's movement rules.
func ValidMovePattern(board board.Board, start, end square.Square) bool {
    piece := board.Positions[start]
    if piece == nil {
        return false
    }

    switch piece.Name {
    case "rook":
        return start.Row == end.Row || start.Col == end.Col
    case "bishop":
        return Abs(end.Row-start.Row) == Abs(end.Col-start.Col)
    case "queen":
        return (start.Row == end.Row || start.Col == end.Col) || Abs(end.Row-start.Row) == Abs(end.Col-start.Col)
    case "knight":
        return (Abs(end.Row-start.Row) == 1 && Abs(end.Col-start.Col) == 2) || (Abs(end.Row-start.Row) == 2 && Abs(end.Col-start.Col) == 1)
    case "king":
        return Abs(end.Row-start.Row) <= 1 && Abs(end.Col-start.Col) <= 1
    case "pawn":
        return ValidPawnMove(board, start, end, piece)
    default:
        return false
    }
}

// Abs returns the absolute value of an integer.
func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// ValidPawnMove checks if a pawn's move is valid.
func ValidPawnMove(board board.Board, start, end square.Square, p piece.Piece) bool {
    // Assuming white is moving up the board (increasing row) and black is moving down
    color := p.Color
    if board.Positions[end] != nil {
        if color == piece.White {
            return end.Row-start.Row == 1 && Abs(end.Col-start.Col) == 1
        } else {
            return end.Row-start.Row == -1 && Abs(end.Col-start.Col) == 1
        }
    }

    if color == piece.White {
        if !p.HasMoved {
            return start.Col == end.Col && (end.Row-start.Row == 1 || end.Row-start.Row == 2)
        }
        return start.Col == end.Col && end.Row-start.Row == 1
    }

    if color == piece.Black {
        if !p.HasMoved {
            return start.Col == end.Col && (end.Row-start.Row == -1 || end.Row-start.Row == -2)
        }
        return start.Col == end.Col && end.Row-start.Row == -1
    }

    return false
}

// JumpingOverPiece checks if a piece jumps over other pieces when moving.
func JumpingOverPiece(board board.Board, start, end square.Square) bool {
    piece := board.Positions[start]
    if piece == nil {
        return false
    }

    pathCrossed := GetPathCrossed(piece.Name, start, end)
    for _, square := range pathCrossed {
        if board.Positions[square] != nil {
            return true
        }
    }
    return false
}

// GetPathCrossed returns the path crossed by a piece from start to end.
func GetPathCrossed(pieceName string, start, end square.Square) []square.Square {
    // Implementation of path calculation based on piece name (rook, bishop, queen, etc.)
    return nil // Placeholder for actual path calculation logic
}

// GetUpdatedBoard returns a new board with the piece moved.
func GetUpdatedBoard(board board.Board, start, end square.Square, piece piece.Piece) board.Board {
    // Implement logic to return a new board with updated piece positions
    return board
}

// GetKingSquare returns the square of the king for the given color.
func GetKingSquare(board board.Board, color piece.Turn) *square.Square {
    for _, p := range board.Positions {
        if p != nil && p.Name == "king" && p.Color == color {
            return &square.Square{Row: p.Position.Row, Col: p.Position.Col}
        }
    }
    return nil
}

// GetAllPiecesOfColor returns all pieces of the opposite color.
func GetAllPiecesOfColor(board board.Board, color piece.Turn) []piece.Piece {
    pieces := []piece.Piece{}
    for _, p := range board.Positions {
        if p != nil && p.Color == color {
            pieces = append(pieces, *p)
        }
    }
    return pieces
}
