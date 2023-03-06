package main

import (
	"image/color"
	"log"
	"math"
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
	width, height int
}

func DrawSierpinskiCarpet(img *ebiten.Image, a, b Point) {
	if math.Abs(b.x-a.x) < 5 && math.Abs(b.y-a.y) < 5 {
		return
	}
	x, y := (b.x-a.x)/3, (b.y-a.y)/3
	ebitenutil.DrawRect(img, a.x+x, a.y+y, x, y, color.White)
	DrawSierpinskiCarpet(img, a, Point{x: a.x + x, y: a.y + y})
	DrawSierpinskiCarpet(img, Point{x: a.x + x, y: a.y}, Point{x: b.x - x, y: a.y + y})
	DrawSierpinskiCarpet(img, Point{x: b.x - x, y: a.y}, Point{x: b.x, y: a.y + y})
	DrawSierpinskiCarpet(img, Point{x: b.x - x, y: a.y + y}, Point{x: b.x, y: b.y - y})
	DrawSierpinskiCarpet(img, Point{x: b.x - x, y: b.y - y}, b)
	DrawSierpinskiCarpet(img, Point{x: a.x + x, y: b.y - y}, Point{x: b.x - x, y: b.y})
	DrawSierpinskiCarpet(img, Point{x: a.x, y: b.y - y}, Point{x: a.x + x, y: b.y})
	DrawSierpinskiCarpet(img, Point{x: a.x, y: a.y + y}, Point{x: a.x + x, y: b.y - y})
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 100, 100, 300, 300, color.RGBA{0, 0, 255, 255})
	DrawSierpinskiCarpet(screen, Point{100, 100}, Point{400, 400})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
