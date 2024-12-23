package movement

import (
    "errors"
    "fmt"
    "github.com/RafaelGervasio/chess-go/board"
    "github.com/RafaelGervasio/chess-go/piece"
    "github.com/RafaelGervasio/chess-go/square"
)

func ValidMove(board board.Board, start, end square.Square, piece *piece.Piece, color piece.Color) bool {
    return !landedOnFriendlyPiece(board, color, end) && validMovePattern(board, start, end, piece, color) && !jumpingOverPiece(board, start, end)
}

// landedOnFriendlyPiece checks if the piece lands on a square occupied by a piece of the same color.
func landedOnFriendlyPiece(board board.Board, colorOfPiece piece.Color, end square.Square) bool {
    pieceInEndSquare := board.GetPieceFromSquare(end)
    return pieceInEndSquare != nil && colorOfPiece == pieceInEndSquare.Color
}

// validMovePattern checks if the move pattern is valid based on the piece's movement rules.
func validMovePattern(board board.Board, start, end square.Square, piece *piece.Piece, color piece.Color) bool {
    switch piece.Name {
    case "rook":
        return start.Row == end.Row || start.Col == end.Col
    case "bishop":
        return abs(end.Row-start.Row) == abs(end.Col-start.Col)
    case "queen":
        return (start.Row == end.Row || start.Col == end.Col) || abs(end.Row-start.Row) == abs(end.Col-start.Col)
    case "knight":
        return (abs(end.Row-start.Row) == 1 && abs(end.Col-start.Col) == 2) || (abs(end.Row-start.Row) == 2 && abs(end.Col-start.Col) == 1)
    case "king":
        return abs(end.Row-start.Row) <= 1 && abs(end.Col-start.Col) <= 1
    case "pawn":
        return validPawnMovePattern(board, start, end, piece, color)
    default:
        return false
    }
}

// validPawnMovePattern checks if a pawn's move follows the allowed movement pattern.
func validPawnMovePattern(board board.Board, start, end square.Square, chessPiece *piece.Piece, color piece.Color) bool {
    targetPiece := board.GetPieceFromSquare(end)

    // Handle capturing moves (diagonal attack).
    if targetPiece != nil {
        if color == piece.White {
            return end.Row-start.Row == 1 && abs(end.Col-start.Col) == 1
        }
        return end.Row-start.Row == -1 && abs(end.Col-start.Col) == 1
    }

    // Handle non-capturing (forward) moves.
    if color == piece.White {
        if !chessPiece.HasMoved {
            return start.Col == end.Col && (end.Row-start.Row == 1 || end.Row-start.Row == 2)
        }
        return start.Col == end.Col && end.Row-start.Row == 1
    }

    if color == piece.Black {
        if !chessPiece.HasMoved {
            return start.Col == end.Col && (end.Row-start.Row == -1 || end.Row-start.Row == -2)
        }
        return start.Col == end.Col && end.Row-start.Row == -1
    }

    return false
}



// jumpingOverPiece checks if a piece jumps over other pieces when moving.
func jumpingOverPiece(board board.Board, start, end square.Square) bool {
    pathCrossed := getPathCrossed(start, end)

    for _, square := range pathCrossed {
        if board.Positions[square] != nil {
            return true
        }
    }

    return false
}

// getPathCrossed returns the path crossed by a piece from start to end.
func getPathCrossed(start, end square.Square) []square.Square {
    var path []square.Square

    if start.Row == end.Row { // Horizontal move
        step := 1
        if start.Col > end.Col {
            step = -1
        }
        for col := start.Col + step; col != end.Col; col += step {
            path = append(path, square.Square{Row: start.Row, Col: col})
        }
    } else if start.Col == end.Col { // Vertical move
        step := 1
        if start.Row > end.Row {
            step = -1
        }
        for row := start.Row + step; row != end.Row; row += step {
            path = append(path, square.Square{Row: row, Col: start.Col})
        }
    } else if abs(end.Row-start.Row) == abs(end.Col-start.Col) { // Diagonal move
        rowStep := 1
        colStep := 1
        if start.Row > end.Row {
            rowStep = -1
        }
        if start.Col > end.Col {
            colStep = -1
        }
        row, col := start.Row+rowStep, start.Col+colStep
        for row != end.Row && col != end.Col {
            path = append(path, square.Square{Row: row, Col: col})
            row += rowStep
            col += colStep
        }
    }

    return path
}

// Check determines if the player of the given color is in check.
func Check(board board.Board, color piece.Color) (bool, error) {
    copyOfBoard := board.GetBoardCopy()

    kingSquare, err := copyOfBoard.GetKingSquare(color)
    if err != nil {
        return false, errors.New("check: king not found")
    }

    squaresAndPiecesOfOppositeColor := copyOfBoard.GetSquaresAndPiecesOfColor(oppositeColor(color))
    for square, piece := range squaresAndPiecesOfOppositeColor {
        if ValidMove(copyOfBoard, square, kingSquare, piece, oppositeColor(color)) {
            return true, nil
        }
    }

    return false, nil
}

// Checkmate determines if the player of the given color is in checkmate.
func Checkmate(board board.Board, color piece.Color) (bool, error) {
    inCheck, err := Check(board, color)
    if err != nil {
        return false, err
    }
    if !inCheck {
        return false, nil // Not in check, so not checkmate
    }

    squaresAndPiecesOfColor := board.GetSquaresAndPiecesOfColor(color)
    for sqr, piece := range squaresAndPiecesOfColor {
        for row := 1; row <= 8; row++ {
            for col := 1; col <= 8; col++ {
                targetSquare := square.Square{Row: row, Col: col}
                if ValidMove(board, sqr, targetSquare, piece, color) {
                    if !LeavesPlayerInCheck(board, sqr, targetSquare, piece, color) {
                        return false, nil
                    }
                }
            }
        }
    }

    return true, nil
}

// LeavesPlayerInCheck simulates a move and checks if it leaves the player in check.
func LeavesPlayerInCheck(board board.Board, start, end square.Square, piece *piece.Piece, color piece.Color) bool {
    hadMoved := piece.HasMoved

    boardToMutate := board.GetBoardCopy()
    boardToMutate.DeleteFromBoard(end)
    boardToMutate.AddToBoard(end, piece)
    boardToMutate.DeleteFromBoard(start)

    piece.HasMoved = hadMoved

    inCheck, err := Check(boardToMutate, color)
    if err != nil {
        fmt.Errorf("Leaves player in check: %w", err)
    }


    return inCheck
}

// abs returns the absolute value of an integer.
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
    
// oppositeColor returns the opposite color of the given color.
func oppositeColor(color piece.Color) piece.Color {
    if color == piece.White {
        return piece.Black
    }
    return piece.White
}