package move

import (
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
)

type Move struct {
	from, to coords.Coordinates
	figure   piece.Piece
	color    color.Color
}

func NewMove(from, to coords.Coordinates, figure piece.Piece, color color.Color) Move {
	return Move{from: from, to: to, figure: figure, color: color}
}
