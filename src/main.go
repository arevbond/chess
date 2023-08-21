package main

import (
	"chess/src/game/board"
	"chess/src/game/render"
)

func main() {
	gameBoard := board.NewBoard()
	gameBoard.SetupDefaultPiecesPositions()
	render.Render(gameBoard)
}
