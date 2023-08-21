package piece

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

type King struct {
	color       color.Color
	coordinates coords.Coordinates
}

func NewKing(color color.Color, coordinates coords.Coordinates) *King {
	return &King{color: color, coordinates: coordinates}
}

func (k *King) Name() string {
	return "King"
}

func (k *King) Color() color.Color {
	return k.color
}

func (k *King) Coordinates() coords.Coordinates {
	return k.coordinates
}

func (k *King) SetCoordinates(coordinates coords.Coordinates) {
	k.coordinates = coordinates
}

func (k *King) SetColor(color color.Color) {
	k.color = color
}

func (k *King) Shifts() map[coords.CoordinatesShift]bool {
	return nil
}
