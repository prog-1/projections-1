package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Height = 480
	Width  = 640
)

type Game struct {
	c cube
}

type Point struct {
	x, y, z float64
}

type cube struct {
	Points []Point
	Joints [][2]int
}

func (c *cube) RotateX(angle float64) {
	for i := range c.Points {
		c.Points[i].y = c.Points[i].y*math.Cos(angle) + c.Points[i].z*math.Sin(angle)
		c.Points[i].z = -c.Points[i].y*math.Sin(angle) + c.Points[i].z*math.Cos(angle)
	}
}

func (c *cube) RotateY(angle float64) {
	for i := range c.Points {
		c.Points[i].x = c.Points[i].x*math.Cos(angle) - c.Points[i].z*math.Sin(angle)
		c.Points[i].z = c.Points[i].x*math.Sin(angle) + c.Points[i].z*math.Cos(angle)
	}
}

func (c *cube) RotateZ(angle float64) {
	for i := range c.Points {
		c.Points[i].x = c.Points[i].x*math.Cos(angle) - c.Points[i].y*math.Sin(angle)
		c.Points[i].y = c.Points[i].x*math.Sin(angle) + c.Points[i].y*math.Cos(angle)
	}
}

func (c *cube) Draw(screen *ebiten.Image) {
	for _, v := range c.Joints {
		z1 := c.Points[v[0]].z + 500
		z2 := c.Points[v[1]].z + 500
		x1 := c.Points[v[0]].x / z1
		y1 := c.Points[v[0]].y / z1
		x2 := c.Points[v[1]].x / z2
		y2 := c.Points[v[1]].y / z2
		ebitenutil.DrawLine(screen, x1*500+Width/2, y1*500+Height/2, x2*500+Width/2, y2*500+Height/2, color.White)
	}
}

func (g *Game) Update() error {
	g.c.RotateX(math.Pi / 1000)
	g.c.RotateY(math.Pi / 750)
	g.c.RotateZ(math.Pi / 450)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.c.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func main() {
	c := cube{
		Points: []Point{
			{-100, -100, 100},
			{-100, 100, 100},
			{-100, -100, -100},
			{100, -100, -100},
			{100, -100, 100},
			{100, 100, -100},
			{100, 100, 100},
			{-100, 100, -100},
		},
		Joints: [][2]int{
			{0, 1},
			{0, 2},
			{0, 4},
			{1, 7},
			{1, 6},
			{2, 3},
			{2, 7},
			{3, 4},
			{3, 5},
			{4, 6},
			{5, 6},
			{5, 7},
		},
	}
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{c}); err != nil {
		log.Fatal(err)
	}
}
