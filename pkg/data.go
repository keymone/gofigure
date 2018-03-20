package pkg

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
	gl.VertexAttribPointer(0, int32(n), gl.FLOAT, false, 8*4, nil)

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, int32(n), gl.FLOAT, false, 8*4, gl.PtrOffset(4*4))

	return vao
}
