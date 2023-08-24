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
	AnsiGreenColor      string = "\u001B[92m"

	AnsiWhiteSquareBackground       string = "\u001B[47m"
	AnsiBlackSquareBackground       string = "\u001B[0;100m"
	AnsiHighlightedSquareBackground        = "\u001B[42m"
)

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
