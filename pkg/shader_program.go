package pkg

import (
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	vertexShaderSource = `
    #version 410

    uniform mat4 mvp;

    layout(location=0) in vec4 vp;
    layout(location=1) in vec4 vc;

    out vec4 color;

    void main() {
      gl_Position = mvp * vp;
      color = vc;
    }
  ` + "\x00"

	fragmentShaderSource = `
    #version 410

    in vec4 color;
    out vec4 fragColor;

    void main() {
      fragColor = color;
    }
  ` + "\x00"
)

func UseDefaultProgram() uint32 {
	program := MakeProgram(vertexShaderSource, fragmentShaderSource)
	gl.UseProgram(program)
	return program
}

func MakeProgram(vsrc, fsrc string) uint32 {
	vertexShader, err := CompileShader(vsrc, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := CompileShader(fsrc, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		logStr := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(logStr))

		log.Panicf("failed to link program: %v", logStr)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}

func CompileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logStr := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logStr))

		log.Panicf("failed to compile %v: %v", source, logStr)
	}

	return shader, nil
}
