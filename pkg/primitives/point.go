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
func UV(u, v float32) mgl32.Vec2 { return mgl32.Vec2{u, v} }

var (
	RGBZ = Vec4{0,0,0,0}
	RGBW = Vec4{1,1,1,1}
	RGBR = Vec4{1,0,0,1}
	RGBG = Vec4{0,1,0,1}
	RGBB = Vec4{0,0,1,1}
	UVZ = mgl32.Vec2{0,0}
)

type Rectf struct {
	Min, Max [2]float32
}

func MakeRectf(x1, y1, x2, y2 float32) Rectf {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return Rectf{
		Min: [2]float32{x1, y1},
		Max: [2]float32{x2, y2},
	}
}

func MakePoint(pos, col Vec4, uv mgl32.Vec2) *Point {
	p := &Point{}
	p.position = pos
	p.color = col
	p.uv = uv
	p.dirty = true
	p.syncFlat()
	return p
}

func MakePointXYUV(x, y, u, v float32) *Point {
	return MakePoint(XY(x, y), RGBZ, UV(u, v))
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

func (p *Point) Translate(x, y, z float32) {
	p.position[0] += x
	p.position[1] += y
	p.position[2] += z
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
