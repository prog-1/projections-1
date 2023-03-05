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
		p      []point
		planes [][]int
	}
)

func (g *game) rotateX() {
	for i, v := range g.p {
		g.p[i].x = v.x*math.Cos(0.0025) - v.y*math.Sin(0.0025)
		g.p[i].y = v.x*math.Sin(0.0025) + v.y*math.Cos(0.0025)
	}
}

func (g *game) rotateY() {
	for i, v := range g.p {
		g.p[i].x = v.x*math.Cos(0.00474533) - v.z*math.Sin(0.00474533)
		g.p[i].z = v.x*math.Sin(0.00474533) + v.z*math.Cos(0.00474533)
	}

}
func (g *game) rotateZ() {
	for i, v := range g.p {
		g.p[i].y = v.y*math.Cos(-0.00474533) - v.z*math.Sin(-0.00474533)
		g.p[i].z = v.y*math.Sin(-0.00474533) + v.z*math.Cos(-0.00474533)
	}

}

func Cross(a, b point) point {
	return point{
		a.y*b.z - b.y*a.z,
		a.z*b.x - b.z*a.x,
		a.x*b.y - b.x*a.y,
	}
}
func Dot(a, b point) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }

func (g *game) Update() error {
	g.rotateX()
	g.rotateY()
	g.rotateZ()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.planes); i += 1 {
		a := point{
			g.p[g.planes[i][1]].x - g.p[g.planes[i][0]].x,
			g.p[g.planes[i][1]].y - g.p[g.planes[i][0]].y,
			g.p[g.planes[i][1]].z - g.p[g.planes[i][0]].z,
		}

		b := point{
			g.p[g.planes[i][2]].x - g.p[g.planes[i][1]].x,
			g.p[g.planes[i][2]].y - g.p[g.planes[i][1]].y,
			g.p[g.planes[i][2]].z - g.p[g.planes[i][1]].z,
		}
		center := point{
			x: (g.p[g.planes[i][2]].x + g.p[g.planes[i][1]].x + g.p[g.planes[i][0]].x) / 3,
			y: (g.p[g.planes[i][2]].y + g.p[g.planes[i][1]].y + g.p[g.planes[i][0]].y) / 3,
			z: (g.p[g.planes[i][2]].x + g.p[g.planes[i][1]].x + g.p[g.planes[i][0]].z) / 3,
		}
		screen.Set(int(center.x)+screenWidth/2, int(center.y)+screenHeight/2, c)
		cross := Cross(a, b)
		if Dot(point{0, 0, 1}, cross) < 0 {
			ebitenutil.DrawLine(screen,
				(g.p[g.planes[i][0]].x/(g.p[g.planes[i][0]].z+1500))*-5000+float64(screenWidth/2),
				(g.p[g.planes[i][0]].y/(g.p[g.planes[i][0]].z+1500))*-5000+float64(screenHeight/2),
				(g.p[g.planes[i][1]].x/(g.p[g.planes[i][1]].z+1500))*-5000+float64(screenWidth/2),
				(g.p[g.planes[i][1]].y/(g.p[g.planes[i][1]].z+1500))*-5000+float64(screenHeight/2),
				color.White)
			ebitenutil.DrawLine(screen,
				(g.p[g.planes[i][1]].x/(g.p[g.planes[i][1]].z+1500))*-5000+float64(screenWidth/2),
				(g.p[g.planes[i][1]].y/(g.p[g.planes[i][1]].z+1500))*-5000+float64(screenHeight/2),
				(g.p[g.planes[i][2]].x/(g.p[g.planes[i][2]].z+1500))*-5000+float64(screenWidth/2),
				(g.p[g.planes[i][2]].y/(g.p[g.planes[i][2]].z+1500))*-5000+float64(screenHeight/2),
				color.White)
			ebitenutil.DrawLine(screen,
				(g.p[g.planes[i][2]].x/(g.p[g.planes[i][2]].z+1500))*-5000+float64(screenWidth/2),
				(g.p[g.planes[i][2]].y/(g.p[g.planes[i][2]].z+1500))*-5000+float64(screenHeight/2),
				(g.p[g.planes[i][0]].x/(g.p[g.planes[i][0]].z+1500))*-5000+float64(screenWidth/2),
				(g.p[g.planes[i][0]].y/(g.p[g.planes[i][0]].z+1500))*-5000+float64(screenHeight/2),
				color.White)
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
	verts := make([]point, 10)
	verts[0] = point{x: 100, z: 50}
	{
		const fi = math.Pi / 5
		cosfi, sinfi := math.Cos(fi), math.Sin(fi)
		for i := 1; i < len(verts); i++ {
			verts[i] = point{
				x: verts[i-1].x*cosfi - verts[i-1].y*sinfi,
				y: verts[i-1].x*sinfi + verts[i-1].y*cosfi,
				z: -verts[i-1].z,
			}
		}
	}
	facets := make([][]int, 10, 12)
	for i := range verts {
		if i%2 == 0 {
			facets[i] = []int{i, (i + 1) % 10, (i + 2) % 10}
		} else {
			facets[i] = []int{(i + 2) % 10, (i + 1) % 10, i}
		}
	}
	verts = append(verts,
		point{z: (50*math.Sqrt(5)-1)/2 + 50},
		point{z: -(50*math.Sqrt(5)-1)/2 - 50},
	)
	for i := 0; i < 10; i += 2 {
		facets = append(facets, []int{10, i, (i + 2) % 10})
		facets = append(facets, []int{(i + 3) % 10, i + 1, 11})
		//facets = append(facets, []int{(i + 2 - 1) % 10, (i + 1) % 10, 11})
	}
	return &game{
		p:      verts,
		planes: facets,
	}
}
