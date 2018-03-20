package pkg

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func Init(width, height int) *glfw.Window {
	runtime.LockOSThread()

	window := initGlfw(width, height)

	initOpenGL()

	return window
}

func MainLoop(
	window *glfw.Window,
	update func(float64),
	render func(),
) {
	var newTime, oldTime float64
	for !window.ShouldClose() {
		newTime = glfw.GetTime()
		update(newTime - oldTime)
		oldTime = newTime

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		render()

		glfw.PollEvents()
		window.SwapBuffers()
	}
}

func Terminate() {
	glfw.Terminate()
}

func initGlfw(width, height int) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "None", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
}
