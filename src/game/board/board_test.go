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
	knight := piece.NewKnight(color.White, knightCoords)
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
