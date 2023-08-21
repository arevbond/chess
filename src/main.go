package main

import (
	"chess/src/game/board"
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
	"fmt"
)

func main() {
	gameBoard := board.NewBoard()
	gameBoard.SetupDefaultPiecesPositions()
	//render.Render(gameBoard)
	knight := piece.NewKnight(color.White, coords.NewCoordinates('E', 4))
	//knight := piece.NewKnight(color.White, coords.NewCoordinates('A', 4))
	//shifts := knight.Shifts()
	//for c, _ := range shifts {
	//	fmt.Printf("%q %d \n", c.RankShift, c.FileShift)
	//}
	avMoves := gameBoard.AvailableMoves(knight)
	//fmt.Println(avMoves)
	for c, _ := range avMoves {
		fmt.Printf("%q %d \n", c.Rank, c.File)
	}
}
