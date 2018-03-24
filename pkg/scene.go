package pkg

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	p "gofigure/pkg/primitives"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type IScene interface {
	Update(float64)
	Render()

	AddEntity(p.Drawer)
	RemoveEntity(p.Drawer) bool
}

type BaseScene struct {
	entities []p.Drawer
	program uint32
}

func MakeBaseScene() *BaseScene {
	return &BaseScene{
		program: UseDefaultProgram(),
	}
}

func (s *BaseScene) Update(timeDelta float64) {
}

func (s *BaseScene) Render() {
	gl.UseProgram(s.program)
	s.SetMode(0)
	for _, e := range s.entities {
		e.Draw(gl.TRIANGLES, 0)
	}
	s.SetMode(1)
	for _, e := range s.entities {
		e.Draw(gl.LINE_LOOP, 0)
	}
}

func (s *BaseScene) AddEntity(toAdd ...p.Drawer) {
	s.entities = append(s.entities, toAdd...)
}

func (s *BaseScene) RemoveEntity(toRemove p.Drawer) bool {
	idx := -1
	for i, e := range s.entities {
		if e == toRemove {
			idx = i
			break
		}
	}

	if idx >= 0 {
		s.entities = append(s.entities[:idx], s.entities[idx+1:]...)
		return true
	}

	return false
}

func (s *BaseScene) SetMode(mode int) {
	fmUniform := gl.GetUniformLocation(s.program, gl.Str("fragmentMode\x00"))
	gl.Uniform1i(fmUniform, int32(mode))
}

func (s *BaseScene) SetMvp(mvp mgl32.Mat4) {
	// projection := mgl32.Ortho(-ratio, ratio, -1, 1, 1, -1)
	// projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(width)/height, 0.1, 10.0)
	// camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	// model := mgl32.Ident4()

	mvpUniform := gl.GetUniformLocation(s.program, gl.Str("modelViewProjection\x00"))
	gl.UniformMatrix4fv(mvpUniform, 1, false, &mvp[0])
}

func (s *BaseScene) SetTexture(file string) (uint32, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return 0, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	img, err := png.Decode(imgFile)
	if err != nil {
		return 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}
