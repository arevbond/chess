package piece

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

type Knight struct {
	color       color.Color
	coordinates coords.Coordinates
}

func NewKnight(color color.Color, coordinates coords.Coordinates) *Knight {
	return &Knight{color: color, coordinates: coordinates}
}

func (k *Knight) Name() string {
	return "Knight"
}

func (k *Knight) Color() color.Color {
	return k.color
}

func (k *Knight) Coordinates() coords.Coordinates {
	return k.coordinates
}

func (k *Knight) SetCoordinates(coordinates coords.Coordinates) {
	k.coordinates = coordinates
}

func (k *Knight) SetColor(color color.Color) {
	k.color = color
}

func (k *Knight) Shifts() map[coords.CoordinatesShift]bool {
	shifts := [][]int{{2, 1}, {1, 2}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}}
	coordsShifts := coords.CalculateCoordinatesShift(shifts, k.Coordinates())
	return coordsShifts
}
