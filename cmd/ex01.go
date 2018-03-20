package main

import (
	"gofigure/pkg"
	prim "gofigure/pkg/primitives"
)

const (
	width  = 800
	height = 600
)

func main() {
	window := pkg.Init(width, height)
	defer pkg.Terminate()

	pkg.UseDefaultProgram()

	tri := prim.MakeTriangle(
		prim.MakePoint2(0, 0),
		prim.MakePoint2(0, 0.9),
		prim.MakePoint2(0.9, 0.9),
	)

	p1 := prim.MakePoint2(-0.1, -0.1)

	s := pkg.MakeBaseScene()
	s.AddEntity(tri)
	s.AddEntity(p1)

	// projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(width)/height, 0.1, 10.0)
	// camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	// model := mgl32.Ident4()

	// mvp = projection * camera * model
	// mvpUniform := gl.GetUniformLocation(program, gl.Str("mvp\x00"))
	// gl.UniformMatrix4fv(modelUniform, 1, false, &mvp[0])

	pkg.MainLoop(window, s.Update, s.Render)
}
