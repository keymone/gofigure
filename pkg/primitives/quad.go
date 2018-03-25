package primitives

import (
	"github.com/go-gl/mathgl/mgl32"
)

const (
	QUAD_FLAT_SIZE = 2 * TRIANGLE_FLAT_SIZE
)

type Quad struct {
	tl *Triangle
	br *Triangle

	DrawPrimitive
}

func MakeQuadPPPP(tl, tr, br, bl *Point) *Quad {
	tlt := MakeTriangle(bl, tl, tr)
	brt := MakeTriangle(tr, br, bl)
	q := &Quad{tl: tlt, br: brt}
	q.setupFlat()
	return q
}

func MakeQuadRCR(pr Rectf, color mgl32.Vec4, tr Rectf) *Quad {
	pbl := MakePoint(mgl32.Vec4{pr.Min[0], pr.Min[1], 0, 1}, color, tr.Min)
	ptl := MakePoint(mgl32.Vec4{pr.Min[0], pr.Max[1], 0, 1}, color, mgl32.Vec2{tr.Min[0], tr.Max[1]})
	ptr := MakePoint(mgl32.Vec4{pr.Max[0], pr.Max[1], 0, 1}, color, tr.Max)
	pbr := MakePoint(mgl32.Vec4{pr.Max[0], pr.Min[1], 0, 1}, color, mgl32.Vec2{tr.Max[0], tr.Min[1]})

	tlt := MakeTriangle(pbl, ptl, ptr)
	brt := MakeTriangle(ptr, pbr, pbl)
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
