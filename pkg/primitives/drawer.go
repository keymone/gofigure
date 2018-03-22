package primitives

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func MakeVbo(vc []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vc), gl.Ptr(vc), gl.STATIC_DRAW)

	return vbo
}

func MakeVao(vbo uint32, n int) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 10*4, gl.PtrOffset(0))

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 10*4, gl.PtrOffset(4*4))

	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 10*4, gl.PtrOffset(8*4))

	return vao
}

type Drawer interface {
	Draw(int, int)
}

type DrawPrimitive struct {
	flat []float32
	vbo  uint32
	vao  uint32
}

func (dp *DrawPrimitive) setupFlat() {}

func (dp *DrawPrimitive) Vertices() int {
	return len(dp.flat) / POINT_FLAT_SIZE
}

func (dp *DrawPrimitive) Setup() {
	dp.setupFlat()

	if dp.vbo == 0 {
		dp.vbo = MakeVbo(dp.flat)
	}

	if dp.vao == 0 {
		dp.vao = MakeVao(dp.vbo, dp.Vertices())
	}
}

func (dp *DrawPrimitive) Draw(mode int, first int) {
	dp.Setup()

	if mode == gl.FALSE {
		mode = gl.TRIANGLES
	}

	//log.Printf("drawing %v vertices of %+v", dp.Vertices(), dp)

	gl.BindVertexArray(dp.vao)
	gl.DrawArrays(uint32(mode), int32(first), int32(dp.Vertices()))
}