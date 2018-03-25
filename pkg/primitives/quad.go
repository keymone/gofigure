package primitives

const (
	QUAD_FLAT_SIZE = 2 * TRIANGLE_FLAT_SIZE
)

type Quad struct {
	tl *Triangle
	br *Triangle

	DrawPrimitive
}

func MakeQuad(tl, tr, br, bl *Point) *Quad {
	tlt := MakeTriangle(bl, tl, tr)
	brt := MakeTriangle(tr, br, bl)
	q := &Quad{tl: tlt, br: brt}
	q.setupFlat()
	return q
}

func (q *Quad) setupFlat() {
	if q.flat == nil {
		q.flat = make([]float32, QUAD_FLAT_SIZE)
		copy(q.flat[0:], q.tl.flat)
		copy(q.flat[TRIANGLE_FLAT_SIZE:], q.br.flat)
	}
}
