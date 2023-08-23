package render

import (
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
)

func GetSpriteForEmptySquare(coordinates coords.Coordinates, isHighLighted bool) string {
	return ColorizeSprite("   ", color.White, IsSquareDark(coordinates), isHighLighted)
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

func GetPieceSprite(figure piece.Piece, isHighLighted bool) string {
	pieceColor := figure.Color()
	pieceCoords := figure.Coordinates()
	sprite := SelectUnicodeSpriteForPiece(figure)
	return ColorizeSprite(sprite, pieceColor, IsSquareDark(pieceCoords), isHighLighted)
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
