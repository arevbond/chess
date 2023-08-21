package main

import (
	"chess/src/game/board"
	"chess/src/game/game"
)

func main() {
	gameBoard := board.NewBoard()
	gameBoard.SetupDefaultPiecesPositions()
	game.GameLoop(gameBoard)
}
