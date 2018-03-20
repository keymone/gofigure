package main

import (
	_ "github.com/go-gl/mathgl/mgl32"

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

	program := pkg.MakeDefaultProgram()

	tri := prim.MakeTriangle(
		prim.MakePoint2(0, 0),
		prim.MakePoint2(0, 0.9),
		prim.MakePoint2(0.9, 0.9),
	)

	p1 := prim.MakePoint2(-0.1, -0.1)

	s := pkg.MakeBaseScene()
	s.AddEntity(tri)
	s.AddEntity(p1)

	pkg.MainLoop(window, program, s.Update, s.Render)
}
