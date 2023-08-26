package render

import (
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
)

const (
	Circle       string = " ● "
	PawnSprite   string = " ♟ "
	KnightSprite string = " ♞ "
	BishopSprite string = " ♝ "
	RockSprite   string = " ♜ "
	QueenSprite  string = " ♛ "
	KingSprite   string = " ♚ "
)

func GetSpriteForEmptySquare(coordinates coords.Coordinates, isHighLighted bool) string {
	return ColorizeSprite("   ", color.White, IsSquareDark(coordinates), isHighLighted)
}

func IsSquareDark(coordinate coords.Coordinates) bool {
	return (int(coordinate.File)+int(coordinate.Rank))%2 == 0
}

func ColorizeSprite(sprite string, pieceColor color.Color, isSquareDark bool, isHighLighted bool) string {
	var pColor, backgroundColor string
	var isHighlightedBgSquare bool

	if isHighLighted && sprite != "   " {
		isHighlightedBgSquare = true
	}

	if isHighLighted && !isHighlightedBgSquare {
		sprite = Circle
		pColor = AnsiGreenColor
	} else if pieceColor == color.White {
		pColor = AnsiWhitePieceColor
	} else {
		pColor = AnsiBlackPieceColor
	}

	if isHighlightedBgSquare {
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

func SelectUnicodeSpriteForPiece(figure piece.Piece) string {
	var sprite string
	switch figure.(type) {
	case *piece.Pawn:
		sprite = PawnSprite
	case *piece.Knight:
		sprite = KnightSprite
	case *piece.Bishop:
		sprite = BishopSprite
	case *piece.Rock:
		sprite = RockSprite
	case *piece.Queen:
		sprite = QueenSprite
	case *piece.King:
		sprite = KingSprite
	}
	return sprite
}
