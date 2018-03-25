package primitives

import (
	"github.com/go-gl/mathgl/mgl32"
)

const (
	// XYZW + RGBA + UV
	POINT_FLAT_SIZE = 4 + 4 + 2
)

type Point struct {
	position mgl32.Vec4
	color    mgl32.Vec4
	uv       mgl32.Vec2

	DrawPrimitive
}

func XY(x, y float32) mgl32.Vec4         { return mgl32.Vec4{x, y, 0, 1} }
func XYZ(x, y, z float32) mgl32.Vec4     { return mgl32.Vec4{x, y, z, 1} }
func RGB(r, g, b float32) mgl32.Vec4     { return mgl32.Vec4{r, g, b, 1} }
func RGBA(r, g, b, a float32) mgl32.Vec4 { return mgl32.Vec4{r, g, b, a} }
func UV(u, v float32) mgl32.Vec2 { return mgl32.Vec2{u, v} }

var (
	RGBZ = mgl32.Vec4{0,0,0,0}
	RGBW = mgl32.Vec4{1,1,1,1}
	RGBR = mgl32.Vec4{1,0,0,1}
	RGBG = mgl32.Vec4{0,1,0,1}
	RGBB = mgl32.Vec4{0,0,1,1}
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

func MakePoint(pos, col mgl32.Vec4, uv mgl32.Vec2) *Point {
	p := &Point{position: pos, color: col, uv: uv}
	p.setupFlat()
	return p
}

func MakePointXYUV(x, y, u, v float32) *Point {
	return MakePoint(XY(x, y), RGBZ, UV(u, v))
}

func (p *Point) setupFlat() {
	if p.flat == nil {
		p.flat = []float32{
			p.position[0], p.position[1], p.position[2], p.position[3],
			p.color[0], p.color[1], p.color[2], p.color[3],
			p.uv[0], p.uv[1],
		}
	}
}
