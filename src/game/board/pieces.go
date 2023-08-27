package board

import (
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/move"
	"chess/src/game/piece"
)

func (b *Board) SetPiece(coordinates coords.Coordinates, piece piece.Piece) {
	piece.SetCoordinates(coordinates)
	b.Pieces[coordinates] = piece
}

func (b *Board) GetPiece(coordinates coords.Coordinates) (piece.Piece, bool) {
	figure, ok := b.Pieces[coordinates]
	return figure, ok
}

func (b *Board) PiecesByColor(figureColor color.Color) []piece.Piece {
	pieces := make([]piece.Piece, 0)
	for _, figure := range b.Pieces {
		if figure.Color() == figureColor {
			pieces = append(pieces, figure)
		}
	}
	return pieces
}

func (b *Board) RemovePiece(coordinates coords.Coordinates) {
	delete(b.Pieces, coordinates)
}

func (b *Board) MovePiece(from, to coords.Coordinates) {
	figure, _ := b.GetPiece(from)
	b.RemovePiece(from)
	b.SetPiece(to, figure)

	curMove := move.NewMove(from, to, figure, figure.Color())
	b.Moves = append(b.Moves, curMove)
}
