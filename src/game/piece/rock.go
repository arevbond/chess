package piece

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

type Rock struct {
	color       color.Color
	coordinates coords.Coordinates
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

func NewRock(color color.Color, coordinates coords.Coordinates) *Rock {
	return &Rock{color: color, coordinates: coordinates}
}
