package main

import (
	"chess/src/game/board"
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/game"
	"fmt"
)

func main() {
	//gameBoard := board.BoardClassicStartPosition()
	gameBoard := board.BoardFromFen("8/3k4/8/8/3Q4/8/8/8 w - - 0 1")
	ans := gameBoard.IsSquareAttackedByColor(coords.Coordinates{coords.D, coords.File(7)}, color.White)
	fmt.Println(ans)
	game.GameLoop(gameBoard)
}
