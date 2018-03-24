package pkg

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/slimsag/binpack"
)

// File layout: <version> <n sprites> <sprite data> <data size> <data>
// Sprite data: <key size> <key data> <x1> <y1> <x2> <y2>

type TexSprite = image.Rectangle

type TexPack struct {
	sprites map[string]TexSprite
	data *image.RGBA
	index []string
}

func (tp *TexPack) PrintInfo() {
	fmt.Printf("Texture pack info\n")
	fmt.Printf("Sprites: %d\n", len(tp.Sprites()))
	fmt.Printf("%+v\n", tp.Sprites())
	fmt.Printf("Texture bounds: %v\n", tp.Data().Bounds())
}

func (tp *TexPack) Sprites() map[string]TexSprite {
	return tp.sprites
}

func (tp *TexPack) Data() *image.RGBA {
	return tp.data
}

func (tp *TexPack) Len() int {
	if tp.index == nil {
		tp.index = make([]string, len(tp.sprites))
		i := 0
		for k := range tp.sprites {
			tp.index[i] = k
			i += 1
		}

		// Put widest sprite first for packing to work
		maxw := 0
		maxi := 0
		for i := range tp.index {
			k := tp.index[i]
			s := tp.sprites[k]
			if s.Dx() >= maxw {
				maxw = s.Dx()
				maxi = i
			}
		}
		kmaxi := tp.index[maxi]
		tp.index[maxi] = tp.index[0]
		tp.index[0] = kmaxi
	}
	return len(tp.index)
}

func (tp *TexPack) Size(n int) (width, height int) {
	s := tp.sprites[tp.index[n]]
	return s.Dx(), s.Dy()
}

func (tp *TexPack) Place(n, x, y int) {
	tp.sprites[tp.index[n]] = tp.sprites[tp.index[n]].Add(image.Pt(x, y))
}

type texFileHeader struct {
	version uint8
	spritesNum uint16
	dataSize uint32
}

type texSpriteHeader struct {
	keySize uint8
	key []byte
	x1, y1, x2, y2 uint32
}

// LoadTexPack loads texture pack from file
func LoadTexPack(file string) (*TexPack, error) {
	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("file read error %s: %s", file, err)
	}

	fh := texFileHeader{}
	buffer := bytes.NewBuffer(fileData)
	err = binary.Read(buffer, binary.LittleEndian, &fh.version)
	if err != nil {
		return nil, fmt.Errorf("parsing version: %s", err)
	}

	err = binary.Read(buffer, binary.LittleEndian, &fh.spritesNum)
	if err != nil {
		return nil, fmt.Errorf("parsing number of sprites: %s", err)
	}

	err = binary.Read(buffer, binary.LittleEndian, &fh.dataSize)
	if err != nil {
		return nil, fmt.Errorf("parsing data size: %s", err)
	}

	fmt.Printf("texpack header: %+v\n", fh)

	t := &TexPack{
		sprites: map[string]TexSprite{},
	}
	h := texSpriteHeader{}

	for i:=uint16(0); i<fh.spritesNum; i++ {
		err = binary.Read(buffer, binary.LittleEndian, &h.keySize)
		if err == nil {
			h.key = make([]byte, h.keySize)
			read, err := buffer.Read(h.key)
			if err == nil && read != int(h.keySize) {
				err = fmt.Errorf("bytes read do not match: %d, expected %d", read, h.keySize)
			}
		}
		if err == nil {
			err = binary.Read(buffer, binary.LittleEndian, &h.x1)
		}
		if err == nil {
			err = binary.Read(buffer, binary.LittleEndian, &h.y1)
		}
		if err == nil {
			err = binary.Read(buffer, binary.LittleEndian, &h.x2)
		}
		if err == nil {
			err = binary.Read(buffer, binary.LittleEndian, &h.y2)
		}
		if err != nil {
			return nil, fmt.Errorf("reading sprite %d: %s", i, err)
		}

		t.sprites[string(h.key[:h.keySize])] = image.Rect(
			int(h.x1), int(h.y1), int(h.x2), int(h.y2),
		)
	}

	data := make([]byte, fh.dataSize)
	read, err := buffer.Read(data)
	if err == nil && read != int(fh.dataSize) {
		err = fmt.Errorf("got %d bytes, expected %d", read, fh.dataSize)
	}
	if err != nil {
		return nil, fmt.Errorf("reading data: %s", err)
	}

	imgBuf := bytes.NewBuffer(data)
	img, err := png.Decode(imgBuf)
	if err != nil {
		return nil, fmt.Errorf("decoding data: %s", err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride data: %s", err)
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	t.data = rgba

	return t, nil
}

func SaveTexPack(tp *TexPack, path string) (int64, error) {
	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var dataBytes bytes.Buffer
	writer := bufio.NewWriter(&dataBytes)
	err = png.Encode(writer, tp.data)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return 0, err
	}
	writer.Flush()

	fileHeader := texFileHeader{
		version: uint8(0),
		spritesNum: uint16(len(tp.sprites)),
		dataSize: uint32(dataBytes.Len()),
	}

	err = binary.Write(file, binary.LittleEndian, fileHeader.version)
	if err == nil {
		err = binary.Write(file, binary.LittleEndian, fileHeader.spritesNum)
	}

	if err == nil {
		err = binary.Write(file, binary.LittleEndian, fileHeader.dataSize)
	}

	for k := range tp.sprites {
		s := tp.sprites[k]
		spriteHeader := texSpriteHeader{
			keySize: uint8(len(k)),
			key: []byte(k),
			x1: uint32(s.Min.X),
			y1: uint32(s.Min.Y),
			x2: uint32(s.Max.X),
			y2: uint32(s.Max.Y),
		}
		if err == nil {
			err = binary.Write(file, binary.LittleEndian, spriteHeader.keySize)
		}
		if err == nil {
			wrote, err := file.Write(spriteHeader.key)
			if err == nil && wrote != len(k) {
				err = fmt.Errorf("key wrote %d bytes, expected %d", wrote, len(k))
			}
		}
		if err == nil {
			err = binary.Write(file, binary.LittleEndian, spriteHeader.x1)
		}
		if err == nil {
			err = binary.Write(file, binary.LittleEndian, spriteHeader.y1)
		}
		if err == nil {
			err = binary.Write(file, binary.LittleEndian, spriteHeader.x2)
		}
		if err == nil {
			err = binary.Write(file, binary.LittleEndian, spriteHeader.y2)
		}
	}
	if err == nil {
		wrote, err := file.Write(dataBytes.Bytes())
		if err == nil && wrote != dataBytes.Len() {
			err = fmt.Errorf("key wrote %d bytes, expected %d", wrote, dataBytes.Len())
		}
	}

	fsize, err := file.Seek(0, 1)
	return fsize, nil
}

func MakeTexPack(dir string) (*TexPack, error) {
	tp := &TexPack{
		sprites: map[string]TexSprite{},
	}
	spriteImgs := map[string]image.Image{}

	filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && filepath.Ext(info.Name()) == ".png" {
				sname := path[strings.Index(path, "/")+1:len(path)-4]

				imgFile, err := os.Open(path)
				if err != nil {
					fmt.Fprintf(os.Stderr, "skipping %s: %s\n", path, err)
					return err
				}
				defer imgFile.Close()

				img, err := png.Decode(imgFile)
				if err != nil {
					fmt.Fprintf(os.Stderr, "skipping %s: %s\n", path, err)
					return err
				}

				spriteImgs[sname] = img
				tp.sprites[sname] = image.Rect(
					0,0, img.Bounds().Dx(), img.Bounds().Dy(),
				)
			}
			return nil
		},
	)

	w, h := binpack.Pack(tp)
	if w <= 0 {
		return nil, fmt.Errorf("texture packing broke: [%d %d]", w, h)
	}
	tp.data = image.NewRGBA(image.Rect(0, 0, w, h))

	for k := range tp.sprites {
		draw.Draw(
			tp.data,
			tp.sprites[k],
			spriteImgs[k],
			tp.sprites[k].Min,
			draw.Src,
		)
	}

	return tp, nil
}
