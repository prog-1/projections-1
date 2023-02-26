package main

import (
	"image/color"
	"log"

	"math/rand"

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
	initRect rect   //initial rectangle
	rects    []rect //segments
}

type point struct {
	x, y float64
}

type rect struct {
	p    point //rectangle position
	size point //height and width of rectangle
	clr  color.RGBA
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//all logic on update
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DrawRect(screen, g.initRect.p.x, g.initRect.p.y, g.initRect.size.x, g.initRect.size.y, g.initRect.clr)

	for i := range g.rects {
		ebitenutil.DrawRect(screen, g.rects[i].p.x, g.rects[i].p.y, g.rects[i].size.x, g.rects[i].size.y, g.rects[i].clr)
	}

}

//-------------------------Functions----------------------------------

//function which makes rainbow Sierpiński carpet from triangle
func carpet(initRect rect, depth int) []rect {

	p := initRect.p
	dx, dy := initRect.size.x, initRect.size.y

	var rects []rect

	//segments
	var a rect
	a.p.x, a.p.y = p.x, p.y
	a.size.x, a.size.y = dx/3, dy/3
	a.clr = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

	var b rect
	b.p.x, b.p.y = p.x+dx/3, p.y
	b.size.x, b.size.y = dx/3, dy/3
	b.clr = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

	var c rect
	c.p.x, c.p.y = p.x+(dx/3)*2, p.y
	c.size.x, c.size.y = dx/3, dy/3
	c.clr = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

	var d rect
	d.p.x, d.p.y = p.x, p.y+(dx/3)
	d.size.x, d.size.y = dx/3, dy/3
	d.clr = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

	var e rect
	e.p.x, e.p.y = p.x+dx/3, p.y+(dx/3)
	e.size.x, e.size.y = dx/3, dy/3
	e.clr = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

	var f rect
	f.p.x, f.p.y = p.x+(dx/3)*2, p.y+(dx/3)
	f.size.x, f.size.y = dx/3, dy/3
	f.clr = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

	var g rect
	g.p.x, g.p.y = p.x, p.y+(dx/3)*2
	g.size.x, g.size.y = dx/3, dy/3
	g.clr = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

	var h rect
	h.p.x, h.p.y = p.x+(dx/3), p.y+(dx/3)*2
	h.size.x, h.size.y = dx/3, dy/3
	h.clr = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

	var i rect
	i.p.x, i.p.y = p.x+(dx/3)*2, p.y+(dx/3)*2
	i.size.x, i.size.y = dx/3, dy/3
	i.clr = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

	//appending middle rectangle
	rects = append(rects, e)

	//recursion
	if depth != 0 {
		depth--
		rects = append(rects, carpet(a, depth)...)
		rects = append(rects, carpet(b, depth)...)
		rects = append(rects, carpet(c, depth)...)
		rects = append(rects, carpet(d, depth)...)
		rects = append(rects, carpet(f, depth)...)
		rects = append(rects, carpet(g, depth)...)
		rects = append(rects, carpet(h, depth)...)
		rects = append(rects, carpet(i, depth)...)
	}

	return rects

}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {

	//Window
	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Sierpiński carpet")

	//Game instance
	g := NewGame(sW, sH)                      //creating game instance
	if err := ebiten.RunGame(g); err != nil { //running game
		log.Fatal(err)
	}
}

//New game instance function
func NewGame(width, height int) *Game {

	//initial rectangle
	var initRect rect
	initRect.size.x, initRect.size.y = 400, 400
	initRect.p.x, initRect.p.y = (sW/2)-(initRect.size.x/2), (sH/2)-(initRect.size.y/2)
	initRect.clr = color.RGBA{230, 230, 230, 255}

	rects := carpet(initRect, 4)

	return &Game{width: width, height: height, initRect: initRect, rects: rects}
}
