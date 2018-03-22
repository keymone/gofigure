package primitives

const (
	TRIANGLE_FLAT_SIZE = 3 * POINT_FLAT_SIZE
)

type Triangle struct {
	a *Point
	b *Point
	c *Point

	DrawPrimitive
}

func MakeTriangle(a, b, c *Point) *Triangle {
	t := &Triangle{a: a, b: b, c: c}
	t.setupFlat()
	return t
}

func (p *Triangle) setupFlat() {
	if p.flat == nil {
		p.flat = make([]float32, TRIANGLE_FLAT_SIZE)
		copy(p.flat[0:], p.a.flat[:])
		copy(p.flat[POINT_FLAT_SIZE:], p.b.flat[:])
		copy(p.flat[POINT_FLAT_SIZE*2:], p.c.flat[:])
	}
}
