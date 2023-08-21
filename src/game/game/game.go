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
		//render
		render.RenderBoard(board)

		// input coords
		if isWhiteToMove {
			curColor = color.White
		} else {
			curColor = color.Black
		}
		pieceCoords, curPiece := inputCoords.InputCoordsOwnPieceCanMove(curColor, board)
		coordsToMove := inputCoords.InputCoordsYourPieceToMove(curPiece, board)

		// make move
		board.MovePiece(pieceCoords, coordsToMove)

		// pass move
		isWhiteToMove = !isWhiteToMove
	}
}
