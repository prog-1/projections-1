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
	screenWidth  = 640
	screenHeight = 480
)

type Point struct {
	x, y float64
}

type Game struct {
	a, b          Point
	img           *ebiten.Image
	width, height int
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
	}
}

func (g *Game) Update() error {
	return nil
}

func DrawCarpet(img *ebiten.Image, a, b Point, depth int) {
	if depth == 0 {
		return
	}
	x := (b.x) / 3
	y := (b.y) / 3
	ebitenutil.DrawRect(img, a.x+x, a.y+y, x, y, color.RGBA{R: 187, G: 245, B: 39, A: 255})
	DrawCarpet(img, Point{x: a.x, y: a.y}, Point{x: x, y: y}, depth-1)
	DrawCarpet(img, Point{x: a.x + x, y: a.y}, Point{x: x, y: y}, depth-1)
	DrawCarpet(img, Point{x: a.x + x*2, y: a.y}, Point{x: x, y: y}, depth-1)
	DrawCarpet(img, Point{x: a.x, y: a.y + y}, Point{x: x, y: y}, depth-1)
	DrawCarpet(img, Point{x: a.x + x*2, y: a.y + y}, Point{x: x, y: y}, depth-1)
	DrawCarpet(img, Point{x: a.x, y: a.y + y*2}, Point{x: x, y: y}, depth-1)
	DrawCarpet(img, Point{x: a.x + x, y: a.y + y*2}, Point{x: x, y: y}, depth-1)
	DrawCarpet(img, Point{x: a.x + x*2, y: a.y + y*2}, Point{x: x, y: y}, depth-1)
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, g.a.x, g.a.y, g.b.x, g.b.y, color.RGBA{R: 245, G: 114, B: 227, A: 255})
	DrawCarpet(screen, g.a, g.b, 6)
}
func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.RunGame(&Game{
		a:   Point{100, 100},
		b:   Point{300, 300},
		img: ebiten.NewImage(screenWidth, screenHeight)}); err != nil {
		log.Fatal(err)
	}
}
