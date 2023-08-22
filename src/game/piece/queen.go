package piece

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

type Queen struct {
	color       color.Color
	coordinates coords.Coordinates
}

func NewQueen(color color.Color, coordinates coords.Coordinates) *Queen {
	return &Queen{color: color, coordinates: coordinates}
}

func (q *Queen) Name() string {
	return "Queen"
}

func (q *Queen) Color() color.Color {
	return q.color
}

func (q *Queen) Coordinates() coords.Coordinates {
	return q.coordinates
}

func (q *Queen) SetCoordinates(coordinates coords.Coordinates) {
	q.coordinates = coordinates
}

func (q *Queen) SetColor(color color.Color) {
	q.color = color
}

func (q *Queen) Shifts() map[coords.CoordinatesShift]bool {
	shiftsRock := [][]int{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0},
		{-1, 0}, {-2, 0}, {-3, 0}, {-4, 0}, {-5, 0}, {-6, 0}, {-7, 0},
		{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7},
		{0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5}, {0, -6}, {0, -7}}
	shiftsBishop := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7},
		{-1, -1}, {-2, -2}, {-3, -3}, {-4, -4}, {-5, -5}, {-6, -6}, {-7, -7},
		{1, -1}, {2, -2}, {3, -3}, {4, -4}, {5, -5}, {6, -6}, {7, -7},
		{-1, 1}, {-2, 2}, {-3, 3}, {-4, 4}, {-5, 5}, {-6, 6}, {-7, 7}}
	shifts := [][]int{}
	shifts = append(shifts, shiftsRock...)
	shifts = append(shifts, shiftsBishop...)
	coordsShifts := coords.CalculateCoordinatesShift(shifts, q.Coordinates())
	return coordsShifts
}
