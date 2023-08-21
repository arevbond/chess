package inputCoords

import (
	"bufio"
	"chess/src/game/board"
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func InputCoordinates() coords.Coordinates {
	scanner := bufio.NewScanner(os.Stdin)
	var r coords.Rank
	var f coords.File
	for {
		fmt.Println("Enter coordinates (ex. g2):")
		scanner.Scan()
		in := scanner.Text()
		in = strings.ReplaceAll(in, " ", "")
		if len(in) != 2 {
			fmt.Println("Invalid coordinates format")
			continue
		}
		runes := []rune(in)
		rankChar := unicode.ToUpper(runes[0])
		fileChar, err := strconv.Atoi(string(runes[1]))
		if err != nil {
			fmt.Println("Invalid coordinates format")
			continue
		}
		if unicode.IsLetter(rankChar) {
			if rankChar < 'A' || rankChar > 'H' || fileChar < 1 || fileChar > 8 {
				fmt.Println("Invalid coordinates value")
				continue
			}
		}
		r = coords.Rank(rankChar)
		f = coords.File(fileChar)

		fmt.Printf("Your choose is %q%d. Correct? (y/n)\n", r, f)
		scanner.Scan()
		ans := scanner.Text()
		if strings.HasPrefix(strings.ToLower(ans), "y") {
			break
		} else {
			continue
		}

	}
	return coords.NewCoordinates(r, f)
}

func InputPieceByColor(color color.Color, board *board.Board) (coords.Coordinates, piece.Piece) {
	for {
		fmt.Println("Please choose your piece")
		curCoords := InputCoordinates()
		curPiece, ok := board.GetPiece(curCoords)
		if !ok {
			fmt.Println("Piece doesn't choose")
			continue
		}
		if curPiece.Color() != color {
			continue
		}
		return curCoords, curPiece
	}
}

func InputCoordsOwnPieceCanMove(color color.Color, board *board.Board) (coords.Coordinates, piece.Piece) {
	var curCoords coords.Coordinates
	var curPiece piece.Piece
	for {
		curCoords, curPiece = InputPieceByColor(color, board)
		if len(board.AvailableMoves(curPiece)) > 0 {
			break
		} else {
			fmt.Println("Your piece can't moves")
		}
	}
	return curCoords, curPiece
}

func InputCoordsYourPieceToMove(piece piece.Piece, board *board.Board) coords.Coordinates {
	var coordsToMove coords.Coordinates
	avaliableMoves := board.AvailableMoves(piece)
	for {
		fmt.Println("Where you want to move?")
		coordsToMove = InputCoordinates()
		if _, ok := avaliableMoves[coordsToMove]; !ok {
			fmt.Println("Impossible move")
			continue
		} else {
			break
		}
	}
	return coordsToMove
}
