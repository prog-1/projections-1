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
	screenWidth  = 960
	screenHeight = 640
)

type Point struct {
	x, y float64
}
type Game struct {
	width, height int
	x, y, s       float64
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	c := color.RGBA{255, 0, 0, 255}
	ebitenutil.DrawRect(screen, g.x, g.y, g.s, g.s, c)
	cube(screen, g.x, g.y, g.s, color.RGBA{0, 0, 0, 255}, 4)
}

func cube(screen *ebiten.Image, x, y, s float64, c color.RGBA, depth int) {
	if depth < 0 {
		return
	}
	l := s / 3
	ebitenutil.DrawRect(screen, x+l, y+l, l, l, c)
	cube(screen, x, y, l, c, depth-1)
	cube(screen, x+l, y, l, c, depth-1)
	cube(screen, x+l*2, y, l, c, depth-1)
	cube(screen, x, y+l, l, c, depth-1)
	cube(screen, x+l*2, y+l, l, c, depth-1)
	cube(screen, x, y+l*2, l, c, depth-1)
	cube(screen, x+l, y+l*2, l, c, depth-1)
	cube(screen, x+l*2, y+l*2, l, c, depth-1)
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		x:      0,
		y:      0,
		s:      300,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
