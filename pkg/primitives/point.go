package primitives

import (
	"github.com/go-gl/mathgl/mgl32"
)

const POINT_FLAT_SIZE = VERTEX_SIZE

type Vec2 = mgl32.Vec2
type Vec3 = mgl32.Vec3
type Vec4 = mgl32.Vec4
type Mat4 = mgl32.Mat4

type Point struct {
	Vertex
	DrawPrimitive
}

func XY(x, y float32) Vec4         { return Vec4{x, y, 0, 1} }
func XYZ(x, y, z float32) Vec4     { return Vec4{x, y, z, 1} }
func RGB(r, g, b float32) Vec4     { return Vec4{r, g, b, 1} }
func RGBA(r, g, b, a float32) Vec4 { return Vec4{r, g, b, a} }
func UV(u, v float32) Vec2         { return Vec2{u, v} }

var (
	RGBW = Vec4{1,1,1,1}
	RGBR = Vec4{1,0,0,1}
	RGBG = Vec4{0,1,0,1}
	RGBB = Vec4{0,0,1,1}
)

func MakePoint(pos, col Vec4, uv mgl32.Vec2) *Point {
	p := &Point{}
	p.position = pos
	p.color = col
	p.uv = uv
	p.dirty = true
	p.syncFlat()
	return p
}

func (p *Point) syncFlat() {
	p.flat = []float32{
		p.position[0], p.position[1], p.position[2], p.position[3],
		p.color[0], p.color[1], p.color[2], p.color[3],
		p.uv[0], p.uv[1],
	}
}

func (p *Point) RotateZ(angle float32) {
	p.position = mgl32.HomogRotate3DZ(angle).Mul4x1(p.position)
	p.dirty = true
	p.Setup()
}

func (p *Point) Translate(delta Vec3) {
	p.position[0] += delta.X()
	p.position[1] += delta.Y()
	p.position[2] += delta.Z()
	p.dirty = true
	p.Setup()
}

func (p *Point) Setup() {
	if p.dirty || p.flat == nil {
		p.syncFlat()
		p.DrawPrimitive.Setup()
		p.dirty = false
	}
}

func (p *Point) Draw(mode int, first int) {
	p.Setup()
	p.DrawPrimitive.Draw(mode, first)
}
