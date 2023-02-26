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
		join   [][2]int
		planes [][2]int
		sin    float64
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
	g.sin += -0.0174533

}
func Dot(a, b point) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func Cross(a, b point) point {
	return point{
		x: a.y*b.z - a.z*b.y,
		y: a.z*b.x - a.x*b.z,
		z: a.x*b.y - a.y*b.x,
	}
}
func Normalize(v point) point {
	magnitude := math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
	return point{
		x: v.x / magnitude,
		y: v.y / magnitude,
		z: v.z / magnitude,
	}
}

func IsFacingScreen(a, b point) bool {
	// Calculate the normal vector of the plane defined by a and b
	normal := Normalize(Cross(a, b))

	// Check if the dot product of the normal and the view vector is positive
	viewVector := point{x: 0, y: 0, z: 1}
	dotProduct := Dot(normal, viewVector)

	return dotProduct > 0
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }

func (g *game) Update() error {
	g.rotateX()
	g.rotateY()
	g.rotateZ()
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	for i := 0; i < 24; i += 4 {
		tmp1 := point{
			g.p[g.planes[i][1]].x - g.p[g.planes[i][0]].x,
			g.p[g.planes[i][1]].y - g.p[g.planes[i][0]].y,
			g.p[g.planes[i][1]].z - g.p[g.planes[i][0]].z}
		tmp2 := point{
			g.p[g.planes[i+1][1]].x - g.p[g.planes[i+1][0]].x,
			g.p[g.planes[i+1][1]].y - g.p[g.planes[i+1][0]].y,
			g.p[g.planes[i+1][1]].z - g.p[g.planes[i+1][0]].z}
		if !IsFacingScreen(tmp1, tmp2) {
			for i1 := i; i1 < i+4; i1++ {
				ebitenutil.DrawLine(screen,
					(g.p[g.planes[i1][0]].x/(g.p[g.planes[i1][0]].z+1000))*900+float64(screenWidth/2),
					(g.p[g.planes[i1][0]].y/(g.p[g.planes[i1][0]].z+1000))*900+float64(screenHeight/2),
					(g.p[g.planes[i1][1]].x/(g.p[g.planes[i1][1]].z+1000))*900+float64(screenWidth/2),
					(g.p[g.planes[i1][1]].y/(g.p[g.planes[i1][1]].z+1000))*900+float64(screenHeight/2),
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
		planes: [][2]int{
			// near plane
			{0, 1},
			{0, 3},
			{2, 1},
			{2, 3},
			// far plane
			{7, 6},
			{7, 4},
			{5, 6},
			{5, 4},
			//  top plane
			{2, 6},
			{2, 1},
			{5, 6},
			{5, 1},
			//left plane
			{5, 1},
			{5, 4},
			{0, 4},
			{0, 1},
			//right plane
			{6, 2},
			{6, 7},
			{3, 7},
			{3, 2},
			//bottom plane
			{3, 7},
			{3, 0},
			{4, 0},
			{4, 7},
		},
	}
}
