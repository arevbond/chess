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
