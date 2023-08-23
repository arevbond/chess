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

func (b *Board) SquaresBetween(from, to coords.Coordinates) []coords.Coordinates {
	coordinates := make([]coords.Coordinates, 0)
	if from.Rank == to.Rank { // ход по вертикали
		if from.File > to.File {
			for f := from.File - 1; f >= to.File; f-- {
				newCoords := coords.NewCoordinates(from.Rank, f)
				coordinates = append(coordinates, newCoords)
			}
		} else {
			for f := from.File + 1; f <= to.File; f++ {
				newCoords := coords.NewCoordinates(from.Rank, f)
				coordinates = append(coordinates, newCoords)
			}
		}
	} else if from.File == to.File { // ход по горизонтали
		if from.Rank > to.Rank {
			for r := from.Rank - 1; r >= to.Rank; r-- {
				newCoords := coords.NewCoordinates(r, from.File)
				coordinates = append(coordinates, newCoords)
			}
		} else {
			for r := from.Rank + 1; r <= to.Rank; r++ {
				newCoords := coords.NewCoordinates(r, from.File)
				coordinates = append(coordinates, newCoords)
			}
		}
	} else { // ход по диагонали
		if from.Rank > to.Rank && from.File > to.File {
			for r, f := from.Rank-1, from.File-1; r >= to.Rank && f >= to.File; r, f = r-1, f-1 {
				newCoords := coords.NewCoordinates(r, f)
				coordinates = append(coordinates, newCoords)
			}
		} else if from.Rank < to.Rank && from.File < to.File {
			for r, f := from.Rank+1, from.File+1; r <= to.Rank && f <= to.File; r, f = r+1, f+1 {
				newCoords := coords.NewCoordinates(r, f)
				coordinates = append(coordinates, newCoords)
			}
		} else if from.Rank > to.Rank && from.File < to.File {
			for r, f := from.Rank-1, from.File+1; r >= to.Rank && f <= to.File; r, f = r-1, f+1 {
				newCoords := coords.NewCoordinates(r, f)
				coordinates = append(coordinates, newCoords)
			}
		} else if from.Rank < to.Rank && from.File > to.File {
			for r, f := from.Rank+1, from.File-1; r <= to.Rank && f >= to.File; r, f = r+1, f-1 {
				newCoords := coords.NewCoordinates(r, f)
				coordinates = append(coordinates, newCoords)
			}
		}
	}
	return coordinates
}

func (b *Board) HasPieceOnWay(from, to coords.Coordinates) bool {
	var flag bool
	coordsBetween := b.SquaresBetween(from, to)
	ourPiece, _ := b.GetPiece(from)
	ourPieceColor := ourPiece.Color()
	for _, coordinates := range coordsBetween {
		if otherPiece, ok := b.GetPiece(coordinates); ok {
			otherPieceColor := otherPiece.Color()
			if flag || otherPiece.Name() == "King" || otherPieceColor == ourPieceColor {
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

func (b *Board) IsSquareAvailableForMove(coordinates coords.Coordinates, figure piece.Piece) bool {
	if !b.IsSquareAvailableForMoveSimple(coordinates, figure) {
		return false
	}
	figureCoords := figure.Coordinates()
	return !b.HasPieceOnWay(figureCoords, coordinates)
}

func (b *Board) AvailableMoves(figure piece.Piece) map[coords.Coordinates]bool {
	// TODO: add for pawn and king
	availableMoves := map[coords.Coordinates]bool{}
	shifts := figure.Shifts()
	for shift, _ := range shifts {
		if figure.Coordinates().CanShift(shift) {
			newCoordinates := figure.Coordinates().Shift(shift)
			if b.IsSquareAvailableForMove(newCoordinates, figure) {
				availableMoves[newCoordinates] = true
			}
		}
	}
	return availableMoves
}
