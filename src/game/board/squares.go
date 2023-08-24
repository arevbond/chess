package board

import (
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
)

func (b *Board) IsSquareDark(coordinates coords.Coordinates) bool {
	return (int(coordinates.File)+int(coordinates.Rank))%2 == 0
}

func (b *Board) IsSquareEmpty(coordinates coords.Coordinates) bool {
	_, ok := b.GetPiece(coordinates)
	return !ok
}

func (b *Board) AttackedSquaresByPiece(figure piece.Piece) map[coords.Coordinates]bool {
	attackedSquares := make(map[coords.Coordinates]bool)
	switch figure.Name() {
	case "Pawn":
		attackedSquares = b.AttackedSquaresByPawn(figure)
	case "King":
		attackedSquares = b.AttackedSquaresByKing(figure)
	default:
		attackedSquares = b.AvailableMoves(figure)
	}
	return attackedSquares
}

func (b *Board) AttackedSquaresByKing(king piece.Piece) map[coords.Coordinates]bool {
	attackedSquares := make(map[coords.Coordinates]bool)
	shifts := king.Shifts()
	for shift, _ := range shifts {
		if king.Coordinates().CanShift(shift) {
			newCoordinates := king.Coordinates().Shift(shift)
			attackedSquares[newCoordinates] = true
		}
	}
	return attackedSquares
}

func (b *Board) AttackedSquaresByPawn(pawn piece.Piece) map[coords.Coordinates]bool {
	attackedSquares := make(map[coords.Coordinates]bool)
	var shifts [][]int
	if pawn.Color() == color.White {
		shifts = [][]int{
			{-1, 1}, {1, 1},
		}
	} else {
		shifts = [][]int{
			{-1, -1}, {1, -1},
		}
	}
	coordsShifts := coords.CalculateCoordinatesShift(shifts, pawn.Coordinates())
	for shift, _ := range coordsShifts {
		if pawn.Coordinates().CanShift(shift) {
			newCoordinates := pawn.Coordinates().Shift(shift)
			attackedSquares[newCoordinates] = true
		}
	}
	return attackedSquares
}

func (b *Board) IsSquareAttackedByColor(coordinates coords.Coordinates, enemyColor color.Color) bool {
	enemyPieces := b.PiecesByColor(enemyColor)
	for _, enemyPiece := range enemyPieces {
		attackedSquare := b.AttackedSquaresByPiece(enemyPiece)
		if _, ok := attackedSquare[coordinates]; ok {
			return true
		}
	}
	return false
}
