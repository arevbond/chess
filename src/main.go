package main

import (
	"chess/src/game/board"
	"chess/src/game/game"
)

func main() {
	//gameBoard := board.BoardClassicStartPosition()
	gameBoard := board.BoardFromFen("8/3k4/8/8/2Q5/8/8/3K4 w - - 0 1")
	//ans := gameBoard.IsSquareAttackedByColor(coords.Coordinates{coords.D, coords.File(7)}, color.White)
	//fmt.Println(ans)
	game.GameLoop(gameBoard)
}
