package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//---------------------------Declaration--------------------------------

const (
	sW = 640
	sH = 480
)

type Game struct {
	width, height int //screen width and height
	//global variables
	r         rotator  //point rotation angle on all axis
	cp        point    //cube central point
	cube      [8]point // cube points
	faces     [][4]int //slice of cube faces
	faceEdges [][2]int //edge sequence for each face
}

type point struct {
	x, y, z float64
}

type rotator struct {
	x, y, z float64
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//cube point rotation
	for i := range g.cube {
		g.cube[i].rotate(g.r, g.cp)
	}
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {

	for i := range g.faces {

		u := sub(g.cube[g.faces[i][0]], g.cube[g.faces[i][1]])
		v := sub(g.cube[g.faces[i][2]], g.cube[g.faces[i][1]])
		n := crossProduct(u, v)

		if dotProduct(normalize(g.cp), n) < 0 { //n.z
			for j := range g.faceEdges {
				g.drawLine(screen, g.cube[g.faces[i][g.faceEdges[j][0]]], g.cube[g.faces[i][g.faceEdges[j][1]]])
			}
		}
	}

	g.drawMP(screen, g.cube[0], g.cube[2]) //used for debug

	// //default drawing
	// for i := range g.edges {
	// 	g.drawLine(screen, g.cube[g.edges[i][0]], g.cube[g.edges[i][1]])
	// }
}

//line draw with central projection
func (g *Game) drawLine(screen *ebiten.Image, a, b point) {

	//central projection
	k := float64(400)
	proj(&a, g.cp, k)
	proj(&b, g.cp, k)

	//draw function
	ebitenutil.DrawLine(screen, g.cp.x+a.x, g.cp.y+a.y, g.cp.x+b.x, g.cp.y+b.y, color.White)
}

//draw middle point
func (g *Game) drawMP(screen *ebiten.Image, a, b point) {
	mp := mp(a, b)
	ebitenutil.DrawCircle(screen, mp.x+g.cp.x, mp.y+g.cp.y, 2, color.White)
}

//-------------------------Functions----------------------------------

func normalize(a point) (res point) {
	length := math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)

	res.x = a.x / length
	res.y = a.y / length
	res.z = a.z / length

	return res
}

//from b subtract a
func sub(b, a point) (res point) {
	res.x = b.x - a.x
	res.y = b.y - a.y
	return res
}

//face middle point (takes face diagonal points)
func mp(a, b point) point {
	dx := b.x - a.x
	dy := b.y - a.y
	dz := b.z - a.z

	mp := point{a.x + dx/2, a.y + dy/2, a.z + dz/2}
	return mp
}

func crossProduct(u, v point) (n point) {
	n.x = u.y*v.z - u.z*v.y
	n.y = u.z*v.x - u.x*v.z
	n.z = u.x*v.y - u.y*v.x
	return n
}

func dotProduct(u, v point) float64 {
	return u.x*v.x + u.y*v.y + u.z*v.z
}

//central projection
func proj(p *point, cp point, k float64) {
	//k - scaling koefficient

	//formulas
	z1 := p.z + k
	x1 := (p.x / z1) * k
	y1 := (p.y / z1) * k

	p.x, p.y = x1, y1

}

//rotates the point on given angle on all axis
func (p *point) rotate(r rotator, cp point) {

	//X plane
	p.y = p.y*math.Cos(r.x) + p.z*math.Sin(r.x)
	p.z = -p.y*math.Sin(r.x) + p.z*math.Cos(r.x)

	//Y plane
	p.x = p.x*math.Cos(r.y) - p.z*math.Sin(r.y)
	p.z = p.x*math.Sin(r.y) + p.z*math.Cos(r.y)

	//Z plane
	p.x = p.x*math.Cos(r.z) - p.y*math.Sin(r.z)
	p.y = p.x*math.Sin(r.z) + p.y*math.Cos(r.z)

}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {

	//Window
	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Normal Cube")
	ebiten.SetWindowResizable(true) //enablening window resize

	//Game instance
	g := NewGame(sW, sH)                      //creating game instance
	if err := ebiten.RunGame(g); err != nil { //running game
		log.Fatal(err)
	}
}

//New game instance function
func NewGame(width, height int) *Game {

	//CUBE

	//center point
	cp := point{sW / 2, sH / 2, 100}

	//cube points
	cube := [8]point{
		/*a*/ {-100, -100, -100},
		/*b*/ {100, -100, -100},
		/*c*/ {100, 100, -100},
		/*d*/ {-100, 100, -100},

		/*e*/ {-100, -100, 100},
		/*f*/ {100, -100, 100},
		/*g*/ {100, 100, 100},
		/*h*/ {-100, 100, 100},
	}

	faceEdges := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0},
	}

	faces := [][4]int{
		/*front face*/ {0, 1, 2, 3},
		/*rear face*/ {5, 4, 7, 6},
		/*top face*/ {4, 5, 1, 0},
		/*bottom face*/ {3, 2, 6, 7},
		/*left face*/ {4, 0, 3, 7},
		/*right face*/ {1, 5, 6, 2},
	}

	//rotator
	var r rotator //rotation angle for each axis
	r.x, r.y, r.z = math.Pi/200, math.Pi/200, math.Pi/200

	return &Game{width: width, height: height, r: r, cp: cp, cube: cube, faces: faces, faceEdges: faceEdges}
}
