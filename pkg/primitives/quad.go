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
	q.syncFlat()
	q.dirty = true
	return q
}

func MakeQuadRCR(pr Rectf, color mgl32.Vec4, tr Rectf) *Quad {
	pbl := MakePoint(mgl32.Vec4{pr.Min[0], pr.Min[1], 0, 1}, color, mgl32.Vec2{tr.Min[0], tr.Max[1]})
	ptl := MakePoint(mgl32.Vec4{pr.Min[0], pr.Max[1], 0, 1}, color, tr.Min)
	ptr := MakePoint(mgl32.Vec4{pr.Max[0], pr.Max[1], 0, 1}, color, mgl32.Vec2{tr.Max[0], tr.Min[1]})
	pbr := MakePoint(mgl32.Vec4{pr.Max[0], pr.Min[1], 0, 1}, color, tr.Max)

	tlt := MakeTriangle(pbl, ptl, ptr)

	// Must copy point structs to avoid sharing
	brt := MakeTriangle(
		&Point{Vertex: ptr.Vertex, DrawPrimitive: ptr.DrawPrimitive},
		pbr,
		&Point{Vertex: pbl.Vertex, DrawPrimitive: pbl.DrawPrimitive})
	q := &Quad{tl: tlt, br: brt}
	q.syncFlat()
	q.dirty = true
	return q
}

func (q *Quad) syncFlat() {
	q.flat = make([]float32, QUAD_FLAT_SIZE)
	copy(q.flat[0:], q.tl.flat)
	copy(q.flat[TRIANGLE_FLAT_SIZE:], q.br.flat)
}

func (q *Quad) RotateZ(angle float32) {
	q.br.RotateZ(angle)
	q.tl.RotateZ(angle)
	q.dirty = true
	q.Setup()
}

func (q *Quad) Translate(x, y, z float32) {
	q.br.Translate(x, y, z)
	q.tl.Translate(x, y, z)
	q.dirty = true
	q.Setup()
}

func (q *Quad) Setup() {
	if q.dirty || q.flat == nil {
		q.syncFlat()
		q.DrawPrimitive.Setup()
		q.dirty = false
	}
}

func (q *Quad) Draw(mode int, first int) {
	q.Setup()
	q.DrawPrimitive.Draw(mode, first)
}