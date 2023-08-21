package piece

import (
	"chess/src/game/color"
	"chess/src/game/coords"
)

type Pawn struct {
	color       color.Color
	coordinates coords.Coordinates
}

func NewPawn(color color.Color, coordinates coords.Coordinates) *Pawn {
	return &Pawn{color: color, coordinates: coordinates}
}

func (p *Pawn) Name() string {
	return "Pawn"
}

func (p *Pawn) Color() color.Color {
	return p.color
}

func (p *Pawn) Coordinates() coords.Coordinates {
	return p.coordinates
}

func (p *Pawn) SetCoordinates(coordinates coords.Coordinates) {
	p.coordinates = coordinates
}

func (p *Pawn) SetColor(color color.Color) {
	p.color = color
}

func (p *Pawn) Shifts() map[coords.CoordinatesShift]bool {
	return nil
}
