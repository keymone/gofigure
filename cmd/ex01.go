package main

import (
	"gofigure/pkg"
	prim "gofigure/pkg/primitives"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	width  = 800
	height = 600
)

func main() {
	window := pkg.Init(width, height)
	defer pkg.Terminate()

	program := pkg.UseDefaultProgram()

	s := pkg.MakeBaseScene()
	s.AddEntity(
		prim.MakeTriangle(
			prim.MakePoint(prim.XY(0, 0), prim.RGB(1, 0, 0)),
			prim.MakePoint(prim.XY(0, 0.5), prim.RGB(0, 1, 0)),
			prim.MakePoint(prim.XY(0.5, 0.5), prim.RGB(0, 0, 1)),
		),
		prim.MakeTriangle(
			prim.MakePoint(prim.XY(0, 0), prim.RGB(1, 0, 0)),
			prim.MakePoint(prim.XY(0, -0.5), prim.RGB(0, 1, 0)),
			prim.MakePoint(prim.XY(-0.5, -0.5), prim.RGB(0, 0, 1)),
		),
		prim.MakePoint(prim.XY(0, 0), prim.RGB(1, 1, 1)),
		prim.MakePoint(prim.XY(.9, 0), prim.RGB(1, 1, 1)),
		prim.MakePoint(prim.XY(0, .9), prim.RGB(1, 1, 1)),
		prim.MakePoint(prim.XY(.9, .9), prim.RGB(1, 1, 1)),
	)

	ratio := float32(width) / height
	ortho := mgl32.Ortho(-ratio, ratio, -1, 1, 1, -1)
	// projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(width)/height, 0.1, 10.0)
	// camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	mvp := ortho // projection.Mul4(camera)
	mvpUniform := gl.GetUniformLocation(program, gl.Str("mvp\x00"))
	gl.UniformMatrix4fv(mvpUniform, 1, false, &mvp[0])

	pkg.MainLoop(window, s.Update, s.Render)
}
