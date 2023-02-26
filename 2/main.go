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
	screenWidth  = 960
	screenHeight = 640
)

type Point struct {
	x, y, z float64
}

type Game struct {
	width, height int
	dots          []Point
}

func rotate(p Point, angle float64) Point {
	p.x = p.x*math.Cos(angle) - p.y*math.Sin(angle)
	p.y = p.x*math.Sin(angle) + p.y*math.Cos(angle)

	p.x = p.x*math.Cos(angle) - p.z*math.Sin(angle)
	p.z = p.x*math.Sin(angle) + p.z*math.Cos(angle)

	p.y = p.y*math.Cos(angle) + p.z*math.Sin(angle)
	p.z = p.z*math.Cos(angle) - p.y*math.Sin(angle)
	return p
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	for i := range g.dots {

		g.dots[i] = rotate(g.dots[i], math.Pi/180)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	c := color.White
	w, h := g.width/2, g.height/2
	ebitenutil.DrawLine(screen, g.dots[0].x+float64(w), g.dots[0].y+float64(h), g.dots[1].x+float64(w), g.dots[1].y+float64(h), c)
	ebitenutil.DrawLine(screen, g.dots[1].x+float64(w), g.dots[1].y+float64(h), g.dots[2].x+float64(w), g.dots[2].y+float64(h), c)
	ebitenutil.DrawLine(screen, g.dots[2].x+float64(w), g.dots[2].y+float64(h), g.dots[3].x+float64(w), g.dots[3].y+float64(h), c)
	ebitenutil.DrawLine(screen, g.dots[3].x+float64(w), g.dots[3].y+float64(h), g.dots[0].x+float64(w), g.dots[0].y+float64(h), c)

	ebitenutil.DrawLine(screen, g.dots[4].x+float64(w), g.dots[4].y+float64(h), g.dots[5].x+float64(w), g.dots[5].y+float64(h), c)
	ebitenutil.DrawLine(screen, g.dots[5].x+float64(w), g.dots[5].y+float64(h), g.dots[6].x+float64(w), g.dots[6].y+float64(h), c)
	ebitenutil.DrawLine(screen, g.dots[6].x+float64(w), g.dots[6].y+float64(h), g.dots[7].x+float64(w), g.dots[7].y+float64(h), c)
	ebitenutil.DrawLine(screen, g.dots[7].x+float64(w), g.dots[7].y+float64(h), g.dots[4].x+float64(w), g.dots[4].y+float64(h), c)

	ebitenutil.DrawLine(screen, g.dots[0].x+float64(w), g.dots[0].y+float64(h), g.dots[4].x+float64(w), g.dots[4].y+float64(h), c)
	ebitenutil.DrawLine(screen, g.dots[1].x+float64(w), g.dots[1].y+float64(h), g.dots[5].x+float64(w), g.dots[5].y+float64(h), c)
	ebitenutil.DrawLine(screen, g.dots[2].x+float64(w), g.dots[2].y+float64(h), g.dots[6].x+float64(w), g.dots[6].y+float64(h), c)
	ebitenutil.DrawLine(screen, g.dots[3].x+float64(w), g.dots[3].y+float64(h), g.dots[7].x+float64(w), g.dots[7].y+float64(h), c)

}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		dots: []Point{
			{-100, -100, -100},
			{-100, 100, -100},
			{100, 100, -100},
			{100, -100, -100},
			{-100, -100, 100},
			{-100, 100, 100},
			{100, 100, 100},
			{100, -100, 100},
		},
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
