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
		attackedSquares = b.AttackedSquaresByLongRangePiece(figure)
	}
	return attackedSquares
}

func (b *Board) AttackedSquaresByLongRangePiece(figure piece.Piece) map[coords.Coordinates]bool {
	attackedSquares := map[coords.Coordinates]bool{}
	shifts := figure.Shifts()
	for shift, _ := range shifts {
		if figure.Coordinates().CanShift(shift) {
			newCoordinates := figure.Coordinates().Shift(shift)
			if !b.IsSquareAvailableForAttackSimple(newCoordinates, figure) {
				break
			}
			if b.IsSquareAttackedByLongRangePiece(newCoordinates, figure) {
				attackedSquares[newCoordinates] = true
			}
		}
	}
	return attackedSquares
}

func (b *Board) IsSquareAttackedByLongRangePiece(coordinates coords.Coordinates, figure piece.Piece) bool {
	figureCoords := figure.Coordinates()
	return !b.HasPieceOnWayForLongRangePiece2(figureCoords, coordinates)
}

func (b *Board) IsSquareAvailableForAttackSimple(coordinates coords.Coordinates, figure piece.Piece) bool {
	// проверяет пустая ли клетка, если нет, стоит ли на ней вражеская фигура
	if b.IsSquareEmpty(coordinates) {
		return true
	}
	otherPiece, _ := b.GetPiece(coordinates)
	return figure.Color() != otherPiece.Color()
}

func (b *Board) HasPieceOnWayForLongRangePiece2(from, to coords.Coordinates) bool {
	var flag bool
	coordsBetween := SquaresBetween(from, to)
	ourPiece, _ := b.GetPiece(from)
	ourPieceColor := ourPiece.Color()
	for _, coordinates := range coordsBetween {
		if otherPiece, ok := b.GetPiece(coordinates); ok {
			otherPieceColor := otherPiece.Color()
			if flag || otherPieceColor == ourPieceColor {
				return true
			} else {
				flag = true
			}
		}
	}
	if flag == true {
		pieceInTo, ok := b.GetPiece(to)
		if ok {
			return ourPieceColor == pieceInTo.Color()
		}
		return true
	}
	return false
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
