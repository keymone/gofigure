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

	shipBounds := tp.SpriteBounds01("ship/hull")
	shipTL := shipBounds.Min
	shipTL[1] = shipBounds.Max[1]
	shipBR := shipBounds.Max
	shipBR[1] = shipBounds.Min[1]

	emptyBounds := tp.SpriteBounds01("empty")

	s.AddEntity(
		p.MakeQuadRCR(
			p.MakeRectf(-.5,-.5, -.1, .5), p.RGBW, shipBounds,
		),
		p.MakeTriangle(
			p.MakePoint(p.XY(0, 0), p.RGBR, emptyBounds.Min),
			p.MakePoint(p.XY(0, 0.5), p.RGBG, emptyBounds.Min),
			p.MakePoint(p.XY(0.5, 0), p.RGBB, emptyBounds.Min),
		),
	)

	ratio := float32(width) / height
	s.SetMvp(mgl32.Ortho(-ratio, ratio, -1, 1, 1, -1))

	pkg.MainLoop(window, s.Update, s.Render)
}
