package pkg

import (
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func MakeVbo(vc []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vc), gl.Ptr(vc), gl.STATIC_DRAW)

	return vbo
}

func MakeVao(vbo uint32, n, stride int, offset int) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	var offsetPtr unsafe.Pointer
	if offset == 0 {
		offsetPtr = nil
	} else {
		offsetPtr = gl.PtrOffset(offset * 4)
	}

	gl.VertexAttribPointer(0, int32(n), gl.FLOAT, false, int32(stride*4), offsetPtr)

	return vao
}
