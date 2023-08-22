package coords

type File int

type Rank int

const (
	A Rank = 'A'
	B Rank = 'B'
	C Rank = 'C'
	D Rank = 'D'
	E Rank = 'E'
	F Rank = 'F'
	G Rank = 'G'
	H Rank = 'H'
)

type Coordinates struct {
	Rank Rank
	File File
}

func NewCoordinates(rank Rank, file File) Coordinates {
	return Coordinates{File: file, Rank: rank}
}

func (c Coordinates) CanShift(shift CoordinatesShift) bool {
	// проверка на то, чтобы новая координата не выходила за пределы доски
	newCoords := NewCoordinates(shift.RankShift, shift.FileShift)
	return newCoords.Rank >= 'A' && newCoords.Rank <= 'H' && newCoords.File >= 1 && newCoords.File <= 8
}

func (c Coordinates) Shift(shift CoordinatesShift) Coordinates {
	return NewCoordinates(shift.RankShift, shift.FileShift)
}
