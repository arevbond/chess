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

func Copy(board *Board) *Board {
	newBoard := BoardFromFen(board.startFen)
	for _, m := range board.Moves {
		newBoard.MovePiece(m.From, m.To)
	}
	return newBoard
}
