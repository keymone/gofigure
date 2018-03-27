package primitives

import (
	"github.com/go-gl/mathgl/mgl32"
)

const (
	QUAD_FLAT_SIZE = 2 * TRIANGLE_FLAT_SIZE
)

type Quad struct {
	DrawPrimitive

	Anchor Vec3

	tl *Triangle
	br *Triangle
}

func MakeQuad(pr Rectf, color mgl32.Vec4, tr Rectf) *Quad {
	pbl := MakePoint(XY(pr.Min[0], pr.Min[1]), color, Vec2{tr.Min[0], tr.Max[1]})
	ptl := MakePoint(XY(pr.Min[0], pr.Max[1]), color, tr.Min)
	ptr := MakePoint(XY(pr.Max[0], pr.Max[1]), color, Vec2{tr.Max[0], tr.Min[1]})
	pbr := MakePoint(XY(pr.Max[0], pr.Min[1]), color, tr.Max)

	tlt := MakeTriangle(pbl, ptl, ptr)

	// Must copy point structs to avoid sharing
	brt := MakeTriangle(
		&Point{Vertex: ptr.Vertex, DrawPrimitive: ptr.DrawPrimitive},
		pbr,
		&Point{Vertex: pbl.Vertex, DrawPrimitive: pbl.DrawPrimitive})

	q := &Quad{tl: tlt, br: brt}
	q.Anchor = pbl.position.Vec3()
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

func (q *Quad) Translate(delta Vec3) {
	q.br.Translate(delta)
	q.tl.Translate(delta)
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

func (q *Quad) MoveTo(newPos Vec3) {
	delta := q.Anchor.Sub(newPos)
  q.Anchor = newPos
	q.Translate(delta)
}

func (q *Quad) RotateBy(angle float32) {
	q.Translate(Vec3{}.Sub(q.Anchor))
	q.RotateZ(angle)
	q.Translate(q.Anchor)
}
