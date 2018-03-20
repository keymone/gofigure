package primitives

import (
	"gofigure/pkg"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Point struct {
	position mgl32.Vec4
	color    mgl32.Vec4

	flat []float32
	vbo  uint32
	vao  uint32
}

func MakePoint3(x, y, z float32) Point {
	return Point{position: mgl32.Vec4{x, y, z, 1}}
}

func MakePoint2(x, y float32) Point {
	return Point{position: mgl32.Vec4{x, y, 0, 1}}
}

func (p *Point) Setup() {
	if p.flat == nil {
		p.flat = []float32{
			p.position.X(),
			p.position.Y(),
			p.position.Z(),
		}
	}

	if p.vbo == 0 {
		p.vbo = pkg.MakeVbo(p.flat)
	}

	if p.vao == 0 {
		p.vao = pkg.MakeVao(p.vbo, 1, 0, 0)
	}
}

func (p *Point) Draw() {
	p.Setup()

	gl.BindVertexArray(p.vao)
	gl.DrawArrays(gl.POINTS, 0, int32(1))
}

func DrawPoints(ps []*Point) {

}
