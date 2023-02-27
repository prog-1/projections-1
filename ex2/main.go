package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	winTitle     = "Cube"
	screenWidth  = 1000
	screenHeight = 1000
	dpi          = 100
)

var c = color.RGBA{R: 255, G: 255, B: 255, A: 255}

type (
	point struct {
		x, y, z float64
	}
	game struct {
		p      [8]point
		planes [][2]int
	}
)

func (g *game) rotateX() {
	for i, v := range g.p {
		g.p[i].x = v.x*math.Cos(0.025) - v.y*math.Sin(0.025)
		g.p[i].y = v.x*math.Sin(0.025) + v.y*math.Cos(0.025)
	}
}

func (g *game) rotateY() {
	for i, v := range g.p {
		g.p[i].x = v.x*math.Cos(0.0174533) - v.z*math.Sin(0.0174533)
		g.p[i].z = v.x*math.Sin(0.0174533) + v.z*math.Cos(0.0174533)
	}

}
func (g *game) rotateZ() {
	for i, v := range g.p {
		g.p[i].y = v.y*math.Cos(-0.0174533) - v.z*math.Sin(-0.0174533)
		g.p[i].z = v.y*math.Sin(-0.0174533) + v.z*math.Cos(-0.0174533)
	}

}

func Cross(a, b point) bool {
	return a.x*b.y-a.y*b.x < 0

}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }

func (g *game) Update() error {
	g.rotateX()
	g.rotateY()
	g.rotateZ()
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {

	for i := 0; i < len(g.planes); i += 4 {
		a := point{
			g.p[g.planes[i][1]].x - g.p[g.planes[i][0]].x,
			g.p[g.planes[i][1]].x - g.p[g.planes[i][0]].y,
			g.p[g.planes[i][1]].x - g.p[g.planes[i][0]].z,
		}

		b := point{
			g.p[g.planes[i+1][1]].x - g.p[g.planes[i+1][0]].x,
			g.p[g.planes[i+1][1]].x - g.p[g.planes[i+1][0]].y,
			g.p[g.planes[i+1][1]].x - g.p[g.planes[i+1][0]].z,
		}
		if Cross(a, b) {
			for i1 := i; i1 < i+4; i1++ {
				ebitenutil.DrawLine(screen,
					(g.p[g.planes[i1][0]].x/(g.p[g.planes[i1][0]].z+1000))*-900+float64(screenWidth/2),
					(g.p[g.planes[i1][0]].y/(g.p[g.planes[i1][0]].z+1000))*-900+float64(screenHeight/2),
					(g.p[g.planes[i1][1]].x/(g.p[g.planes[i1][1]].z+1000))*-900+float64(screenWidth/2),
					(g.p[g.planes[i1][1]].y/(g.p[g.planes[i1][1]].z+1000))*-900+float64(screenHeight/2),
					color.White)
			}
		}
	}

}

func main() {
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizable(true)
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
func NewGame() *game {

	return &game{
		p: [8]point{
			{300, -300, -300},  //0
			{-300, -300, -300}, //1
			{-300, 300, -300},  //2
			{300, 300, -300},   //3

			{300, -300, 300},  //4
			{-300, -300, 300}, //5
			{-300, 300, 300},  //6
			{300, 300, 300},   //7
		},
		planes: [][2]int{
			// near plane
			{0, 1},
			{1, 2},
			{2, 3},
			{3, 0},
			// far plane
			{4, 7},
			{7, 6},
			{6, 5},
			{5, 4},
			//  top plane
			{5, 1},
			{1, 0},
			{0, 4},
			{4, 5},
			//left plane
			{6, 2},
			{2, 1},
			{1, 5},
			{5, 6},
			//right plane
			{4, 0},
			{0, 3},
			{3, 7},
			{7, 4},
			//bottom plane
			{6, 7},
			{7, 3},
			{3, 2},
			{2, 6},
		},
	}
}
