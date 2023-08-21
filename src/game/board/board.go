package board

import (
	"chess/src/game/color"
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

func (b *Board) IsSquareDark(coordinates coords.Coordinates) bool {
	return (int(coordinates.File)+int(coordinates.Rank))%2 == 0
}

func (b *Board) IsSquareEmpty(coordinates coords.Coordinates) bool {
	_, ok := b.GetPiece(coordinates)
	return !ok
}

func (b *Board) GetPiece(coordinates coords.Coordinates) (piece.Piece, bool) {
	curPiece, ok := b.Pieces[coordinates]
	return curPiece, ok
}

func (b *Board) SetupDefaultPiecesPositions() {
	// set pawns
	for r := coords.A; r <= coords.H; r++ {
		fWhite := coords.File(2)
		coordsWhite := coords.NewCoordinates(r, fWhite)
		pawnWhite := piece.NewPawn(color.White, coordsWhite)

		fBlack := coords.File(7)
		coordsBlack := coords.NewCoordinates(r, fBlack)
		pawnBlack := piece.NewPawn(color.Black, coordsBlack)

		b.SetPiece(coordsWhite, pawnWhite)
		b.SetPiece(coordsBlack, pawnBlack)
	}

	// set rocks
	coordsBlack := coords.NewCoordinates('A', 8)
	b.SetPiece(coordsBlack, piece.NewRock(color.Black, coordsBlack))
	coordsBlack = coords.NewCoordinates('H', 8)
	b.SetPiece(coordsBlack, piece.NewRock(color.Black, coordsBlack))

	coordsWhite := coords.NewCoordinates('A', 1)
	b.SetPiece(coordsWhite, piece.NewRock(color.White, coordsWhite))
	coordsWhite = coords.NewCoordinates('H', 1)
	b.SetPiece(coordsWhite, piece.NewRock(color.White, coordsWhite))

	// set knights
	coordsBlack = coords.NewCoordinates('B', 8)
	b.SetPiece(coordsBlack, piece.NewKnight(color.Black, coordsBlack))
	coordsBlack = coords.NewCoordinates('G', 8)
	b.SetPiece(coordsBlack, piece.NewKnight(color.Black, coordsBlack))

	coordsWhite = coords.NewCoordinates('B', 1)
	b.SetPiece(coordsWhite, piece.NewKnight(color.White, coordsWhite))
	coordsWhite = coords.NewCoordinates('G', 1)
	b.SetPiece(coordsWhite, piece.NewKnight(color.White, coordsWhite))

	// set bishops
	coordsBlack = coords.NewCoordinates('C', 8)
	b.SetPiece(coordsBlack, piece.NewBishop(color.Black, coordsBlack))
	coordsBlack = coords.NewCoordinates('F', 8)
	b.SetPiece(coordsBlack, piece.NewBishop(color.Black, coordsBlack))

	coordsWhite = coords.NewCoordinates('C', 1)
	b.SetPiece(coordsWhite, piece.NewBishop(color.White, coordsWhite))
	coordsWhite = coords.NewCoordinates('F', 1)
	b.SetPiece(coordsWhite, piece.NewBishop(color.White, coordsWhite))
	// set queens
	coordsBlack = coords.NewCoordinates('D', 8)
	b.SetPiece(coordsBlack, piece.NewQueen(color.Black, coordsBlack))

	coordsWhite = coords.NewCoordinates('D', 1)
	b.SetPiece(coordsWhite, piece.NewQueen(color.White, coordsWhite))

	// set kings
	coordsBlack = coords.NewCoordinates('E', 8)
	b.SetPiece(coordsBlack, piece.NewKing(color.Black, coordsBlack))

	coordsWhite = coords.NewCoordinates('E', 1)
	b.SetPiece(coordsWhite, piece.NewKing(color.White, coordsWhite))
}

func (b *Board) IsSquareAvailableForMoveSimple(coordinates coords.Coordinates, curPiece piece.Piece) bool {
	if b.IsSquareEmpty(coordinates) {
		return true
	}
	otherPiece, _ := b.GetPiece(coordinates)
	return curPiece.Color() != otherPiece.Color() && otherPiece.Name() != "King"
}

func (b *Board) IsSquareAvailableForMove(coordinates coords.Coordinates, curPiece piece.Piece) bool {
	if !b.IsSquareAvailableForMoveSimple(coordinates, curPiece) {
		return false
	}
	return true
}

func (b *Board) AvailableMoves(curPiece piece.Piece) map[coords.Coordinates]bool {
	availableMoves := map[coords.Coordinates]bool{}
	shifts := curPiece.Shifts()
	for shift, _ := range shifts {
		if curPiece.Coordinates().CanShift(shift) {
			newCoordinates := curPiece.Coordinates().Shift(shift)
			if b.IsSquareAvailableForMove(newCoordinates, curPiece) {
				availableMoves[newCoordinates] = true
			}
		}
	}
	return availableMoves
}
