package main

import (
	"chess/src/game/factory"
	"chess/src/game/game"
)

func main() {
	//gameBoard := factory.BoardClassicStartPosition()
	gameBoard := factory.BoardFromFen("1nbqkbn1/8/7r/3r4/3R4/B6B/7R/1N1QK1N1 w - - 0 1")
	game.GameLoop(gameBoard)
}
