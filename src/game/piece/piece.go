package piece

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

type Piece interface {
	Color() color.Color
	Coordinates() coords.Coordinates
	SetCoordinates(coordinates coords.Coordinates)
	SetColor(color color.Color)

	//Shifts() map[*coords.CoordinatesShift]bool
	//AvailableMoves(board.Board) map[*coords.Coordinates]bool
}
