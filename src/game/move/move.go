package move

import (
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
)

type Move struct {
	From, To    coords.Coordinates
	Figure      piece.Piece
	FigureColor color.Color
}

func NewMove(from, to coords.Coordinates, figure piece.Piece, color color.Color) Move {
	return Move{From: from, To: to, Figure: figure, FigureColor: color}
}
