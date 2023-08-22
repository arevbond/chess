package coords

type CoordinatesShift struct {
	RankShift Rank
	FileShift File
}

func NewCoordinatesShift(r Rank, f File) CoordinatesShift {
	return CoordinatesShift{RankShift: r, FileShift: f}
}

func CalculateCoordinatesShift(shifts [][]int, curCoords Coordinates) map[CoordinatesShift]bool {
	coordsShifts := make(map[CoordinatesShift]bool)
	for _, shift := range shifts {
		rShift, fShift := Rank(shift[0]), File(shift[1])
		curCoordsShifts := NewCoordinatesShift(curCoords.Rank+rShift, curCoords.File+fShift)
		coordsShifts[curCoordsShifts] = true
	}
	return coordsShifts
}
