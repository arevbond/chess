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
		if figure, ok := board.GetPiece(tt.to); !ok || figure != queen {
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

func TestBoard_PieceFromFenChar(t *testing.T) {
	coordinates := coords.NewCoordinates(coords.A, 8)
	var tests = []struct {
		symbol    rune
		wantPiece string
		wantColor color.Color
	}{
		{'K', "King", color.White},
		{'k', "King", color.Black},
		{'q', "Queen", color.Black},
		{'Q', "Queen", color.White},
		{'p', "Pawn", color.Black},
		{'P', "Pawn", color.White},
		{'n', "Knight", color.Black},
		{'N', "Knight", color.White},
		{'B', "Bishop", color.White},
		{'b', "Bishop", color.Black},
		{'r', "Rock", color.Black},
		{'R', "Rock", color.White},
	}

	for _, tt := range tests {
		figure := PieceFromFenChar(tt.symbol, coordinates)
		if figure.Name() != tt.wantPiece || figure.Color() != tt.wantColor || figure.Coordinates() != coordinates {
			t.Errorf("want: %s color: %d coords: %q%d - have: %s color: %d coords: %q%d", tt.wantPiece, tt.wantColor,
				coordinates.Rank, coordinates.File, figure.Name(), figure.Color(), figure.Coordinates().Rank, figure.Coordinates().File)
		}
	}
}

func TestBoard_SetupPositionFromFEN(t *testing.T) {
	startPositionFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	gameBoard := NewBoard()
	gameBoard.SetupPositionFromFEN(startPositionFen)

	var tests = []struct {
		coordinates coords.Coordinates
		figureName  string
		figureColor color.Color
	}{
		{coords.NewCoordinates(coords.A, coords.File(8)), "Rock", color.Black},
		{coords.NewCoordinates(coords.H, coords.File(8)), "Rock", color.Black},
		{coords.NewCoordinates(coords.A, coords.File(1)), "Rock", color.White},
		{coords.NewCoordinates(coords.H, coords.File(1)), "Rock", color.White},

		{coords.NewCoordinates(coords.A, coords.File(2)), "Pawn", color.White},
		{coords.NewCoordinates(coords.B, coords.File(2)), "Pawn", color.White},
		{coords.NewCoordinates(coords.C, coords.File(2)), "Pawn", color.White},
		{coords.NewCoordinates(coords.D, coords.File(2)), "Pawn", color.White},
		{coords.NewCoordinates(coords.E, coords.File(2)), "Pawn", color.White},
		{coords.NewCoordinates(coords.F, coords.File(2)), "Pawn", color.White},
		{coords.NewCoordinates(coords.G, coords.File(2)), "Pawn", color.White},
		{coords.NewCoordinates(coords.H, coords.File(2)), "Pawn", color.White},
		{coords.NewCoordinates(coords.A, coords.File(7)), "Pawn", color.Black},
		{coords.NewCoordinates(coords.B, coords.File(7)), "Pawn", color.Black},
		{coords.NewCoordinates(coords.C, coords.File(7)), "Pawn", color.Black},
		{coords.NewCoordinates(coords.D, coords.File(7)), "Pawn", color.Black},
		{coords.NewCoordinates(coords.E, coords.File(7)), "Pawn", color.Black},
		{coords.NewCoordinates(coords.F, coords.File(7)), "Pawn", color.Black},
		{coords.NewCoordinates(coords.G, coords.File(7)), "Pawn", color.Black},
		{coords.NewCoordinates(coords.H, coords.File(7)), "Pawn", color.Black},

		{coords.NewCoordinates(coords.B, coords.File(8)), "Knight", color.Black},
		{coords.NewCoordinates(coords.G, coords.File(8)), "Knight", color.Black},
		{coords.NewCoordinates(coords.B, coords.File(1)), "Knight", color.White},
		{coords.NewCoordinates(coords.G, coords.File(1)), "Knight", color.White},

		{coords.NewCoordinates(coords.C, coords.File(8)), "Bishop", color.Black},
		{coords.NewCoordinates(coords.F, coords.File(8)), "Bishop", color.Black},
		{coords.NewCoordinates(coords.C, coords.File(1)), "Bishop", color.White},
		{coords.NewCoordinates(coords.F, coords.File(1)), "Bishop", color.White},

		{coords.NewCoordinates(coords.D, coords.File(8)), "Queen", color.Black},
		{coords.NewCoordinates(coords.D, coords.File(1)), "Queen", color.White},

		{coords.NewCoordinates(coords.E, coords.File(8)), "King", color.Black},
		{coords.NewCoordinates(coords.E, coords.File(1)), "King", color.White},
	}
	for _, tt := range tests {
		figure, ok := gameBoard.GetPiece(tt.coordinates)
		if !ok {
			t.Errorf("Piece %s not on board", tt.figureName)
		}
		if tt.figureName != figure.Name() || tt.figureColor != figure.Color() {
			t.Errorf("want: %s %d - have: %s %d", tt.figureName, tt.figureColor, figure.Name(), figure.Color())
		}
	}
}

func TestBoard_KingAvailableMoves(t *testing.T) {
	var gameBoard *Board
	var availableMoves map[coords.Coordinates]bool

	// check available moves for king
	// white king
	var tests3 = []struct {
		fen           string
		kingCoords    coords.Coordinates
		lenLegalMoves int
		legalMoves    []coords.Coordinates
	}{
		{"3r4/8/8/r3B3/4K3/8/4k3/8 w - - 0 1",
			coords.NewCoordinates(coords.E, coords.File(4)),
			2,
			[]coords.Coordinates{
				{coords.F, coords.File(4)}, {coords.F, coords.File(5)},
			}},
		{"8/8/8/8/4K3/8/8/8 w - - 0 1",
			coords.NewCoordinates(coords.E, coords.File(4)),
			8,
			[]coords.Coordinates{
				{coords.F, coords.File(4)},
				{coords.F, coords.File(5)},
				{coords.F, coords.File(3)},
				{coords.E, coords.File(3)},
				{coords.D, coords.File(3)},
				{coords.D, coords.File(4)},
				{coords.D, coords.File(5)},
				{coords.E, coords.File(5)},
			}},
	}
	for _, tt := range tests3 {
		gameBoard = BoardFromFen(tt.fen)
		king, ok := gameBoard.GetPiece(tt.kingCoords)
		if !ok {
			t.Errorf("invalid king coords: %q%d", tt.kingCoords.Rank, tt.kingCoords.File)
		}
		availableMoves = gameBoard.AvailableMoves(king)
		if tt.lenLegalMoves != len(availableMoves) {
			t.Errorf("want len legal moves: %d - have len legal moves: %d", tt.lenLegalMoves, len(availableMoves))
		}
		for _, coordinates := range tt.legalMoves {
			if _, ok3 := availableMoves[coordinates]; !ok3 {
				t.Errorf("move piece %q from %q%d to %q%d should be legal\nCurrent fen: %s", king.Name(), king.Coordinates().Rank,
					king.Coordinates().File, coordinates.Rank, coordinates.File, tt.fen)
			}
		}
	}
	// black king
	var tests = []struct {
		fen           string
		kingCoords    coords.Coordinates
		lenLegalMoves int
		legalMoves    []coords.Coordinates
	}{
		{"3R4/8/8/R3b3/4k3/8/4K3/8 w - - 0 1",
			coords.NewCoordinates(coords.E, coords.File(4)),
			2,
			[]coords.Coordinates{
				{coords.F, coords.File(4)}, {coords.F, coords.File(5)},
			}},
		{"8/8/8/8/4k3/8/8/8 w - - 0 1",
			coords.NewCoordinates(coords.E, coords.File(4)),
			8,
			[]coords.Coordinates{
				{coords.F, coords.File(4)},
				{coords.F, coords.File(5)},
				{coords.F, coords.File(3)},
				{coords.E, coords.File(3)},
				{coords.D, coords.File(3)},
				{coords.D, coords.File(4)},
				{coords.D, coords.File(5)},
				{coords.E, coords.File(5)},
			}},
	}
	for _, tt := range tests {
		gameBoard = BoardFromFen(tt.fen)
		king, ok := gameBoard.GetPiece(tt.kingCoords)
		if !ok {
			t.Errorf("invalid king coords: %q%d", tt.kingCoords.Rank, tt.kingCoords.File)
		}
		availableMoves = gameBoard.AvailableMoves(king)
		if tt.lenLegalMoves != len(availableMoves) {
			t.Errorf("want len legal moves: %d - have len legal moves: %d", tt.lenLegalMoves, len(availableMoves))
		}
		for _, coordinates := range tt.legalMoves {
			if _, ok3 := availableMoves[coordinates]; !ok3 {
				t.Errorf("move piece %q from %q%d to %q%d should be legal\nCurrent fen: %s", king.Name(), king.Coordinates().Rank,
					king.Coordinates().File, coordinates.Rank, coordinates.File, tt.fen)
			}
		}
	}
}

func TestBoard_PawnAvailableMoves(t *testing.T) {
	var fen string
	var gameBoard *Board
	var availableMoves map[coords.Coordinates]bool

	// check available moves for pawn
	// white pawn
	fen = "8/pN1N4/1P1P1n1n/6P1/1n6/n4P2/PP1P4/8 w - - 0 1"
	gameBoard = BoardFromFen(fen)
	var tests = []struct {
		pieceCoords   coords.Coordinates
		lenLegalMoves int
		legalMoves    []coords.Coordinates
	}{
		{coords.NewCoordinates(coords.A, coords.File(2)), 0, nil},
		{coords.NewCoordinates(coords.B, coords.File(2)), 2, []coords.Coordinates{
			coords.NewCoordinates(coords.A, coords.File(3)),
			coords.NewCoordinates(coords.B, coords.File(3)),
		}},
		{coords.NewCoordinates(coords.B, coords.File(6)), 1, []coords.Coordinates{
			coords.NewCoordinates(coords.A, coords.File(7)),
		}},
		{coords.NewCoordinates(coords.D, coords.File(2)), 2, []coords.Coordinates{
			coords.NewCoordinates(coords.D, coords.File(3)),
			coords.NewCoordinates(coords.D, coords.File(4)),
		}},
		{coords.NewCoordinates(coords.F, coords.File(3)), 1, []coords.Coordinates{
			coords.NewCoordinates(coords.F, coords.File(4)),
		}},
		{coords.NewCoordinates(coords.G, coords.File(5)), 3, []coords.Coordinates{
			coords.NewCoordinates(coords.G, coords.File(6)),
			coords.NewCoordinates(coords.F, coords.File(6)),
			coords.NewCoordinates(coords.H, coords.File(6)),
		}},
	}
	for _, tt := range tests {
		pawn, ok := gameBoard.GetPiece(tt.pieceCoords)
		if !ok {
			t.Errorf("not piece on %q%d", tt.pieceCoords.Rank, tt.pieceCoords.File)
		}
		availableMoves = gameBoard.AvailableMoves(pawn)
		if len(availableMoves) != tt.lenLegalMoves {
			t.Errorf("want len legal moves: %d - have len legal moves: %d", tt.lenLegalMoves, len(availableMoves))
		}
		for _, coordinates := range tt.legalMoves {
			if _, ok2 := availableMoves[coordinates]; !ok2 {
				t.Errorf("move piece %q from %q%d to %q%d should be legal\nCurrent fen: %s", pawn.Name(), pawn.Coordinates().Rank,
					pawn.Coordinates().File, coordinates.Rank, coordinates.File, fen)
			}
		}
	}
	// black pawn
	fen = "8/2p1p2p/2pP1Pp1/1p6/1R6/3p4/2N1Q3/8 w - - 0 1"
	gameBoard = BoardFromFen(fen)
	var tests2 = []struct {
		pieceCoords   coords.Coordinates
		lenLegalMoves int
		legalMoves    []coords.Coordinates
	}{
		{coords.NewCoordinates(coords.B, coords.File(5)), 0, nil},
		{coords.NewCoordinates(coords.C, coords.File(7)), 1, []coords.Coordinates{
			coords.NewCoordinates(coords.D, coords.File(6)),
		}},
		{coords.NewCoordinates(coords.C, coords.File(6)), 1, []coords.Coordinates{
			coords.NewCoordinates(coords.C, coords.File(5)),
		}},
		{coords.NewCoordinates(coords.D, coords.File(3)), 3, []coords.Coordinates{
			coords.NewCoordinates(coords.D, coords.File(2)),
			coords.NewCoordinates(coords.C, coords.File(2)),
			coords.NewCoordinates(coords.E, coords.File(2)),
		}},
		{coords.NewCoordinates(coords.E, coords.File(7)), 4, []coords.Coordinates{
			coords.NewCoordinates(coords.E, coords.File(6)),
			coords.NewCoordinates(coords.E, coords.File(5)),
			coords.NewCoordinates(coords.D, coords.File(6)),
			coords.NewCoordinates(coords.F, coords.File(6)),
		}},
		{coords.NewCoordinates(coords.G, coords.File(6)), 1, []coords.Coordinates{
			coords.NewCoordinates(coords.G, coords.File(5)),
		}},
		{coords.NewCoordinates(coords.H, coords.File(7)), 2, []coords.Coordinates{
			coords.NewCoordinates(coords.H, coords.File(6)),
			coords.NewCoordinates(coords.H, coords.File(5)),
		}},
	}
	for _, tt := range tests2 {
		pawn, ok := gameBoard.GetPiece(tt.pieceCoords)
		if !ok {
			t.Errorf("not piece on %q%d", tt.pieceCoords.Rank, tt.pieceCoords.File)
		}
		availableMoves = gameBoard.AvailableMoves(pawn)
		if len(availableMoves) != tt.lenLegalMoves {
			t.Errorf("want len legal moves: %d - have len legal moves: %d", tt.lenLegalMoves, len(availableMoves))
		}
		for _, coordinates := range tt.legalMoves {
			if _, ok2 := availableMoves[coordinates]; !ok2 {
				t.Errorf("move piece %q from %q%d to %q%d should be legal\nCurrent fen: %s", pawn.Name(), pawn.Coordinates().Rank,
					pawn.Coordinates().File, coordinates.Rank, coordinates.File, fen)
			}
		}
	}
}

func TestBoard_QueenAvailableMoves(t *testing.T) {
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
	fen = "8/3P4/1P3P2/8/1P1Q2P1/8/1P3P2/3P4 w - - 0 1"
	gameBoard = BoardFromFen(fen) // queen on d4
	queen, _ = gameBoard.GetPiece(coords.NewCoordinates(coords.D, coords.File(4)))
	availableMoves = gameBoard.AvailableMoves(queen)
	legalMoves = []coords.Coordinates{
		{coords.C, coords.File(3)},
		{coords.C, coords.File(4)},
		{coords.C, coords.File(5)},
		{coords.D, coords.File(6)},
		{coords.D, coords.File(5)},
		{coords.D, coords.File(3)},
		{coords.D, coords.File(2)},
		{coords.E, coords.File(3)},
		{coords.E, coords.File(4)},
		{coords.E, coords.File(5)},
		{coords.F, coords.File(4)},
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
	fen = "8/3p4/1p3p2/8/1p1q2p1/8/1p3p2/3p4 w - - 0 1"
	gameBoard = BoardFromFen(fen) // queen on d4
	queen, _ = gameBoard.GetPiece(coords.NewCoordinates(coords.D, coords.File(4)))
	availableMoves = gameBoard.AvailableMoves(queen)
	legalMoves = []coords.Coordinates{
		{coords.C, coords.File(3)},
		{coords.C, coords.File(4)},
		{coords.C, coords.File(5)},
		{coords.D, coords.File(6)},
		{coords.D, coords.File(5)},
		{coords.D, coords.File(3)},
		{coords.D, coords.File(2)},
		{coords.E, coords.File(3)},
		{coords.E, coords.File(4)},
		{coords.E, coords.File(5)},
		{coords.F, coords.File(4)},
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
}

func TestBoard_KnightAvailableMoves(t *testing.T) {
	var fen string
	var gameBoard *Board
	var availableMoves map[coords.Coordinates]bool
	var legalMoves []coords.Coordinates

	// check available moves for knight
	// white knight
	fen = "8/8/4n3/1NP1Pp2/3N4/2P1P3/2Q5/8 w - - 0 1"
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
	fen = "8/8/4p3/1Nn2p2/2nn4/2p1P3/2Q1N3/8 w - - 0 1"
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
}

func TestBoard_BishopAvailableMoves(t *testing.T) {
	var fen string
	var gameBoard *Board
	var availableMoves map[coords.Coordinates]bool
	var legalMoves []coords.Coordinates

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
}

func TestBoard_RockAvailableMoves(t *testing.T) {
	var fen string
	var gameBoard *Board
	var availableMoves map[coords.Coordinates]bool
	var legalMoves []coords.Coordinates

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
