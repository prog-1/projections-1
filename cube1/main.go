package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//Implement a program that projects a 3d convex shape (e.g. a cube) into the plane $z = -1$ using a central projection with all facets visible.

const (
	screenWidth  = 640
	screenHeight = 480
)

type vector struct {
	x, y, z float64
}

type connection struct {
	a, b int
}

type game struct {
	points      []vector
	connections []connection
	c           vector
}

func (g *game) rotateX(rad float64) {
	for i, v := range g.points {
		g.points[i].x = v.x*math.Cos(rad) - v.y*math.Sin(rad)
		g.points[i].y = v.x*math.Sin(rad) + v.y*math.Cos(rad)
	}
}
func (g *game) rotateY(rad float64) {
	for i, v := range g.points {
		g.points[i].x = v.x*math.Cos(rad) - v.z*math.Sin(rad)
		g.points[i].z = v.x*math.Sin(rad) + v.z*math.Cos(rad)
	}
}
func (g *game) rotateZ(rad float64) {
	for i, v := range g.points {
		g.points[i].y = v.y*math.Cos(rad) - v.z*math.Sin(rad)
		g.points[i].z = v.y*math.Sin(rad) + v.z*math.Cos(rad)
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	for _, v := range g.points {
		v.x -= g.c.x
		v.y -= g.c.y
		v.z -= g.c.z
		g.rotateX(0.002) // sign affects the direction of rotation
		g.rotateY(0.001) // if /10, slow motion effect
		g.rotateZ(-0.002)
		v.x += g.c.x
		v.y += g.c.y
		v.z += g.c.z
	}
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	for _, c := range g.connections {
		ebitenutil.DrawLine(screen, g.points[c.a].x+screenWidth/2, g.points[c.a].y+screenHeight/2, g.points[c.b].x+screenWidth/2, g.points[c.b].y+screenHeight/2, color.White)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	points := []vector{
		{100, 100, 100},
		{200, 100, 100},
		{200, 200, 100},
		{100, 200, 100},
		{100, 100, 200},
		{200, 100, 200},
		{200, 200, 200},
		{100, 200, 200},
	}
	connections := []connection{
		{0, 4},
		{1, 5},
		{2, 6},
		{3, 7},
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 0},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 4},
	}
	var c vector
	for _, v := range points {
		c.x += v.x
		c.y += v.y
		c.z += v.z
	}
	c.x /= float64(len(points))
	c.y /= float64(len(points))
	c.z /= float64(len(points))

	if err := ebiten.RunGame(&game{points: points, c: c, connections: connections}); err != nil {
		log.Fatal(err)
	}
}
