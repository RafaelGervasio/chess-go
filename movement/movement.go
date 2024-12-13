
// The color and piece will be deduced in the driver / input handling. teh funcs here that don't need them at all can not have them, btu to thee ones that do i'll just pass it in. that way i'm accessing a piece and a turn from an squser once rather than in every func
// This package should hanlde ALL needed cecks for a vlaid move, and EXPORT ONE function called ValidMove (the rest should be internal to the driver package.)
    // The driver package should get user input and determine the square piece etc... and then amke ONE func call that hides all the stuff away.





package movement

import (
    "github.com/RafaelGervasio/chess-go/board"
    "github.com/RafaelGervasio/chess-go/piece"
    "github.com/RafaelGervasio/chess-go/square"
    "errors"
)


func ValidMove(board board.Board, start, end square.Square, piece piece.Piece, color piece.Color) bool {
    return !landedOnFriendlyPiece(board, color, end) && validMovePattern(board, start, end, piece, color) &&  !jumpingOverPiece(board, start, end)
}



// landedOnFriendlyPiece checks if the piece lands on a square occupied by a piece of the same color.
func landedOnFriendlyPiece(board board.Board, colorOfPiece Color, end square.Square) bool {
    pieceInEndSquare = board.GetPieceFromSquare(end)
    return pieceInEndSquare != nil && color == pieceInEndSquare.Color
}

// validMovePattern checks if the move pattern is valid based on the piece's movement rules.
func validMovePattern(board board.Board, start, end square.Square, piece piece.Piece, color piece.Color) bool {
    
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
        return validPawnMovePattern(board, start, end, piece)
    default:
        return false
    }
}

// validPawnMovePattern checks if a pawn's move is valid.
func validPawnMovePattern(board board.Board, start, end square.Square, piece piece.Piece, color piece.Color) bool {
    // Assuming white is moving up the board (increasing row) and black is moving down
    
    if board.GetPieceFromSquare(end) != nil {
        if color == White {
            return end.Row-start.Row == 1 && abs(end.Col-start.Col) == 1
        } else {
            return end.Row-start.Row == -1 && abs(end.Col-start.Col) == 1
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

// jumpingOverPiece checks if a piece jumps over other pieces when moving.
func jumpingOverPiece(board board.Board, start, end square.Square) bool {
    pathCrossed := getPathCrossed(piece.Name, start, end)

    for _, square := range pathCrossed {
        if board.Positions[square] != nil {
            return true
        }
    }

    return false
}


// getPathCrossed returns the path crossed by a piece from start to end.
func getPathCrossed(pieceName string, start, end square.Square) []square.Square {
    var path []square.Square

    switch pieceName {
    case "Rook":
        if start.Row == end.Row { // Moving horizontally
            step := 1
            if start.Col > end.Col { // Moving left
                step = -1
            }
            for col := start.Col + step; col != end.Col; col += step {
                path = append(path, square.Square{Row: start.Row, Col: col})
            }
        } else if start.Col == end.Col { // Moving vertically
            step := 1
            if start.Row > end.Row { // Moving up
                step = -1
            }
            for row := start.Row + step; row != end.Row; row += step {
                path = append(path, square.Square{Row: row, Col: start.Col})
            }
        }

    case "Bishop":
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

    case "Queen":
        if start.Row == end.Row || start.Col == end.Col {
            // Use Rook's logic
            path = getPathCrossed("Rook", start, end)
        } else {
            // Use Bishop's logic
            path = getPathCrossed("Bishop", start, end)
        }

    case "King", "Knight":
        // No path crossed for King or Knight
        // They jump directly to the destination
        return []square.Square{}

    case "Pawn":
        // Ignoring for now as per your request
        return []square.Square{}
    }

    return path
}




// Check determines if the player of the given color is in check.
func Check(board board.Board, color piece.Color) bool, err {
    copyOfBoard := b.GetBoardCopy()

    kingSquare, err := copyOfBoard.GetKingSquare(color)
    
    if err != nil {
        return false, fmt.Errof("Check: %w", err)
    }

    squaresAndPiecesOfOppositeColor := copyOfBoard.GetSquaresAndPiecesOfColor(oppositeColor(color))
    
    for square, piece := range squaresAndPiecesOfOppositeColor {
        if movement.ValidMove(copyOfBoard, square, kingSquare, piece, oppositeColor(color)) {
            return true
        }
    }

    return false
}



func Checkmate(board board.Board, color Color) (bool, error) {
    if inCheck, err := Check(board, color); err != nil {
        return false, fmt.Errorf("Checkmate: %w", err)
    } else if !inCheck {
        return false, nil // Not in check, so not Checkmate
    }

    copyOfBoard := board.GetBoardCopy()
    squaresAndPiecesOfColor := copyOfBoard.GetSquaresAndPiecesOfColor(color)

    for square, piece := range squaresAndPiecesOfColor {
        for row := 1; row <= 8; row++ {
            for col := 1; col <= 8; col++ {
                targetSquare := Square{Row: row, Col: col}

                if ValidMove(copyOfBoard, square, targetSquare, piece, color) {
                    if !LeavesPlayerInCheck(copyOfBoard, square, targetSquare, piece, color) {
                        return false
                    }
                }
            }
        }
    }

    return true, nil 
}



func LeavesPlayerInCheck(board board.Board, start, end square.Square, piece piece.Piece, color piece.Color) bool {
    hadMoved := piece.HasMoved
    
    boardToMutate := GetBoardCopy()
    boardToMutate.DeleteFromBoard(end)
    boardToMutate.AddToBoard(end, piece)
    boardToMutate.DeleteFromBoard(start)

    piece.HasMoved = hadMoved

    return Check(boardToMutate, color)
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

