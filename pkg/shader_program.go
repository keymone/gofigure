package pkg

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	// uniform mat4 mvp;
	// in vec4 vc;
	vertexShaderSource = `
    #version 410
    in vec4 vp;
    out vec4 color;
    void main() {
      gl_Position = vp;
      color = vec4(1,0,0,0);
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

func UseDefaultProgram() {
	program := MakeProgram(vertexShaderSource, fragmentShaderSource)
	gl.UseProgram(program)
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

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
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

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
