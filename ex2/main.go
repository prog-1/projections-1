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
	pos1, pos2, pos3, pos4, pos5, pos6, pos7, pos8 Point
}

func (g *Game) RotationX(p, p2, p3, p4, p5, p6, p7, p8 Point) {
	ang := math.Pi / 120
	g.pos1.x = p.x*math.Cos(ang) - p.y*math.Sin(ang)
	g.pos1.y = p.x*math.Sin(ang) + p.y*math.Cos(ang)
	g.pos2.x = p2.x*math.Cos(ang) - p2.y*math.Sin(ang)
	g.pos2.y = p2.x*math.Sin(ang) + p2.y*math.Cos(ang)
	g.pos3.x = p3.x*math.Cos(ang) - p3.y*math.Sin(ang)
	g.pos3.y = p3.x*math.Sin(ang) + p3.y*math.Cos(ang)
	g.pos4.x = p4.x*math.Cos(ang) - p4.y*math.Sin(ang)
	g.pos4.y = p4.x*math.Sin(ang) + p4.y*math.Cos(ang)
	g.pos5.x = p5.x*math.Cos(ang) - p5.y*math.Sin(ang)
	g.pos5.y = p5.x*math.Sin(ang) + p5.y*math.Cos(ang)
	g.pos6.x = p6.x*math.Cos(ang) - p6.y*math.Sin(ang)
	g.pos6.y = p6.x*math.Sin(ang) + p6.y*math.Cos(ang)
	g.pos7.x = p7.x*math.Cos(ang) - p7.y*math.Sin(ang)
	g.pos7.y = p7.x*math.Sin(ang) + p7.y*math.Cos(ang)
	g.pos8.x = p8.x*math.Cos(ang) - p8.y*math.Sin(ang)
	g.pos8.y = p8.x*math.Sin(ang) + p8.y*math.Cos(ang)
}
func (g *Game) RotationY(p, p2, p3, p4, p5, p6, p7, p8 Point) {
	ang := math.Pi / 120
	g.pos1.x = p.x*math.Cos(ang) - p.z*math.Sin(ang)
	g.pos1.z = p.x*math.Sin(ang) + p.z*math.Cos(ang)
	g.pos2.x = p2.x*math.Cos(ang) - p2.z*math.Sin(ang)
	g.pos2.z = p2.x*math.Sin(ang) + p2.z*math.Cos(ang)
	g.pos3.x = p3.x*math.Cos(ang) - p3.z*math.Sin(ang)
	g.pos3.z = p3.x*math.Sin(ang) + p3.z*math.Cos(ang)
	g.pos4.x = p4.x*math.Cos(ang) - p4.z*math.Sin(ang)
	g.pos4.z = p4.x*math.Sin(ang) + p4.z*math.Cos(ang)
	g.pos5.x = p5.x*math.Cos(ang) - p5.z*math.Sin(ang)
	g.pos5.z = p5.x*math.Sin(ang) + p5.z*math.Cos(ang)
	g.pos6.x = p6.x*math.Cos(ang) - p6.z*math.Sin(ang)
	g.pos6.z = p6.x*math.Sin(ang) + p6.z*math.Cos(ang)
	g.pos7.x = p7.x*math.Cos(ang) - p7.z*math.Sin(ang)
	g.pos7.z = p7.x*math.Sin(ang) + p7.z*math.Cos(ang)
	g.pos8.x = p8.x*math.Cos(ang) - p8.z*math.Sin(ang)
	g.pos8.z = p8.x*math.Sin(ang) + p8.z*math.Cos(ang)
}
func (g *Game) RotationZ(p, p2, p3, p4, p5, p6, p7, p8 Point) {
	ang := math.Pi / 120
	g.pos1.y = p.y*math.Cos(ang) + p.z*math.Sin(ang)
	g.pos1.z = -p.y*math.Sin(ang) + p.z*math.Cos(ang)
	g.pos2.y = p2.y*math.Cos(ang) + p2.z*math.Sin(ang)
	g.pos2.z = -p2.y*math.Sin(ang) + p2.z*math.Cos(ang)
	g.pos3.y = p3.y*math.Cos(ang) + p3.z*math.Sin(ang)
	g.pos3.z = -p3.y*math.Sin(ang) + p3.z*math.Cos(ang)
	g.pos4.y = p4.y*math.Cos(ang) + p4.z*math.Sin(ang)
	g.pos4.z = -p4.y*math.Sin(ang) + p4.z*math.Cos(ang)
	g.pos5.y = p5.y*math.Cos(ang) + p5.z*math.Sin(ang)
	g.pos5.z = -p5.y*math.Sin(ang) + p5.z*math.Cos(ang)
	g.pos6.y = p6.y*math.Cos(ang) + p6.z*math.Sin(ang)
	g.pos6.z = -p6.y*math.Sin(ang) + p6.z*math.Cos(ang)
	g.pos7.y = p7.y*math.Cos(ang) + p7.z*math.Sin(ang)
	g.pos7.z = -p7.y*math.Sin(ang) + p7.z*math.Cos(ang)
	g.pos8.y = p8.y*math.Cos(ang) + p8.z*math.Sin(ang)
	g.pos8.z = -p8.y*math.Sin(ang) + p8.z*math.Cos(ang)
}

func (g *Game) Update() error {
	g.RotationX(g.pos1, g.pos2, g.pos3, g.pos4, g.pos5, g.pos6, g.pos7, g.pos8)
	g.RotationY(g.pos1, g.pos2, g.pos3, g.pos4, g.pos5, g.pos6, g.pos7, g.pos8)
	g.RotationZ(g.pos1, g.pos2, g.pos3, g.pos4, g.pos5, g.pos6, g.pos7, g.pos8)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	a, b := screenWidth, screenHeight
	ebitenutil.DrawLine(screen,
		(g.pos1.x/(g.pos1.z+1000))*-900+float64(a/2), (g.pos1.y/(g.pos1.z+1000))*-900+float64(b/2),
		(g.pos2.x/(g.pos2.z+1000))*-900+float64(a/2), (g.pos2.y/(g.pos2.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen,
		(g.pos2.x/(g.pos2.z+1000))*-900+float64(a/2), (g.pos2.y/(g.pos2.z+1000))*-900+float64(b/2),
		(g.pos3.x/(g.pos3.z+1000))*-900+float64(a/2), (g.pos3.y/(g.pos3.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen,
		(g.pos3.x/(g.pos3.z+1000))*-900+float64(a/2), (g.pos3.y/(g.pos3.z+1000))*-900+float64(b/2),
		(g.pos4.x/(g.pos4.z+1000))*-900+float64(a/2), (g.pos4.y/(g.pos4.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen,
		(g.pos4.x/(g.pos4.z+1000))*-900+float64(a/2), (g.pos4.y/(g.pos4.z+1000))*-900+float64(b/2),
		(g.pos1.x/(g.pos1.z+1000))*-900+float64(a/2), (g.pos1.y/(g.pos1.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})

	ebitenutil.DrawLine(screen,
		(g.pos5.x/(g.pos5.z+1000))*-900+float64(a/2), (g.pos5.y/(g.pos5.z+1000))*-900+float64(b/2),
		(g.pos6.x/(g.pos6.z+1000))*-900+float64(a/2), (g.pos6.y/(g.pos6.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen,
		(g.pos6.x/(g.pos6.z+1000))*-900+float64(a/2), (g.pos6.y/(g.pos6.z+1000))*-900+float64(b/2),
		(g.pos8.x/(g.pos8.z+1000))*-900+float64(a/2), (g.pos8.y/(g.pos8.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen,
		(g.pos8.x/(g.pos8.z+1000))*-900+float64(a/2), (g.pos8.y/(g.pos8.z+1000))*-900+float64(b/2),
		(g.pos7.x/(g.pos7.z+1000))*-900+float64(a/2), (g.pos7.y/(g.pos7.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen,
		(g.pos7.x/(g.pos7.z+1000))*-900+float64(a/2), (g.pos7.y/(g.pos7.z+1000))*-900+float64(b/2),
		(g.pos5.x/(g.pos5.z+1000))*-900+float64(a/2), (g.pos5.y/(g.pos5.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})

	ebitenutil.DrawLine(screen,
		(g.pos1.x/(g.pos1.z+1000))*-900+float64(a/2), (g.pos1.y/(g.pos1.z+1000))*-900+float64(b/2),
		(g.pos5.x/(g.pos5.z+1000))*-900+float64(a/2), (g.pos5.y/(g.pos5.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen,
		(g.pos2.x/(g.pos2.z+1000))*-900+float64(a/2), (g.pos2.y/(g.pos2.z+1000))*-900+float64(b/2),
		(g.pos6.x/(g.pos6.z+1000))*-900+float64(a/2), (g.pos6.y/(g.pos6.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen,
		(g.pos3.x/(g.pos3.z+1000))*-900+float64(a/2), (g.pos3.y/(g.pos3.z+1000))*-900+float64(b/2),
		(g.pos8.x/(g.pos8.z+1000))*-900+float64(a/2), (g.pos8.y/(g.pos8.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen,
		(g.pos4.x/(g.pos4.z+1000))*-900+float64(a/2), (g.pos4.y/(g.pos4.z+1000))*-900+float64(b/2),
		(g.pos7.x/(g.pos7.z+1000))*-900+float64(a/2), (g.pos7.y/(g.pos7.z+1000))*-900+float64(b/2), color.RGBA{R: 227, G: 76, B: 235, A: 255})
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano()) /////////////    Pos7________Pos8
	if err := ebiten.RunGame(&Game{  /////////////   /|         / |
		pos1: Point{x: 100, y: 100, z: 100},   //Pos5_|_______Pos6|
		pos2: Point{x: 100, y: -100, z: 100},  //  |  |       |   |
		pos3: Point{x: -100, y: -100, z: 100}, //  |  Pos4____|___| Pos3
		pos4: Point{x: -100, y: 100, z: 100},  //  | /        |  /
		pos5: Point{x: 100, y: 100, z: -100},  //  Pos1_______Pos2
		pos6: Point{x: 100, y: -100, z: -100},
		pos7: Point{x: -100, y: 100, z: -100},
		pos8: Point{x: -100, y: -100, z: -100}}); err != nil {
		log.Fatal(err)
	}
}
