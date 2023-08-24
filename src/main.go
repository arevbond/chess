package main

import (
	"chess/src/game/board"
	"chess/src/game/game"
)

func main() {
	//gameBoard := board.BoardClassicStartPosition()
	gameBoard := board.BoardFromFen("8/2n5/8/4Q3/8/8/R7/8 w - - 0 1")
	game.GameLoop(gameBoard)
}
