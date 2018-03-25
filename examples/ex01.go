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
	shipTL := shipBounds.Min
	shipTL[1] = shipBounds.Max[1]
	shipBR := shipBounds.Max
	shipBR[1] = shipBounds.Min[1]

	//emptyBounds := tp.SpriteBounds01("empty")

	q := p.MakeQuadRCR(
		p.MakeRectf(0,0, 130, 344), p.RGBW, shipBounds,
	)

	s.AddEntity(
		q,
		//p.MakeTriangle(
		//	p.MakePoint(p.XY(0, 0), p.RGBR, emptyBounds.Min),
		//	p.MakePoint(p.XY(0, 0.5), p.RGBG, emptyBounds.Min),
		//	p.MakePoint(p.XY(0.5, 0), p.RGBB, emptyBounds.Min),
		//),
	)

	pkg.MainLoop(window, s.Update, s.Render)
}

func (s *Ex01Scene) Update(timeDelta float64) {
	for _, e := range s.Entities {
		e.Translate(-130/2, -344/2, 0)
		e.RotateZ(float32(timeDelta))
		e.Translate(130/2, 344/2, 0)
	}
}
