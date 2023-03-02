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
	angle float64 //point rotation angle
	cp    point   //cube central point
	cube  [8]point
}

type point struct {
	x, y, z float64
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//all logic on update
	for i := range g.cube {
		g.cube[i].rotate(g.angle, g.cp)
	}
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawCube(screen, g.cube)
}

func (g *Game) DrawCube(screen *ebiten.Image, cube [8]point) {

	//1st plane
	drawLine(screen, cube[0], cube[1], g.cp)
	drawLine(screen, cube[1], cube[2], g.cp)
	drawLine(screen, cube[2], cube[3], g.cp)
	drawLine(screen, cube[3], cube[0], g.cp)

	//2nd plane
	drawLine(screen, cube[4], cube[5], g.cp)
	drawLine(screen, cube[5], cube[6], g.cp)
	drawLine(screen, cube[6], cube[7], g.cp)
	drawLine(screen, cube[7], cube[4], g.cp)

	//connectors
	drawLine(screen, cube[0], cube[4], g.cp)
	drawLine(screen, cube[1], cube[5], g.cp)
	drawLine(screen, cube[3], cube[7], g.cp)
	drawLine(screen, cube[2], cube[6], g.cp)

}

func drawLine(screen *ebiten.Image, a, b, cp point) {
	proj(&a, cp, 400)
	proj(&b, cp, 400)
	ebitenutil.DrawLine(screen, a.x, a.y, b.x, b.y, color.White)
}

//central projection
func proj(p *point, cp point, k float64) {
	//k - scaling koefficient

	//moving to top left corner
	p.x = cp.x - p.x
	p.y = cp.y - p.y
	p.z = cp.z - p.z

	//formulas
	z1 := p.z + k
	x1 := (p.x / z1) * k
	y1 := (p.y / z1) * k

	p.x, p.y = x1, y1

	//moving back
	p.x += cp.x
	p.y += cp.y
	p.z += cp.z
}

//-------------------------Functions----------------------------------

//rotates the point on given angle on all axis
func (p *point) rotate(angle float64, cp point) {

	//moving to top left corner
	p.x = cp.x - p.x
	p.y = cp.y - p.y
	p.z = cp.z - p.z

	//X plane
	p.y = p.y*math.Cos(angle) + p.z*math.Sin(angle)
	p.z = -p.y*math.Sin(angle) + p.z*math.Cos(angle)

	//Y plane
	p.x = p.x*math.Cos(angle) - p.z*math.Sin(angle)
	p.z = p.x*math.Sin(angle) + p.z*math.Cos(angle)

	//Z plane
	p.x = p.x*math.Cos(angle) - p.y*math.Sin(angle)
	p.y = p.x*math.Sin(angle) + p.y*math.Cos(angle)

	//moving back
	p.x += cp.x
	p.y += cp.y
	p.z += cp.z

}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {

	//Window
	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Cube")
	ebiten.SetWindowResizable(true) //enablening window resize

	//Game instance
	g := NewGame(sW, sH)                      //creating game instance
	if err := ebiten.RunGame(g); err != nil { //running game
		log.Fatal(err)
	}
}

//New game instance function
func NewGame(width, height int) *Game {

	//Cube
	var cp point //center point
	cp.x, cp.y, cp.z = sW/2, sH/2, 100

	var a, b, c, d, e, f, g, h point

	a.x, a.y, a.z = cp.x-100, cp.y-100, cp.z+100
	b.x, b.y, b.z = cp.x+100, cp.y-100, cp.z+100
	c.x, c.y, c.z = cp.x+100, cp.y+100, cp.z+100
	d.x, d.y, d.z = cp.x-100, cp.y+100, cp.z+100

	e.x, e.y, e.z = cp.x-100, cp.y-100, cp.z-100
	f.x, f.y, f.z = cp.x+100, cp.y-100, cp.z-100
	g.x, g.y, g.z = cp.x+100, cp.y+100, cp.z-100
	h.x, h.y, h.z = cp.x-100, cp.y+100, cp.z-100

	return &Game{width: width, height: height, angle: math.Pi / 720, cp: cp, cube: [8]point{a, b, c, d, e, f, g, h}}
}
