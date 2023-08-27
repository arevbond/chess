package main

import (
	"chess/src/game/board"
	"chess/src/game/game"
)

func main() {
	gameBoard := board.BoardClassicStartPosition()
	game.GameLoop(gameBoard)
}
