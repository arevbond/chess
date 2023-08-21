package piece

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

type Queen struct {
	color       color.Color
	coordinates coords.Coordinates
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

func NewQueen(color color.Color, coordinates coords.Coordinates) *Queen {
	return &Queen{color: color, coordinates: coordinates}
}
