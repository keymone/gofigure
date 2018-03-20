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

func XY(x, y float32) mgl32.Vec4         { return mgl32.Vec4{x, y, 0, 1} }
func XYZ(x, y, z float32) mgl32.Vec4     { return mgl32.Vec4{x, y, z, 1} }
func XYZW(x, y, z, w float32) mgl32.Vec4 { return mgl32.Vec4{x, y, z, w} }

func RGB(r, g, b float32) mgl32.Vec4     { return mgl32.Vec4{r, g, b, 1} }
func RGBA(r, g, b, a float32) mgl32.Vec4 { return mgl32.Vec4{r, g, b, a} }

func MakePoint(pos, col mgl32.Vec4) *Point {
	return &Point{position: pos, color: col}
}

func (p *Point) Setup() {
	if p.flat == nil {
		p.flat = []float32{
			p.position.X(),
			p.position.Y(),
			p.position.Z(),
			p.position.W(),
			p.color.X(),
			p.color.Y(),
			p.color.Z(),
			p.color.W(),
		}
	}

	if p.vbo == 0 {
		p.vbo = pkg.MakeVbo(p.flat)
	}

	if p.vao == 0 {
		p.vao = pkg.MakeVao(p.vbo, 1)
	}
}

func (p *Point) Draw() {
	p.Setup()

	gl.BindVertexArray(p.vao)
	gl.DrawArrays(gl.POINTS, 0, 1)
}

func DrawPoints(ps []*Point) {

}
