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

func RenderBoardForWhite(gameBoard *board.Board) {
	for f := coords.File(8); f >= 1; f-- {
		line := ""
		for r := coords.A; r <= coords.H; r++ {
			coordinates := coords.NewCoordinates(r, f)
			if gameBoard.IsSquareEmpty(coordinates) {
				line += GetSpriteForEmptySquare(coordinates, false)
			} else {
				figure, ok := gameBoard.GetPiece(coordinates)
				if ok {
					line += GetPieceSprite(figure, false)
				}
			}
		}
		fmt.Println(line)
	}
}

func RenderBoardForBlack(gameBoard *board.Board) {
	for f := coords.File(1); f <= 8; f++ {
		line := ""
		for r := coords.H; r >= coords.A; r-- {
			coordinates := coords.NewCoordinates(r, f)
			if gameBoard.IsSquareEmpty(coordinates) {
				line += GetSpriteForEmptySquare(coordinates, false)
			} else {
				curPiece, ok := gameBoard.GetPiece(coordinates)
				if ok {
					line += GetPieceSprite(curPiece, false)
				}
			}
		}
		fmt.Println(line)
	}
}

func RenderBoard(figureColor color.Color, board *board.Board) {
	if figureColor == color.White {
		RenderBoardForWhite(board)
	} else if figureColor == color.Black {
		RenderBoardForBlack(board)
	}
}

func RenderBoardWithAvailablePieceMoves(figureColor color.Color, board *board.Board, figure piece.Piece) {
	if figureColor == color.White {
		RenderBoardForWhiteWithAvailablePieceMoves(board, figure)
	} else if figureColor == color.Black {
		RenderBoardForBlackWithAvailablePieceMoves(board, figure)
	}
}

func RenderBoardForWhiteWithAvailablePieceMoves(gameBoard *board.Board, figure piece.Piece) {
	availableMoves := gameBoard.AvailableMoves(figure)
	for f := coords.File(8); f >= 1; f-- {
		line := ""
		for r := coords.A; r <= coords.H; r++ {
			coordinates := coords.NewCoordinates(r, f)
			var isHighlight bool
			if _, ok := availableMoves[coordinates]; ok {
				isHighlight = true
			}
			if gameBoard.IsSquareEmpty(coordinates) {
				line += GetSpriteForEmptySquare(coordinates, isHighlight)
			} else {
				curPiece, ok := gameBoard.GetPiece(coordinates)
				if ok {
					line += GetPieceSprite(curPiece, isHighlight)
				}
			}
		}
		fmt.Println(line)
	}

}

func RenderBoardForBlackWithAvailablePieceMoves(gameBoard *board.Board, figure piece.Piece) {
	availableMoves := gameBoard.AvailableMoves(figure)
	for f := coords.File(1); f <= 8; f++ {
		line := ""
		for r := coords.H; r >= coords.A; r-- {
			coordinates := coords.NewCoordinates(r, f)
			var isHighlight bool
			if _, ok := availableMoves[coordinates]; ok {
				isHighlight = true
			}
			if gameBoard.IsSquareEmpty(coordinates) {
				line += GetSpriteForEmptySquare(coordinates, isHighlight)
			} else {
				curPiece, ok := gameBoard.GetPiece(coordinates)
				if ok {
					line += GetPieceSprite(curPiece, isHighlight)
				}
			}
		}
		fmt.Println(line)
	}
}
