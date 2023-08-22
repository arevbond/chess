package board

import (
	"chess/src/game/coords"
	"chess/src/game/piece"
)

func (b *Board) IsSquareAvailableForMoveSimple(coordinates coords.Coordinates, curPiece piece.Piece) bool {
	// проверяет пустая ли клетка, если нет, стоит ли на ней вражеская фигура, кроме короля
	if b.IsSquareEmpty(coordinates) {
		return true
	}
	otherPiece, _ := b.GetPiece(coordinates)
	return curPiece.Color() != otherPiece.Color() && otherPiece.Name() != "King"
}

func (b *Board) CheckPieceOnWay(from, to coords.Coordinates) bool {
	// TODO: завершить функцию проверки фигур на пути
	return false
}

func (b *Board) HasPieceOnWay(coordinatesToMove coords.Coordinates, curPiece piece.Piece) bool {
	// проверка на то, чтобы на пути фигуры не было других фигур
	// то есть чтобы фигура не перепрыгивала через другие фигуры
	if curPiece.Name() == "Knight" {
		return false
	}
	curCoords := curPiece.Coordinates()
	return b.CheckPieceOnWay(curCoords, coordinatesToMove)
}

func (b *Board) IsSquareAvailableForMove(coordinates coords.Coordinates, curPiece piece.Piece) bool {
	if !b.IsSquareAvailableForMoveSimple(coordinates, curPiece) {
		return false
	}
	//if b.HasPieceOnWay()
	return true
}

func (b *Board) AvailableMoves(curPiece piece.Piece) map[coords.Coordinates]bool {
	availableMoves := map[coords.Coordinates]bool{}
	shifts := curPiece.Shifts()
	for shift, _ := range shifts {
		if curPiece.Coordinates().CanShift(shift) {
			newCoordinates := curPiece.Coordinates().Shift(shift)
			if b.IsSquareAvailableForMove(newCoordinates, curPiece) {
				availableMoves[newCoordinates] = true
			}
		}
	}
	return availableMoves
}
