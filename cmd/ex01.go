package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	_ "github.com/go-gl/mathgl/mgl32"

	"gofigure/pkg"
)

const (
	width  = 800
	height = 600
)

var (
	triangle = []float32{
		0, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
	}
)

func main() {
	window := pkg.Init(width, height)
	defer pkg.Terminate()

	program := pkg.MakeDefaultProgram()

	vbo := pkg.MakeVbo(triangle)
	vao := pkg.MakeVao(vbo, len(triangle)/3, 0, 0)

	for !window.ShouldClose() {
		draw(vao, window, program)
	}
}

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}
