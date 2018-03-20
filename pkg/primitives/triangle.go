package primitives

import (
	"gofigure/pkg"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Triangle struct {
	a *Point
	b *Point
	c *Point

	flat []float32
	vbo  uint32
	vao  uint32
}

func MakeTriangle(a, b, c *Point) *Triangle {
	return &Triangle{
		a: a,
		b: b,
		c: c,
	}
}

func (p *Triangle) Draw() {
	p.Setup()

	gl.BindVertexArray(p.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(3))
}

func (p *Triangle) Setup() {
	if p.flat == nil {
		p.flat = []float32{
			p.a.position.X(),
			p.a.position.Y(),
			p.a.position.Z(),
			p.a.position.W(),
			p.b.position.X(),
			p.b.position.Y(),
			p.b.position.Z(),
			p.b.position.W(),
			p.c.position.X(),
			p.c.position.Y(),
			p.c.position.Z(),
			p.c.position.W(),
		}
	}

	if p.vbo == 0 {
		p.vbo = pkg.MakeVbo(p.flat)
	}

	if p.vao == 0 {
		p.vao = pkg.MakeVao(p.vbo, 3, 0, 0)
	}
}

func DrawTriangles(ps []*Triangle) {

}
