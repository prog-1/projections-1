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
	screenWidth  = 2560
	screenHeight = 1440
)

type point struct {
	x, y, z float64
}

type Game struct {
	o    [8]point
	join [][2]int
}

var col = color.RGBA{244, 212, 124, 255}

func NewGame(width, height int) *Game {
	return &Game{
		o: [8]point{
			{-300, -300, -300},
			{-300, 300, -300},
			{300, 300, -300},
			{300, -300, -300},
			{-300, -300, 300},
			{-300, 300, 300},
			{300, 300, 300},
			{300, -300, 300},
		},
		join: [][2]int{
			{0, 1},
			{0, 3},
			{2, 1},
			{2, 3},
			{1, 5},
			{0, 4},
			{2, 6},
			{3, 7},
			{4, 5},
			{5, 6},
			{6, 7},
			{7, 4},
		},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return screenWidth, screenHeight
}

func (g *Game) rotX() {
	for i, j := range g.o {
		g.o[i].y = j.y*math.Cos(math.Pi/180) + j.z*math.Sin(math.Pi/180)
		g.o[i].z = j.y*math.Sin(math.Pi/180) - j.z*math.Cos(math.Pi/180)
	}
}

func (g *Game) rotY() {
	for i, j := range g.o {
		g.o[i].x = j.x*math.Sin(math.Pi/180) - j.z*math.Cos(math.Pi/180)
		g.o[i].z = j.x*math.Cos(math.Pi/180) + j.z*math.Sin(math.Pi/180)
	}
}

func (g *Game) rotZ() {
	for i, j := range g.o {
		g.o[i].x = j.x*math.Cos(math.Pi/180) - j.y*math.Sin(math.Pi/180)
		g.o[i].y = j.x*math.Sin(math.Pi/180) + j.y*math.Cos(math.Pi/180)
	}
}

func (g *Game) Update() error {
	g.rotX()
	g.rotY()
	g.rotZ()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.join {
		ebitenutil.DrawLine(screen,
			(g.o[v[0]].x/(g.o[v[0]].z+1000))*900+float64(screenWidth/2),
			(g.o[v[0]].y/(g.o[v[0]].z+1000))*900+float64(screenHeight/2),
			(g.o[v[1]].x/(g.o[v[1]].z+1000))*900+float64(screenWidth/2),
			(g.o[v[1]].y/(g.o[v[1]].z+1000))*900+float64(screenHeight/2),
			col)
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
