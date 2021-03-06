package main

import (
	"log"

	"gofigure/pkg"
	p "gofigure/pkg/primitives"
)

const (
	width  = 800
	height = 600
	scale = float32(500)
)

type Ex01Scene struct {
	pkg.BaseScene
}

func main() {
	window := pkg.Init(width, height)
	defer pkg.Terminate()

	s := &Ex01Scene{}
	s.Program = pkg.UseDefaultProgram()

	ratio := float32(width) / height
	s.SetMvpOrtho(ratio, scale)

	tp, err := pkg.LoadTexPack("resources/assets.tp")
	if err != nil {
		log.Panic(err)
	}

	_, err = s.SetTextureFromTexPack(tp)
	if err != nil {
		log.Panic(err)
	}

	shipBounds := tp.SpriteBounds01("ship/hull")
	q := p.MakeQuad(
		p.MakeRectf(0,0, 130, 344), p.RGBW, shipBounds,
	)
	s.AddEntity(q)

	emptyBounds := tp.SpriteBounds01("empty")
	t := p.MakeTriangle(
		p.MakePoint(p.XY(-100, -100), p.RGBR, emptyBounds.Min),
		p.MakePoint(p.XY(-100, -50), p.RGBG, emptyBounds.Min),
		p.MakePoint(p.XY(50, -100), p.RGBB, emptyBounds.Min),
	)
	s.AddEntity(t)

	pkg.MainLoop(window, s.Update, s.Render)
}

func (s *Ex01Scene) Update(timeDelta float64) {
	for _, e := range s.Entities {
		e.RotateBy(float32(timeDelta))
	}
}
