package board

import (
	"chess/src/game/coords"
	"chess/src/game/move"
	"chess/src/game/piece"
)

type Board struct {
	Pieces map[coords.Coordinates]piece.Piece
	Moves  []move.Move
}

func NewBoard() *Board {
	return &Board{Pieces: map[coords.Coordinates]piece.Piece{}, Moves: []move.Move{}}
}
