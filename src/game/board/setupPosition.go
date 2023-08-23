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

func (b *Board) SetupPositionFromFEN(fen string) {
	// fen: "rnbqkbn1/pppppppp/8/7r/8/8/PPPPPPPP/RNBQKBN1 w Qq - 0 1"
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
				rank += coords.Rank(digit)
				continue
			}
			curCoords := coords.NewCoordinates(rank, file)
			curPiece := PieceFromFenChar(symbol, curCoords)
			b.SetPiece(curCoords, curPiece)
			rank++
		}
	}
}

func PieceFromFenChar(symbol rune, coordinates coords.Coordinates) piece.Piece {
	var figure piece.Piece
	var figureColor color.Color

	figureColor = color.White
	if unicode.IsLower(symbol) {
		figureColor = color.Black
	}
	symbol = unicode.ToUpper(symbol)

	switch symbol {
	case 'R':
		figure = piece.NewRock(figureColor, coordinates)
	case 'N':
		figure = piece.NewKnight(figureColor, coordinates)
	case 'B':
		figure = piece.NewBishop(figureColor, coordinates)
	case 'Q':
		figure = piece.NewQueen(figureColor, coordinates)
	case 'K':
		figure = piece.NewKing(figureColor, coordinates)
	case 'P':
		figure = piece.NewPawn(figureColor, coordinates)
	default:
		log.Fatalf("Unknow fen symbol for piece - %q", symbol)
	}
	return figure
}

func (b *Board) SetupDefaultPiecesPositions() {
	startPosition := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	b.SetupPositionFromFEN(startPosition)
}
