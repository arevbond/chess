package board

import (
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
	"log"
	"strconv"
	"strings"
	"unicode"
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
	curPiece, ok := b.Pieces[coordinates]
	return curPiece, ok
}

func (b *Board) RemovePiece(coordinates coords.Coordinates) {
	delete(b.Pieces, coordinates)
}

func (b *Board) MovePiece(from, to coords.Coordinates) {
	curPiece, _ := b.GetPiece(from)
	b.RemovePiece(from)
	b.SetPiece(to, curPiece)
}

func (b *Board) IsSquareDark(coordinates coords.Coordinates) bool {
	return (int(coordinates.File)+int(coordinates.Rank))%2 == 0
}

func (b *Board) IsSquareEmpty(coordinates coords.Coordinates) bool {
	_, ok := b.GetPiece(coordinates)
	return !ok
}

func (b *Board) IsSquareAvailableForMoveSimple(coordinates coords.Coordinates, curPiece piece.Piece) bool {
	// проверяет пустая ли клетка, если нет, стоит ли на ней вражеская фигура, кроме короля
	if b.IsSquareEmpty(coordinates) {
		return true
	}
	otherPiece, _ := b.GetPiece(coordinates)
	return curPiece.Color() != otherPiece.Color() && otherPiece.Name() != "King"
}

func (b *Board) CheckPieceOnWay(from, to coords.Coordinates) bool {
	// TODO: завершить функцию проверки фигур на пути
	return false
}

func (b *Board) HasPieceOnWay(coordinatesToMove coords.Coordinates, curPiece piece.Piece) bool {
	// проверка на то, чтобы на пути фигуры не было других фигур
	// то есть чтобы фигура не перепрыгивала через другие фигуры
	if curPiece.Name() == "Knight" {
		return false
	}
	curCoords := curPiece.Coordinates()
	return b.CheckPieceOnWay(curCoords, coordinatesToMove)
}

func (b *Board) IsSquareAvailableForMove(coordinates coords.Coordinates, curPiece piece.Piece) bool {
	if !b.IsSquareAvailableForMoveSimple(coordinates, curPiece) {
		return false
	}
	//if b.HasPieceOnWay()
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

func (b *Board) SetupPositionFromFEN(fen string) {
	// rnbqkbn1/pppppppp/8/7r/8/8/PPPPPPPP/RNBQKBN1 w Qq - 0 1
	parts := strings.Split(fen, " ")
	piecePositions := parts[0]
	fenRows := strings.Split(piecePositions, "/")
	for i := 0; i < len(fenRows); i++ {
		file := coords.File(8 - i)
		rank := coords.A
		for j := 0; j < len(fenRows[i]); j++ {
			symbol := rune(fenRows[i][j])
			if unicode.IsDigit(symbol) {
				digit, err := strconv.Atoi(string(symbol))
				if err != nil {
					log.Fatal(err)
				}
				rank += coords.Rank(digit - 1)
				continue
			}
			var curColor color.Color
			var curPiece piece.Piece
			curCoords := coords.NewCoordinates(rank, file)
			curColor = color.White
			if unicode.IsLower(symbol) {
				curColor = color.Black
			}
			symbol = unicode.ToUpper(symbol)

			switch symbol {
			case 'R':
				curPiece = piece.NewRock(curColor, curCoords)
			case 'N':
				curPiece = piece.NewKnight(curColor, curCoords)
			case 'B':
				curPiece = piece.NewBishop(curColor, curCoords)
			case 'Q':
				curPiece = piece.NewQueen(curColor, curCoords)
			case 'K':
				curPiece = piece.NewKing(curColor, curCoords)
			case 'P':
				curPiece = piece.NewPawn(curColor, curCoords)
			}
			b.SetPiece(curCoords, curPiece)
			rank++
		}
	}
}

func (b *Board) SetupDefaultPiecesPositions() {
	startPosition := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	b.SetupPositionFromFEN(startPosition)
}
