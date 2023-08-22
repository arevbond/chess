package piece

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

type Rock struct {
	color       color.Color
	coordinates coords.Coordinates
}

func NewRock(color color.Color, coordinates coords.Coordinates) *Rock {
	return &Rock{color: color, coordinates: coordinates}
}

func (r *Rock) Name() string {
	return "Rock"
}

func (r *Rock) Color() color.Color {
	return r.color
}

func (r *Rock) Coordinates() coords.Coordinates {
	return r.coordinates
}

func (r *Rock) SetCoordinates(coordinates coords.Coordinates) {
	r.coordinates = coordinates
}

func (r *Rock) SetColor(color color.Color) {
	r.color = color
}

func (r *Rock) Shifts() map[coords.CoordinatesShift]bool {
	shifts := [][]int{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0},
		{-1, 0}, {-2, 0}, {-3, 0}, {-4, 0}, {-5, 0}, {-6, 0}, {-7, 0},
		{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7},
		{0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5}, {0, -6}, {0, -7}}
	coordsShifts := coords.CalculateCoordinatesShift(shifts, r.Coordinates())
	return coordsShifts
}
