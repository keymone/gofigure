package primitives

import "github.com/go-gl/mathgl/mgl32"

type Point struct {
	position mgl32.Vec4
	color    mgl32.Vec4
}

func MakePoint(x, y, z float32) Point {
	return Point{
		position: mgl32.Vec4{x, y, z, 1},
	}
}

func (p *Point) Draw() {

}

func DrawPoints(ps []*Point) {

}
