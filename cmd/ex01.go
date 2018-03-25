package main

import (
	"log"

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

	tp, err := pkg.LoadTexPack("resources/assets.tp")
	if err != nil {
		log.Panic(err)
	}

	_, err = s.SetTextureFromTexPack(tp)
	if err != nil {
		log.Panic(err)
	}

	s.AddEntity(
		p.MakeQuad(
			p.MakePointXYUV(-0.5, 0.5, 0, 1),
			p.MakePointXYUV(-0.1, 0.5, 1, 1),
			p.MakePointXYUV(-0.1, -0.5, 1, 0),
			p.MakePointXYUV(-0.5, -0.5, 0, 0),
		),
		p.MakeTriangle(
			p.MakePoint(p.XY(0, 0), p.RGB(1,0,0), p.UV(0,0)),
			p.MakePoint(p.XY(0, 0.5), p.RGB(0,1,0), p.UV(0,0)),
			p.MakePoint(p.XY(0.5, 0), p.RGB(0,0,1), p.UV(0,0)),
		),
	)

	ratio := float32(width) / height
	s.SetMvp(mgl32.Ortho(-ratio, ratio, -1, 1, 1, -1))

	pkg.MainLoop(window, s.Update, s.Render)
}
