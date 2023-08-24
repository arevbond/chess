package color

type Color int

const (
	White Color = 0
	Black Color = 1
)

func Opposite(value Color) Color {
	if value == White {
		return Black
	}
	return White
}
