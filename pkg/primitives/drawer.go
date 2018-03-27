package primitives

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Drawer interface {
	Draw(int, int)

	RotateZ(float32)
	Translate(Vec3)

	RotateBy(float32)
	MoveTo(Vec3)
}

type DrawPrimitive struct {
	anchor Vec3
	flat   []float32
	vbo    uint32
	vao    uint32
	dirty  bool
}

func (dp *DrawPrimitive) syncVbo() {
	if dp.vbo == 0 {
		gl.GenBuffers(1, &dp.vbo)
	}
	gl.BindBuffer(gl.ARRAY_BUFFER, dp.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(dp.flat), gl.Ptr(dp.flat), gl.STATIC_DRAW)
}

func (dp *DrawPrimitive) syncVao() {
	if dp.vao == 0 {
		gl.GenVertexArrays(1, &dp.vao)
		gl.BindVertexArray(dp.vao)
		gl.BindBuffer(gl.ARRAY_BUFFER, dp.vbo)
		gl.EnableVertexAttribArray(0)
		gl.VertexAttribPointer(
			0, 4, gl.FLOAT, false, VERTEX_SIZE*4,
			gl.PtrOffset(0))
		gl.EnableVertexAttribArray(1)
		gl.VertexAttribPointer(
			1, 4, gl.FLOAT, false, VERTEX_SIZE*4,
			gl.PtrOffset(4*4))
		gl.EnableVertexAttribArray(2)
		gl.VertexAttribPointer(
			2, 2, gl.FLOAT, false, VERTEX_SIZE*4,
			gl.PtrOffset(8*4))
	}
}

func (dp *DrawPrimitive) Vertices() int {
	return len(dp.flat) / POINT_FLAT_SIZE
}

func (dp *DrawPrimitive) Setup() {
	if dp.dirty {
		dp.syncVbo()
		dp.syncVao()
		dp.dirty = false
	}
}

func (dp *DrawPrimitive) Draw(mode int, first int) {
	dp.Setup()

	if mode == gl.FALSE {
		mode = gl.TRIANGLES
	}

	gl.BindVertexArray(dp.vao)
	gl.DrawArrays(uint32(mode), int32(first), int32(dp.Vertices()))
}