package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	sWidth  = 1000
	sHeight = 1000
)

// type point struct {
//     x, y int
// }

type Game struct {
	width, height int
	a, b, c       float64
	img           *ebiten.Image
}

var col = color.RGBA{244, 212, 124, 255}
var grey = color.RGBA{89, 96, 98, 255}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		a:      0,
		b:      0,
		c:      1000,
		// d:      400,
		img: ebiten.NewImage(sWidth, sHeight),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	return nil
}

func carpet(screen *ebiten.Image, a, b, c float64) {
	s := c / 3
	ebitenutil.DrawRect(screen, a+s, b+s, s, s, color.Black)

	if s >= 2 {
		carpet(screen, a, b, s)
		carpet(screen, a+s, b, s)
		carpet(screen, a+s*2, b, s)
		carpet(screen, a, b+s, s)
		carpet(screen, a+s*2, b+s, s)
		carpet(screen, a, b+s*2, s)
		carpet(screen, a+s, b+s*2, s)
		carpet(screen, a+s*2, b+s*2, s)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, g.a, g.b, g.c, g.c, col)
	carpet(screen, g.a, g.b, g.c)

}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(sWidth, sHeight)
	if err := ebiten.RunGame(NewGame(sWidth, sHeight)); err != nil {
		log.Fatal(err)
	}
}
