package game

import (
	"chess/src/game/board"
	"chess/src/game/color"
	"chess/src/game/inputCoords"
	"chess/src/game/render"
)

func GameLoop(board *board.Board) {
	var isWhiteToMove bool = true
	var pieceColor color.Color
	for {
		if isWhiteToMove {
			pieceColor = color.White
		} else {
			pieceColor = color.Black
		}

		//render
		render.RenderBoard(pieceColor, board)

		// input coords
		figureCoords, figure := inputCoords.InputCoordsOwnPieceCanMove(pieceColor, board)

		// render board with available moves
		render.RenderBoardWithAvailablePieceMoves(pieceColor, board, figure)

		targetCoords := inputCoords.InputCoordsYourPieceToMove(figure, board)

		// make move
		board.MovePiece(figureCoords, targetCoords)

		// pass move
		isWhiteToMove = !isWhiteToMove
	}
}
