package main

import (
	"chess/src/game/factory"
	"chess/src/game/game"
)

func main() {
	//gameBoard := factory.BoardClassicStartPosition()
	//gameBoard := factory.BoardFromFen("8/6p1/5p2/8/3B4/8/8/8 w - - 0 1")
	gameBoard := factory.BoardFromFen("3k4/6p1/3B1p2/2N5/1P1Q1N2/8/1N3P2/3K4 w - - 0 1")
	game.GameLoop(gameBoard)
}
