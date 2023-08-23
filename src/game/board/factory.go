package board

func BoardFromFen(fen string) *Board {
	newBoard := NewBoard()
	newBoard.SetupPositionFromFEN(fen)
	return newBoard
}

func BoardClassicStartPosition() *Board {
	newBoard := NewBoard()
	newBoard.SetupDefaultPiecesPositions()
	return newBoard
}
