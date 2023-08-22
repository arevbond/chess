package render

import (
	"chess/src/game/board"
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
	"fmt"
)

const (
	AnsiReset string = "\u001B[0m"

	AnsiWhitePieceColor string = "\u001B[97m"
	AnsiBlackPieceColor string = "\u001B[30m"

	AnsiWhiteSquareBackground       string = "\u001B[47m"
	AnsiBlackSquareBackground       string = "\u001B[0;100m"
	AnsiHighlightedSquareBackground        = "\u001B[45m"
)

func GetSpriteForEmptySquare(coordinates coords.Coordinates) string {
	return ColorizeSprite("   ", color.White, IsSquareDark(coordinates), false)
}

func IsSquareDark(coordinate coords.Coordinates) bool {
	return (int(coordinate.File)+int(coordinate.Rank))%2 == 0
}

func ColorizeSprite(sprite string, pieceColor color.Color, isSquareDark bool, isHighLighted bool) string {
	var pColor, backgroundColor string

	if pieceColor == color.White {
		pColor = AnsiWhitePieceColor
	} else {
		pColor = AnsiBlackPieceColor
	}

	if isHighLighted {
		backgroundColor = AnsiHighlightedSquareBackground
	} else if isSquareDark {
		backgroundColor = AnsiBlackSquareBackground
	} else {
		backgroundColor = AnsiWhiteSquareBackground
	}
	result := backgroundColor + pColor + sprite
	return result + AnsiReset
}

func GetPieceSprite(curPiece piece.Piece) string {
	pieceColor := curPiece.Color()
	pieceCoords := curPiece.Coordinates()
	sprite := SelectUnicodeSpriteForPiece(curPiece)
	return ColorizeSprite(sprite, pieceColor, IsSquareDark(pieceCoords), false)
}

func SelectUnicodeSpriteForPiece(curPiece piece.Piece) string {
	var sprite string
	switch curPiece.(type) {
	case *piece.Pawn:
		sprite = " ♟ "
	case *piece.Knight:
		sprite = " ♞ "
	case *piece.Bishop:
		sprite = " ♝ "
	case *piece.Rock:
		sprite = " ♜ "
	case *piece.Queen:
		sprite = " ♛ "
	case *piece.King:
		sprite = " ♚ "
	}
	return sprite
}

func RenderBoardForWhite(curBoard *board.Board, withMoves bool) {
	for f := coords.File(8); f >= 1; f-- {
		line := ""
		for r := coords.A; r <= coords.H; r++ {
			coordinates := coords.NewCoordinates(r, f)
			if curBoard.IsSquareEmpty(coordinates) {
				line += GetSpriteForEmptySquare(coordinates)
			} else {
				curPiece, ok := curBoard.GetPiece(coordinates)
				if ok {
					line += GetPieceSprite(curPiece)
				}
			}
		}
		fmt.Println(line)
	}
}

func RenderBoardForBlack(curBoard *board.Board) {
	for f := coords.File(1); f <= 8; f++ {
		line := ""
		for r := coords.H; r >= coords.A; r-- {
			coordinates := coords.NewCoordinates(r, f)
			if curBoard.IsSquareEmpty(coordinates) {
				line += GetSpriteForEmptySquare(coordinates)
			} else {
				curPiece, ok := curBoard.GetPiece(coordinates)
				if ok {
					line += GetPieceSprite(curPiece)
				}
			}
		}
		fmt.Println(line)
	}
}

func RenderBoard(curColor color.Color, curBoard *board.Board) {
	if curColor == color.White {
		RenderBoardForWhite(curBoard, false)
	} else if curColor == color.Black {
		RenderBoardForBlack(curBoard)
	}
}
