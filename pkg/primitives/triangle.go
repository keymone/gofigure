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
	t.syncFlat()
	return t
}

func (p *Triangle) syncFlat() {
	p.flat = make([]float32, TRIANGLE_FLAT_SIZE)
	copy(p.flat[0:], p.a.flat[:])
	copy(p.flat[POINT_FLAT_SIZE:], p.b.flat[:])
	copy(p.flat[POINT_FLAT_SIZE*2:], p.c.flat[:])
}

func (p *Triangle) RotateZ(angle float32) {
	p.a.RotateZ(angle)
	p.b.RotateZ(angle)
	p.c.RotateZ(angle)
	p.dirty = true
	p.Setup()
}

func (p *Triangle) Translate(x, y, z float32) {
	p.a.Translate(x, y, z)
	p.b.Translate(x, y, z)
	p.c.Translate(x, y, z)
	p.dirty = true
	p.Setup()
}

func (p *Triangle) Setup() {
	if p.dirty || p.flat == nil {
		p.syncFlat()
		p.DrawPrimitive.Setup()
		p.dirty = false
	}
}

func (p *Triangle) Draw(mode int, first int) {
	p.Setup()
	p.DrawPrimitive.Draw(mode, first)
}
