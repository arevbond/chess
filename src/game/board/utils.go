package board

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

func (b *Board) HasPieceOnWay(from, to coords.Coordinates) bool {
	coordsBetween := SquaresBetween(from, to)
	for _, coordinates := range coordsBetween {
		if _, ok := b.GetPiece(coordinates); ok {
			return true
		}
	}
	return false
}

func (b *Board) HasPieceOnWayForLongRangePiece(from, to coords.Coordinates) bool {
	var flag bool
	coordsBetween := SquaresBetween(from, to)
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

func VerticalSquaresBetween(from, to coords.Coordinates) []coords.Coordinates {
	coordinates := make([]coords.Coordinates, 0)
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
	return coordinates
}

func HorizontalSquaresBetween(from, to coords.Coordinates) []coords.Coordinates {
	coordinates := make([]coords.Coordinates, 0)
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
	return coordinates
}

func DiagonalSquaresBetween(from, to coords.Coordinates) []coords.Coordinates {
	coordinates := make([]coords.Coordinates, 0)
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
	return coordinates
}

func SquaresBetween(from, to coords.Coordinates) []coords.Coordinates {
	coordinates := make([]coords.Coordinates, 0)
	if from.Rank == to.Rank { // ход по вертикали
		coordinates = VerticalSquaresBetween(from, to)
	} else if from.File == to.File { // ход по горизонтали
		coordinates = HorizontalSquaresBetween(from, to)
	} else { // ход по диагонали
		coordinates = DiagonalSquaresBetween(from, to)
	}
	return coordinates
}

func (b *Board) IsKingInCheck(kingColor color.Color) bool {
	king := b.GetKing(kingColor)
	return b.IsSquareAttackedByColor(king.Coordinates(), color.Opposite(kingColor))
}
