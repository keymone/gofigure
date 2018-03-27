package primitives

type Rectf struct {
	Min, Max [2]float32
}

func MakeRectf(x1, y1, x2, y2 float32) Rectf {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return Rectf{
		Min: [2]float32{x1, y1},
		Max: [2]float32{x2, y2},
	}
}