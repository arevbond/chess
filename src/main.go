package main

import (
	"chess/src/game/factory"
	"chess/src/game/game"
)

func main() {
	gameBoard := factory.BoardClassicStartPosition()
	//gameBoard := factory.BoardFromFen("3k4/8/1N3n2/8/3B4/8/8/3K4 w - - 0 1")
	game.GameLoop(gameBoard)
}
