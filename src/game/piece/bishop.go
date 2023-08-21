package piece

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

type Bishop struct {
	color       color.Color
	coordinates coords.Coordinates
}

func NewBishop(color color.Color, coordinates coords.Coordinates) *Bishop {
	return &Bishop{color: color, coordinates: coordinates}
}

func (b *Bishop) Name() string {
	return "Bishop"
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

func (b *Bishop) Shifts() map[coords.CoordinatesShift]bool {
	return map[coords.CoordinatesShift]bool{}
}
