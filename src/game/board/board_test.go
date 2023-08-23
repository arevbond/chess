package board

import (
	"chess/src/game/color"
	"chess/src/game/coords"
	"chess/src/game/piece"
	"testing"
)

func TestBoard_SetPiece(t *testing.T) {
	curBoard := NewBoard()
	knightCoords := coords.Coordinates{File: 1, Rank: 'B'}
	knight := PieceFromFenChar('N', knightCoords)
	curBoard.SetPiece(knightCoords, knight)
	if k, ok := curBoard.Pieces[knightCoords]; !ok || k != knight {
		t.Errorf("func Board.SetPiece doesn't set knight")
	}

	bishopCoords := coords.Coordinates{File: 4, Rank: 'G'}
	bishop := piece.NewBishop(color.Black, bishopCoords)
	curBoard.SetPiece(bishopCoords, bishop)
	if b, ok := curBoard.Pieces[bishopCoords]; !ok || b != bishop {
		t.Errorf("func Board.SetPiece doesn't set bishop")
	}

	pawnCoords := coords.Coordinates{File: 2, Rank: 'A'}
	pawn := piece.NewPawn(color.Black, pawnCoords)
	curBoard.SetPiece(pawnCoords, pawn)
	if p, ok := curBoard.Pieces[pawnCoords]; !ok || p != pawn {
		t.Errorf("func Board.SetPiece doesn't set pawn")
	}

	rockCoords := coords.Coordinates{File: 6, Rank: 'C'}
	rock := piece.NewRock(color.Black, rockCoords)
	curBoard.SetPiece(rockCoords, rock)
	if r, ok := curBoard.Pieces[rockCoords]; !ok || r != rock {
		t.Errorf("func Board.SetPiece doesn't set rock")
	}

	queenCoords := coords.Coordinates{File: 3, Rank: 'H'}
	queen := piece.NewQueen(color.Black, queenCoords)
	curBoard.SetPiece(queenCoords, queen)
	if q, ok := curBoard.Pieces[queenCoords]; !ok || q != queen {
		t.Errorf("func Board.SetPiece doesn't set queen")
	}

	kingCoords := coords.Coordinates{File: 8, Rank: 'F'}
	king := piece.NewKing(color.Black, kingCoords)
	curBoard.SetPiece(kingCoords, king)
	if k, ok := curBoard.Pieces[kingCoords]; !ok || k != king {
		t.Errorf("func Board.SetPiece doesn't set queen")
	}

	// установка фигуры на место где уже стоит фигура
	newBishop := piece.NewBishop(color.White, kingCoords)
	curBoard.SetPiece(kingCoords, newBishop)
	if b, ok := curBoard.Pieces[kingCoords]; !ok || b != newBishop {
		t.Errorf("func Board.SetPiece doesn't update new piece")
	}
}

func TestBoard_RemovePiece(t *testing.T) {
	curBoard := NewBoard()
	knightCoords := coords.Coordinates{File: 1, Rank: 'B'}
	knight := piece.NewKnight(color.White, knightCoords)
	curBoard.SetPiece(knightCoords, knight)
	curBoard.RemovePiece(knightCoords)
	if len(curBoard.Pieces) != 0 {
		t.Errorf("Board.RemovePiece doesn't remove the piece")
	}
}

func TestBoard_MovePiece(t *testing.T) {
	board := NewBoard()

	var tests = []struct {
		from, to coords.Coordinates
	}{
		{coords.Coordinates{Rank: coords.A, File: coords.File(1)}, coords.Coordinates{Rank: coords.A, File: coords.File(2)}},
		{coords.Coordinates{Rank: coords.B, File: coords.File(1)}, coords.Coordinates{Rank: coords.G, File: coords.File(2)}},
		{coords.Coordinates{Rank: coords.C, File: coords.File(1)}, coords.Coordinates{Rank: coords.H, File: coords.File(2)}},
	}
	for _, tt := range tests {
		queen := piece.NewQueen(color.White, tt.from)
		board.SetPiece(tt.from, queen)
		board.MovePiece(tt.from, tt.to)
		if curPiece, ok := board.GetPiece(tt.to); !ok || curPiece != queen {
			t.Errorf("piece %s doen't move from %q%d to %q%d", queen.Name(), tt.from.Rank, tt.from.File,
				tt.to.Rank, tt.to.File)
		}
	}
}

func TestBoard_IsSquareDark(t *testing.T) {
	board := NewBoard()
	var tests = []struct {
		input coords.Coordinates
		want  bool
	}{
		{coords.Coordinates{Rank: coords.A, File: coords.File(1)}, true},
		{coords.Coordinates{Rank: coords.B, File: coords.File(1)}, false},
		{coords.Coordinates{Rank: coords.C, File: coords.File(1)}, true},
		{coords.Coordinates{Rank: coords.D, File: coords.File(1)}, false},
		{coords.Coordinates{Rank: coords.E, File: coords.File(1)}, true},
		{coords.Coordinates{Rank: coords.F, File: coords.File(1)}, false},
		{coords.Coordinates{Rank: coords.G, File: coords.File(1)}, true},
		{coords.Coordinates{Rank: coords.H, File: coords.File(1)}, false},

		{coords.Coordinates{Rank: coords.A, File: coords.File(2)}, false},
		{coords.Coordinates{Rank: coords.B, File: coords.File(2)}, true},
		{coords.Coordinates{Rank: coords.C, File: coords.File(2)}, false},
		{coords.Coordinates{Rank: coords.D, File: coords.File(2)}, true},
		{coords.Coordinates{Rank: coords.E, File: coords.File(2)}, false},
		{coords.Coordinates{Rank: coords.F, File: coords.File(2)}, true},
		{coords.Coordinates{Rank: coords.G, File: coords.File(2)}, false},
		{coords.Coordinates{Rank: coords.H, File: coords.File(2)}, true},

		{coords.Coordinates{Rank: coords.A, File: coords.File(3)}, true},
		{coords.Coordinates{Rank: coords.B, File: coords.File(3)}, false},
		{coords.Coordinates{Rank: coords.C, File: coords.File(3)}, true},
		{coords.Coordinates{Rank: coords.D, File: coords.File(3)}, false},
		{coords.Coordinates{Rank: coords.E, File: coords.File(3)}, true},
		{coords.Coordinates{Rank: coords.F, File: coords.File(3)}, false},
		{coords.Coordinates{Rank: coords.G, File: coords.File(3)}, true},
		{coords.Coordinates{Rank: coords.H, File: coords.File(3)}, false},

		{coords.Coordinates{Rank: coords.A, File: coords.File(4)}, false},
		{coords.Coordinates{Rank: coords.B, File: coords.File(4)}, true},
		{coords.Coordinates{Rank: coords.C, File: coords.File(4)}, false},
		{coords.Coordinates{Rank: coords.D, File: coords.File(4)}, true},
		{coords.Coordinates{Rank: coords.E, File: coords.File(4)}, false},
		{coords.Coordinates{Rank: coords.F, File: coords.File(4)}, true},
		{coords.Coordinates{Rank: coords.G, File: coords.File(4)}, false},
		{coords.Coordinates{Rank: coords.H, File: coords.File(4)}, true},

		{coords.Coordinates{Rank: coords.A, File: coords.File(5)}, true},
		{coords.Coordinates{Rank: coords.B, File: coords.File(5)}, false},
		{coords.Coordinates{Rank: coords.C, File: coords.File(5)}, true},
		{coords.Coordinates{Rank: coords.D, File: coords.File(5)}, false},
		{coords.Coordinates{Rank: coords.E, File: coords.File(5)}, true},
		{coords.Coordinates{Rank: coords.F, File: coords.File(5)}, false},
		{coords.Coordinates{Rank: coords.G, File: coords.File(5)}, true},
		{coords.Coordinates{Rank: coords.H, File: coords.File(5)}, false},

		{coords.Coordinates{Rank: coords.A, File: coords.File(6)}, false},
		{coords.Coordinates{Rank: coords.B, File: coords.File(6)}, true},
		{coords.Coordinates{Rank: coords.C, File: coords.File(6)}, false},
		{coords.Coordinates{Rank: coords.D, File: coords.File(6)}, true},
		{coords.Coordinates{Rank: coords.E, File: coords.File(6)}, false},
		{coords.Coordinates{Rank: coords.F, File: coords.File(6)}, true},
		{coords.Coordinates{Rank: coords.G, File: coords.File(6)}, false},
		{coords.Coordinates{Rank: coords.H, File: coords.File(6)}, true},

		{coords.Coordinates{Rank: coords.A, File: coords.File(7)}, true},
		{coords.Coordinates{Rank: coords.B, File: coords.File(7)}, false},
		{coords.Coordinates{Rank: coords.C, File: coords.File(7)}, true},
		{coords.Coordinates{Rank: coords.D, File: coords.File(7)}, false},
		{coords.Coordinates{Rank: coords.E, File: coords.File(7)}, true},
		{coords.Coordinates{Rank: coords.F, File: coords.File(7)}, false},
		{coords.Coordinates{Rank: coords.G, File: coords.File(7)}, true},
		{coords.Coordinates{Rank: coords.H, File: coords.File(7)}, false},

		{coords.Coordinates{Rank: coords.A, File: coords.File(8)}, false},
		{coords.Coordinates{Rank: coords.B, File: coords.File(8)}, true},
		{coords.Coordinates{Rank: coords.C, File: coords.File(8)}, false},
		{coords.Coordinates{Rank: coords.D, File: coords.File(8)}, true},
		{coords.Coordinates{Rank: coords.E, File: coords.File(8)}, false},
		{coords.Coordinates{Rank: coords.F, File: coords.File(8)}, true},
		{coords.Coordinates{Rank: coords.G, File: coords.File(8)}, false},
		{coords.Coordinates{Rank: coords.H, File: coords.File(8)}, true},
	}
	for _, tt := range tests {
		ans := board.IsSquareDark(tt.input)
		if ans != tt.want {
			t.Errorf("got %t, want %t", ans, tt.want)
		}
	}
}

func TestBoard_IsSquareEmpty(t *testing.T) {
	board := NewBoard()
	piecesCoords := []coords.Coordinates{
		{'A', 1},
		{'B', 1},
		{'C', 1},
		{'D', 1},
		{'E', 8},
		{'F', 8},
		{'G', 8},
		{'H', 8},
	}
	for i, curCoords := range piecesCoords {
		var curColor color.Color = color.Black
		if i%2 == 0 {
			curColor = color.White
		}
		board.SetPiece(curCoords, piece.NewPawn(curColor, curCoords))
	}

	var tests = []struct {
		input coords.Coordinates
		want  bool
	}{
		{coords.Coordinates{Rank: coords.A, File: coords.File(1)}, false},
		{coords.Coordinates{Rank: coords.B, File: coords.File(1)}, false},
		{coords.Coordinates{Rank: coords.C, File: coords.File(1)}, false},
		{coords.Coordinates{Rank: coords.D, File: coords.File(1)}, false},
		{coords.Coordinates{Rank: coords.E, File: coords.File(8)}, false},
		{coords.Coordinates{Rank: coords.F, File: coords.File(8)}, false},
		{coords.Coordinates{Rank: coords.G, File: coords.File(8)}, false},
		{coords.Coordinates{Rank: coords.H, File: coords.File(8)}, false},

		{coords.Coordinates{Rank: coords.A, File: coords.File(2)}, true},
		{coords.Coordinates{Rank: coords.B, File: coords.File(3)}, true},
		{coords.Coordinates{Rank: coords.C, File: coords.File(4)}, true},
		{coords.Coordinates{Rank: coords.D, File: coords.File(5)}, true},
		{coords.Coordinates{Rank: coords.E, File: coords.File(6)}, true},
		{coords.Coordinates{Rank: coords.F, File: coords.File(7)}, true},
		{coords.Coordinates{Rank: coords.G, File: coords.File(2)}, true},
		{coords.Coordinates{Rank: coords.H, File: coords.File(3)}, true},
	}
	for _, tt := range tests {
		ans := board.IsSquareEmpty(tt.input)
		if ans != tt.want {
			t.Errorf("got %t, want %t", ans, tt.want)
		}
	}
}

func TestBoard_IsSquareAvailableForMoveSimple(t *testing.T) {
	board := NewBoard()

	whiteKingCoords := coords.NewCoordinates(coords.A, coords.File(1))
	board.SetPiece(whiteKingCoords, piece.NewKing(color.White, whiteKingCoords))

	blackKingCoords := coords.NewCoordinates(coords.B, coords.File(1))
	board.SetPiece(blackKingCoords, piece.NewKing(color.Black, blackKingCoords))

	whitePawnCoords := coords.NewCoordinates(coords.C, coords.File(1))
	board.SetPiece(whitePawnCoords, piece.NewPawn(color.White, whitePawnCoords))

	blackPawnCoords := coords.NewCoordinates(coords.D, coords.File(1))
	board.SetPiece(blackPawnCoords, piece.NewPawn(color.Black, blackPawnCoords))

	var tests = []struct {
		name        string
		inputCoords coords.Coordinates
		inputPiece  piece.Piece
		want        bool
	}{
		{"Black piece try move on white King",
			whiteKingCoords,
			piece.NewQueen(color.Black, coords.NewCoordinates(coords.A, coords.File(2))),
			false},
		{"White piece try move on black King",
			blackKingCoords,
			piece.NewQueen(color.White, coords.NewCoordinates(coords.A, coords.File(2))),
			false},
		{"Black piece try move on white piece",
			whitePawnCoords,
			piece.NewQueen(color.Black, coords.NewCoordinates(coords.A, coords.File(2))),
			true},
		{"White piece try move on black piece",
			blackPawnCoords,
			piece.NewQueen(color.White, coords.NewCoordinates(coords.A, coords.File(2))),
			true},
		{"White piece try move on empty square",
			coords.NewCoordinates(coords.B, coords.File(5)),
			piece.NewQueen(color.White, coords.NewCoordinates(coords.A, coords.File(2))),
			true},
		{"Black piece try move on empty square",
			coords.NewCoordinates(coords.B, coords.File(6)),
			piece.NewQueen(color.Black, coords.NewCoordinates(coords.A, coords.File(2))),
			true},
	}

	for _, tt := range tests {
		ans := board.IsSquareAvailableForMoveSimple(tt.inputCoords, tt.inputPiece)
		if ans != tt.want {
			t.Errorf("%s - got %t, want %t", tt.name, ans, tt.want)
		}
	}
}

//func TestBoard_SetupPositionFromFEN(t *testing.T) {
//
//}

func TestBoard_AvailableMoves(t *testing.T) {
	// TODO: add for king and pawn
	var fen string
	var gameBoard *Board
	var availableMoves map[coords.Coordinates]bool
	var legalMoves []coords.Coordinates

	// check available moves for queen
	// white queen
	fen = "3k4/6p1/3B1p2/2N5/1r1Q1N2/8/1N3P2/3K4 w - - 0 1"
	gameBoard = BoardFromFen(fen) // queen on d4
	queen, _ := gameBoard.GetPiece(coords.NewCoordinates(coords.D, coords.File(4)))
	availableMoves = gameBoard.AvailableMoves(queen)
	legalMoves = []coords.Coordinates{
		{coords.B, coords.File(4)},
		{coords.C, coords.File(4)},
		{coords.C, coords.File(3)},
		{coords.E, coords.File(4)},
		{coords.D, coords.File(5)},
		{coords.D, coords.File(3)},
		{coords.D, coords.File(2)},
		{coords.E, coords.File(5)},
		{coords.E, coords.File(3)},
		{coords.F, coords.File(6)},
	}
	if len(legalMoves) != len(availableMoves) {
		t.Errorf("len available moves don't match len legal moves")
	}
	for _, coordinates := range legalMoves {
		if _, ok := availableMoves[coordinates]; !ok {
			t.Errorf("move piece %q to %q%d should be legal\nCurrent fen: %s", queen.Name(), coordinates.Rank,
				coordinates.File, fen)

		}
	}
	// black queen
	fen = "3p4/8/5n2/2p5/2pq2N1/2R5/5P2/8 w - - 0 1"
	gameBoard = BoardFromFen(fen) // queen on d4
	queen, _ = gameBoard.GetPiece(coords.NewCoordinates(coords.D, coords.File(4)))
	availableMoves = gameBoard.AvailableMoves(queen)
	legalMoves = []coords.Coordinates{
		{coords.C, coords.File(3)},
		{coords.D, coords.File(7)},
		{coords.D, coords.File(6)},
		{coords.D, coords.File(5)},
		{coords.D, coords.File(3)},
		{coords.D, coords.File(2)},
		{coords.D, coords.File(1)},
		{coords.E, coords.File(3)},
		{coords.E, coords.File(4)},
		{coords.E, coords.File(5)},
		{coords.F, coords.File(2)},
		{coords.F, coords.File(4)},
		{coords.G, coords.File(4)},
	}
	if len(legalMoves) != len(availableMoves) {
		t.Errorf("len available moves don't match len legal moves")
	}
	for _, coordinates := range legalMoves {
		if _, ok := availableMoves[coordinates]; !ok {
			t.Errorf("move piece %q to %q%d should be legal\nCurrent fen: %s", queen.Name(), coordinates.Rank,
				coordinates.File, fen)

		}
	}

	// check available moves for knight
	// white knight
	fen = "8/8/4n3/1N3p2/3N4/8/2Q5/8 w - - 0 1"
	gameBoard = BoardFromFen(fen) // knight on d4
	knight, _ := gameBoard.GetPiece(coords.NewCoordinates(coords.D, coords.File(4)))
	availableMoves = gameBoard.AvailableMoves(knight)
	legalMoves = []coords.Coordinates{
		{coords.B, coords.File(3)},
		{coords.C, coords.File(6)},
		{coords.E, coords.File(6)},
		{coords.E, coords.File(2)},
		{coords.F, coords.File(3)},
		{coords.F, coords.File(5)},
	}
	if len(legalMoves) != len(availableMoves) {
		t.Errorf("len available moves don't match len legal moves")
	}
	for _, coordinates := range legalMoves {
		if _, ok := availableMoves[coordinates]; !ok {
			t.Errorf("move piece %q to %q%d should be legal\nCurrent fen: %s", knight.Name(), coordinates.Rank,
				coordinates.File, fen)

		}
	}
	// black knight
	fen = "8/8/4p3/1N3p2/3n4/8/2Q1N3/8 w - - 0 1"
	gameBoard = BoardFromFen(fen) // knight on d4
	knight, _ = gameBoard.GetPiece(coords.NewCoordinates(coords.D, coords.File(4)))
	availableMoves = gameBoard.AvailableMoves(knight)
	legalMoves = []coords.Coordinates{
		{coords.B, coords.File(3)},
		{coords.B, coords.File(5)},
		{coords.C, coords.File(2)},
		{coords.C, coords.File(6)},
		{coords.E, coords.File(2)},
		{coords.F, coords.File(3)},
	}
	if len(legalMoves) != len(availableMoves) {
		t.Errorf("len available moves don't match len legal moves")
	}
	for _, coordinates := range legalMoves {
		if _, ok := availableMoves[coordinates]; !ok {
			t.Errorf("move piece %q to %q%d should be legal\nCurrent fen: %s", knight.Name(), coordinates.Rank,
				coordinates.File, fen)

		}
	}
	// check available moves for bishop
	// white bishop
	fen = "8/8/5n2/8/3B4/2N5/5N2/8 w - - 0 1"
	gameBoard = BoardFromFen(fen) // bishop on d4
	bishop, _ := gameBoard.GetPiece(coords.NewCoordinates(coords.D, coords.File(4)))
	availableMoves = gameBoard.AvailableMoves(bishop)
	legalMoves = []coords.Coordinates{
		{coords.A, coords.File(7)},
		{coords.B, coords.File(6)},
		{coords.C, coords.File(5)},
		{coords.E, coords.File(3)},
		{coords.E, coords.File(5)},
		{coords.F, coords.File(6)},
	}
	if len(legalMoves) != len(availableMoves) {
		t.Errorf("len available moves don't match len legal moves")
	}
	for _, coordinates := range legalMoves {
		if _, ok := availableMoves[coordinates]; !ok {
			t.Errorf("move piece %q to %q%d should be legal\nCurrent fen: %s", bishop.Name(), coordinates.Rank,
				coordinates.File, fen)

		}
	}
	// black bishop
	fen = "8/8/5n2/8/3b4/2N5/5N2/8 w - - 0 1"
	gameBoard = BoardFromFen(fen) // bishop on d4
	bishop, _ = gameBoard.GetPiece(coords.NewCoordinates(coords.D, coords.File(4)))
	availableMoves = gameBoard.AvailableMoves(bishop)
	legalMoves = []coords.Coordinates{
		{coords.A, coords.File(7)},
		{coords.B, coords.File(6)},
		{coords.C, coords.File(5)},
		{coords.C, coords.File(3)},
		{coords.E, coords.File(5)},
		{coords.E, coords.File(3)},
		{coords.F, coords.File(2)},
	}
	if len(legalMoves) != len(availableMoves) {
		t.Errorf("len available moves don't match len legal moves")
	}
	for _, coordinates := range legalMoves {
		if _, ok := availableMoves[coordinates]; !ok {
			t.Errorf("move piece %q to %q%d should be legal\nCurrent fen: %s", bishop.Name(), coordinates.Rank,
				coordinates.File, fen)

		}
	}

	// check available moves for rock
	// white rock
	fen = "8/3r4/8/8/1N1Rn3/8/8/8 w - - 0 1"
	gameBoard = BoardFromFen(fen) // rock on d4
	rock, _ := gameBoard.GetPiece(coords.NewCoordinates(coords.D, coords.File(4)))
	availableMoves = gameBoard.AvailableMoves(rock)
	legalMoves = []coords.Coordinates{
		{coords.C, coords.File(4)},
		{coords.D, coords.File(5)},
		{coords.D, coords.File(6)},
		{coords.D, coords.File(7)},
		{coords.D, coords.File(3)},
		{coords.D, coords.File(2)},
		{coords.D, coords.File(1)},
		{coords.E, coords.File(4)},
	}
	if len(legalMoves) != len(availableMoves) {
		t.Errorf("len available moves don't match len legal moves")
	}
	for _, coordinates := range legalMoves {
		if _, ok := availableMoves[coordinates]; !ok {
			t.Errorf("move piece %q to %q%d should be legal\nCurrent fen: %s", rock.Name(), coordinates.Rank,
				coordinates.File, fen)

		}
	} // black rock
	fen = "8/3r4/8/8/1N1r2n1/8/3R4/8 w - - 0 1"
	gameBoard = BoardFromFen(fen) // rock on d4
	rock, _ = gameBoard.GetPiece(coords.NewCoordinates(coords.D, coords.File(4)))
	availableMoves = gameBoard.AvailableMoves(rock)
	legalMoves = []coords.Coordinates{
		{coords.B, coords.File(4)},
		{coords.C, coords.File(4)},
		{coords.D, coords.File(5)},
		{coords.D, coords.File(6)},
		{coords.D, coords.File(3)},
		{coords.D, coords.File(2)},
		{coords.E, coords.File(4)},
		{coords.F, coords.File(4)},
	}
	if len(legalMoves) != len(availableMoves) {
		t.Errorf("len available moves don't match len legal moves")
	}
	for _, coordinates := range legalMoves {
		if _, ok := availableMoves[coordinates]; !ok {
			t.Errorf("move piece %q to %q%d should be legal\nCurrent fen: %s", rock.Name(), coordinates.Rank,
				coordinates.File, fen)

		}
	}
}
