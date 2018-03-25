package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"gofigure/pkg"
)

func main() {
	var in, out, info, ext string

	flag.StringVar(&in, "in", "", "input directory")
	flag.StringVar(&out, "out", "", "output file")
	flag.StringVar(&info, "info", "", "describe file")
	flag.StringVar(&ext, "ext", "", "extract -info pack into file")
	flag.Parse()

	if info != "" {
		tp, err := pkg.LoadTexPack(info)
		if err != nil {
			panic(err.Error())
		}
		tp.PrintInfo()

		if ext != "" {
			fh, err := os.Create(ext)
			defer fh.Close()
			if err != nil {
				panic(err.Error())
			}
			err = png.Encode(fh, tp.Data())
			if err != nil {
				panic(err.Error())
			}
		}

		return
	}

	if in == "" || out == "" {
		fmt.Printf("Error: must provide -in and -out arguments\n")
		os.Exit(1)
	}

	fmt.Printf(
		"Processing sprites in %s into %s.tp texture pack\n", in, out,
	)

	tp, err := pkg.MakeTexPack(in)
	if err != nil {
		panic(err.Error())
	}

	tp.PrintInfo()

	wrote, err := pkg.SaveTexPack(tp, out)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Wrote %d bytes\n", wrote)
}