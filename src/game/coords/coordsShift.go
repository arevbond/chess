package coords

type CoordinatesShift struct {
	RankShift Rank
	FileShift File
}

func NewCoordinatesShift(r Rank, f File) CoordinatesShift {
	return CoordinatesShift{RankShift: r, FileShift: f}
}
