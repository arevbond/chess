package factory

import "chess/src/game/board"

func BoardFromFen(fen string) *board.Board {
	newBoard := board.NewBoard()
	newBoard.SetupPositionFromFEN(fen)
	return newBoard
}

func BoardClassicStartPosition() *board.Board {
	newBoard := board.NewBoard()
	newBoard.SetupDefaultPiecesPositions()
	return newBoard
}
