package main

import (
	"chess/src/game/board"
	"chess/src/game/game"
)

func main() {
	//gameBoard := board.BoardClassicStartPosition()
	gameBoard := board.BoardFromFen(
		//"8/6p1/5p2/8/3B4/8/8/8 w - - 0 1",
		//"3k4/6p1/3B1p2/2N5/1r1Q1N2/8/1N3P2/3K4 w - - 0 1",
		// "8/3r4/8/8/1N1Rn3/8/8/8 w - - 0 1",
		//"8/5pp1/4p1N1/8/2p5/2n5/1p3p2/N1N5 w - - 0 1",
		"8/8/8/r3B3/4K3/8/4k3/8 w - - 0 1",
	)
	game.GameLoop(gameBoard)
}
