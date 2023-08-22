package game

import (
	"chess/src/game/board"
	"chess/src/game/color"
	"chess/src/game/inputCoords"
	"chess/src/game/render"
)

func GameLoop(board *board.Board) {
	var isWhiteToMove bool = true
	var curColor color.Color
	for {
		if isWhiteToMove {
			curColor = color.White
		} else {
			curColor = color.Black
		}

		//render
		render.RenderBoard(curColor, board)

		// input coords
		pieceCoords, curPiece := inputCoords.InputCoordsOwnPieceCanMove(curColor, board)
		targetCoords := inputCoords.InputCoordsYourPieceToMove(curPiece, board)

		// make move
		board.MovePiece(pieceCoords, targetCoords)

		// pass move
		isWhiteToMove = !isWhiteToMove
	}
}
