package main

import (
	"gofigure/pkg"
	p "gofigure/pkg/primitives"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	width  = 800
	height = 600
)

func main() {
	window := pkg.Init(width, height)
	defer pkg.Terminate()

	s := pkg.MakeBaseScene()
	s.AddEntity(
		p.MakeTriangle(
			p.MakePoint(p.XY(0, 0), p.RGB(1, 0, 0), p.UV(0,0)),
			p.MakePoint(p.XY(0, 0.5), p.RGB(0, 1, 0), p.UV(0,0)),
			p.MakePoint(p.XY(0.5, 0.5), p.RGB(0, 0, 1), p.UV(0,0)),
		),
		p.MakeTriangle(
			p.MakePoint(p.XY(0, 0), p.RGB(1, 0, 0), p.UV(0,0)),
			p.MakePoint(p.XY(0, -0.5), p.RGB(0, 1, 0), p.UV(0,0)),
			p.MakePoint(p.XY(-0.5, -0.5), p.RGB(0, 0, 1), p.UV(0,0)),
		),
		p.MakePoint(p.XY(0, 0), p.RGB(1, 1, 1), p.UV(0,0)),
		p.MakePoint(p.XY(.9, 0), p.RGB(1, 1, 1), p.UV(0,0)),
		p.MakePoint(p.XY(0, .9), p.RGB(1, 1, 1), p.UV(0,0)),
		p.MakePoint(p.XY(.9, .9), p.RGB(1, 1, 1), p.UV(0,0)),
		p.MakeQuad(
			p.MakePoint(p.XY(-0.5, 0.5), p.RGB(1, 0, 0), p.UV(0,0)),
			p.MakePoint(p.XY(-0.1, 0.5), p.RGB(0, 0, 1), p.UV(0,0)),
			p.MakePoint(p.XY(-0.1, -0.5), p.RGB(0, 1, 0), p.UV(0,0)),
			p.MakePoint(p.XY(-0.5, -0.5), p.RGB(1, 0, 1), p.UV(0,0)),
		),
	)

	ratio := float32(width) / height
	s.SetMvp(mgl32.Ortho(-ratio, ratio, -1, 1, 1, -1))
	s.SetTexture("resources/empty.rgb")

	pkg.MainLoop(window, s.Update, s.Render)
}
