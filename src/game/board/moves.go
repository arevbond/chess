package board

import (
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
	"chess/src/utils"
)

func (b *Board) AvailableMoves(figure piece.Piece) map[coords.Coordinates]bool {
	availableMoves := map[coords.Coordinates]bool{}
	isKingCheck := b.IsKingInCheck(figure.Color())
	shifts := figure.Shifts()
	for shift, _ := range shifts {
		if figure.Coordinates().CanShift(shift) {
			newCoordinates := figure.Coordinates().Shift(shift)
			if b.IsSquareAvailableForMove(newCoordinates, figure, isKingCheck) {
				availableMoves[newCoordinates] = true
			}
		}
	}
	return availableMoves
}

func (b *Board) IsSquareAvailableForMoveSimple(coordinates coords.Coordinates, figure piece.Piece) bool {
	// проверяет пустая ли клетка, если нет, стоит ли на ней вражеская фигура, кроме короля
	if b.IsSquareEmpty(coordinates) {
		return true
	}
	otherPiece, _ := b.GetPiece(coordinates)
	return figure.Color() != otherPiece.Color() && otherPiece.Name() != "King"
}

func (b *Board) IsSquareAvailableForMove(coordinates coords.Coordinates, figure piece.Piece, isKingCheck bool) bool {
	if !b.IsSquareAvailableForMoveSimple(coordinates, figure) {
		return false
	}
	var ans bool
	if isKingCheck {
		ans = b.IsSquareAvailableForMoveWithCheck(coordinates, figure)
	} else {
		ans = b.IsSquareAvailableForMoveByPiece(coordinates, figure)
	}
	if ans {
		ans = b.IsSquareAvailableForMoveWithCheck(coordinates, figure)
	}
	return ans
}

func (b *Board) IsSquareAvailableForMoveByPiece(coordinates coords.Coordinates, figure piece.Piece) bool {
	var ans bool
	switch figure.Name() {
	case "Knight":
		ans = b.IsSquareAvailableForMoveKnight(coordinates, figure)
	case "Bishop", "Rock", "Queen":
		ans = b.IsSquareAvailableForMoveLongRangePiece(coordinates, figure)
	case "Pawn":
		ans = b.IsSquareAvailableForMovePawn(coordinates, figure)
	case "King":
		ans = b.IsSquareAvailableForMoveKing(coordinates, figure)
	}
	return ans
}

func (b *Board) IsSquareAvailableForMoveWithCheck(coordinates coords.Coordinates, figure piece.Piece) bool {
	copyBoard := Copy(b)
	copyFigure, _ := copyBoard.GetPiece(figure.Coordinates())
	if copyBoard.IsSquareAvailableForMoveByPiece(coordinates, copyFigure) {
		copyBoard.MovePiece(copyFigure.Coordinates(), coordinates)
		if !copyBoard.IsKingInCheck(copyFigure.Color()) {
			return true
		}
	}
	return false
}

func (b *Board) IsSquareAvailableForMoveKnight(coordinates coords.Coordinates, figure piece.Piece) bool {
	if b.IsSquareEmpty(coordinates) {
		return true
	}
	otherPiece, _ := b.GetPiece(coordinates)
	return figure.Color() != otherPiece.Color() && otherPiece.Name() != "King"
}

func (b *Board) IsSquareAvailableForMovePawn(coordinates coords.Coordinates, figure piece.Piece) bool {
	// TODO: Добавить взятие на проходе
	var result bool
	figureCoords := figure.Coordinates()
	if figureCoords.Rank == coordinates.Rank {
		if utils.Abs(int(figureCoords.File-coordinates.File)) == 2 {
			if figure.Color() == color.White && figureCoords.File == coords.File(2) {
				result = !b.HasPieceOnWay(figureCoords, coordinates)
			} else if figure.Color() == color.Black && figureCoords.File == coords.File(7) {
				result = !b.HasPieceOnWay(figureCoords, coordinates)
			} else {
				result = false
			}
		} else if utils.Abs(int(figureCoords.File-coordinates.File)) == 1 {
			result = !b.HasPieceOnWay(figureCoords, coordinates)
		}
	} else {
		otherPiece, ok := b.GetPiece(coordinates)
		if ok && otherPiece.Color() != figure.Color() {
			result = true
		} else {
			result = false
		}
	}
	return result
}

func (b *Board) IsSquareAvailableForMoveKing(coordinates coords.Coordinates, figure piece.Piece) bool {
	return !b.IsSquareAttackedByColor(coordinates, color.Opposite(figure.Color()))
}

func (b *Board) IsSquareAvailableForMoveLongRangePiece(coordinates coords.Coordinates, figure piece.Piece) bool {
	figureCoords := figure.Coordinates()
	return !b.HasPieceOnWayForLongRangePiece(figureCoords, coordinates)
}
