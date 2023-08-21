package piece

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

func NewBishop(color color.Color, coordinates coords.Coordinates) *Bishop {
	return &Bishop{color: color, coordinates: coordinates}
}

type Bishop struct {
	color       color.Color
	coordinates coords.Coordinates
}

func (b *Bishop) Color() color.Color {
	return b.color
}

func (b *Bishop) Coordinates() coords.Coordinates {
	return b.coordinates
}

func (b *Bishop) SetCoordinates(coordinates coords.Coordinates) {
	b.coordinates = coordinates
}

func (b *Bishop) SetColor(color color.Color) {
	b.color = color
}

//
//func (b *Bishop) AvailableMoves(curBoard board.Board) map[*coords.Coordinates]bool {
//	availableMoves := map[*coords.Coordinates]bool{}
//	shifts := b.Shifts()
//	for shift, _ := range shifts {
//		if b.coordinates.CanShift(shift) {
//			newCoordinates := b.coordinates.Shift(shift)
//			if b.IsSquareAvailableForMove(newCoordinates, curBoard) {
//				availableMoves[newCoordinates] = true
//			}
//		}
//
//	}
//	return availableMoves
//}
//
//func (b *Bishop) Shifts() map[*coords.CoordinatesShift]bool {
//	return map[*coords.CoordinatesShift]bool{}
//}
//
//func (b *Bishop) IsSquareAvailableForMove(coordinates *coords.Coordinates, curBoard board.Board) bool {
//	if !curBoard.IsSquareAvailableForMoveSimple(coordinates, b) {
//		return false
//	}
//	return true
//}
