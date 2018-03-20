package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
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

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program)

		s.Render()

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
