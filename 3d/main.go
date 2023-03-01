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
	Joints [][4]int
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
		a := crossProduct(c.Points[v[0]], c.Points[v[1]])
		if a.z > 0 {
			continue
		}
		for i := 1; i <= 3; i++ {

			z1 := c.Points[v[i-1]].z + 500
			z2 := c.Points[v[i]].z + 500
			x1 := c.Points[v[i-1]].x / z1
			y1 := c.Points[v[i-1]].y / z1
			x2 := c.Points[v[i]].x / z2
			y2 := c.Points[v[i]].y / z2
			ebitenutil.DrawLine(screen, x1*500+Width/2, y1*500+Height/2, x2*500+Width/2, y2*500+Height/2, color.White)
		}
		z1 := c.Points[v[0]].z + 500
		z2 := c.Points[v[3]].z + 500
		x1 := c.Points[v[0]].x / z1
		y1 := c.Points[v[0]].y / z1
		x2 := c.Points[v[3]].x / z2
		y2 := c.Points[v[3]].y / z2
		ebitenutil.DrawLine(screen, x1*500+Width/2, y1*500+Height/2, x2*500+Width/2, y2*500+Height/2, color.White)
	}

	// for _, v := range c.Joints {
	// }
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

func crossProduct(u, v Point) Point {
	return Point{u.y*v.z - u.z*v.y, u.z*v.x - u.x*v.z, u.x*v.y - u.y*v.x}
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
		Joints: [][4]int{
			{4, 0, 1, 6},
			{3, 5, 7, 2},
			{3, 4, 6, 5},
			{2, 0, 4, 3},
			{7, 1, 0, 2},
			{7, 5, 6, 1},
		},
	}
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{c}); err != nil {
		log.Fatal(err)
	}
}
