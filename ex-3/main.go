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
	x, y, z float64
}

type Game struct {
	width, height int
	verts         []*Point
	planes        [][4]int
}

func Rotate(a *Point, angle float64) {
	a.x = a.x*math.Cos(angle) - a.y*math.Sin(angle)
	a.y = a.x*math.Sin(angle) + a.y*math.Cos(angle)

	a.x = a.x*math.Cos(angle) - a.z*math.Sin(angle)
	a.z = a.x*math.Sin(angle) + a.z*math.Cos(angle)

	a.y = a.y*math.Cos(angle) + a.z*math.Sin(angle)
	a.z = -a.y*math.Sin(angle) + a.z*math.Cos(angle)
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		verts: []*Point{
			{100, 100, 100},
			{100, -100, 100},
			{-100, -100, 100},
			{-100, 100, 100},
			{100, 100, -100},
			{100, -100, -100},
			{-100, -100, -100},
			{-100, 100, -100},
		},
		planes: [][4]int{
			{0, 1, 2, 3}, {7, 6, 5, 4}, {2, 6, 7, 3}, {3, 7, 4, 0}, {1, 0, 4, 5}, {6, 2, 1, 5},
		},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	for _, v := range g.verts {
		Rotate(v, math.Pi/360)
	}
	return nil
}

func Cross(a, b *Point) *Point {
	return &Point{x: a.y*b.z - a.z*b.y, y: a.z*b.x - a.x*b.z, z: a.x*b.y - a.y*b.x}
}

func Sub(a, b *Point) *Point {
	return &Point{x: a.x - b.x, y: a.y - b.y, z: a.z - b.z}
}

func Dot(a, b *Point) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (g *Game) DrawLine(img *ebiten.Image, a, b *Point) {
	z1 := a.z + 500
	z2 := b.z + 500
	x1 := a.x / z1
	y1 := a.y / z1
	x2 := b.x / z2
	y2 := b.y / z2
	ebitenutil.DrawLine(img, x1*500+float64(g.width)/2, y1*500+float64(g.height)/2, x2*500+float64(g.width)/2, y2*500+float64(g.height)/2, color.White)
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.planes {
		normal := Cross(Sub(g.verts[v[3]], g.verts[v[0]]), Sub(g.verts[v[1]], g.verts[v[0]]))
		if Dot(normal, Sub(g.verts[v[0]], &Point{0, 0, -500})) < 0 {
			g.DrawLine(screen, g.verts[v[0]], g.verts[v[1]])
			g.DrawLine(screen, g.verts[v[1]], g.verts[v[2]])
			g.DrawLine(screen, g.verts[v[2]], g.verts[v[3]])
			g.DrawLine(screen, g.verts[v[3]], g.verts[v[0]])
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
