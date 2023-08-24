package board

import (
	"chess/src/game/coords"
	"chess/src/game/piece"
)

type Board struct {
	Pieces map[coords.Coordinates]piece.Piece
}

func NewBoard() *Board {
	return &Board{Pieces: map[coords.Coordinates]piece.Piece{}}
}
