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

func (b *Board) SetPiece(coordinates coords.Coordinates, piece piece.Piece) {
	piece.SetCoordinates(coordinates)
	b.Pieces[coordinates] = piece
}

func (b *Board) GetPiece(coordinates coords.Coordinates) (piece.Piece, bool) {
	figure, ok := b.Pieces[coordinates]
	return figure, ok
}

func (b *Board) RemovePiece(coordinates coords.Coordinates) {
	delete(b.Pieces, coordinates)
}

func (b *Board) MovePiece(from, to coords.Coordinates) {
	figure, _ := b.GetPiece(from)
	b.RemovePiece(from)
	b.SetPiece(to, figure)
}

func (b *Board) IsSquareDark(coordinates coords.Coordinates) bool {
	return (int(coordinates.File)+int(coordinates.Rank))%2 == 0
}

func (b *Board) IsSquareEmpty(coordinates coords.Coordinates) bool {
	_, ok := b.GetPiece(coordinates)
	return !ok
}
