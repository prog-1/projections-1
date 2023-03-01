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
	planes        [][4]int
}

func sub(a, b Point) Point {
	return Point{b.x - a.x, b.y - a.y, b.z - a.z}
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
	w, h := g.width/2, g.height/2
	for _, j := range g.planes {
		g.check(screen, color.RGBA{255, 255, 255, 255}, w, h, j)
	}
}

func (g *Game) check(screen *ebiten.Image, c color.RGBA, hw, hh int, plane [4]int) {
	u, v := sub(g.dots[plane[0]], g.dots[plane[1]]), sub(g.dots[plane[1]], g.dots[plane[2]])
	if u.x*v.y-u.y*v.x < 0 {
		ebitenutil.DrawLine(screen, g.dots[plane[0]].x+float64(hw), g.dots[plane[0]].y+float64(hh), g.dots[plane[1]].x+float64(hw), g.dots[plane[1]].y+float64(hh), c)
		ebitenutil.DrawLine(screen, g.dots[plane[1]].x+float64(hw), g.dots[plane[1]].y+float64(hh), g.dots[plane[2]].x+float64(hw), g.dots[plane[2]].y+float64(hh), c)
		ebitenutil.DrawLine(screen, g.dots[plane[2]].x+float64(hw), g.dots[plane[2]].y+float64(hh), g.dots[plane[3]].x+float64(hw), g.dots[plane[3]].y+float64(hh), c)
		ebitenutil.DrawLine(screen, g.dots[plane[3]].x+float64(hw), g.dots[plane[3]].y+float64(hh), g.dots[plane[0]].x+float64(hw), g.dots[plane[0]].y+float64(hh), c)
	} else {
		ebitenutil.DrawLine(screen, g.dots[plane[0]].x+float64(hw), g.dots[plane[0]].y+float64(hh), g.dots[plane[1]].x+float64(hw), g.dots[plane[1]].y+float64(hh), color.RGBA{25, 25, 25, 150})
		ebitenutil.DrawLine(screen, g.dots[plane[1]].x+float64(hw), g.dots[plane[1]].y+float64(hh), g.dots[plane[2]].x+float64(hw), g.dots[plane[2]].y+float64(hh), color.RGBA{25, 25, 25, 150})
		ebitenutil.DrawLine(screen, g.dots[plane[2]].x+float64(hw), g.dots[plane[2]].y+float64(hh), g.dots[plane[3]].x+float64(hw), g.dots[plane[3]].y+float64(hh), color.RGBA{25, 25, 25, 150})
		ebitenutil.DrawLine(screen, g.dots[plane[3]].x+float64(hw), g.dots[plane[3]].y+float64(hh), g.dots[plane[0]].x+float64(hw), g.dots[plane[0]].y+float64(hh), color.RGBA{25, 25, 25, 150})
	}
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
		planes: [][4]int{
			{0, 1, 2, 3},
			{7, 6, 5, 4},
			{0, 4, 5, 1},
			{1, 5, 6, 2},
			{3, 2, 6, 7},
			{4, 0, 3, 7},
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
