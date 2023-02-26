package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//Implement a program that draws the [Sierpiński carpet]
//https://en.wikipedia.org/wiki/Sierpi%C5%84ski_carpet

const (
	screenWidth  = 729
	screenHeight = 729
	depth        = 3
)

type rectangle struct {
	x, y, w, h float64
}

type game struct{}

func CarpetStore(screen *ebiten.Image, r rectangle, d int) {
	red := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	a := uint8(rand.Intn(255))
	r.w, r.h = r.w/3, r.h/3
	ebitenutil.DrawRect(screen, r.x+r.w, r.y+r.h, r.w, r.h, color.RGBA{red, g, b, a})
	d++
	if d < depth {
		CarpetStore(screen, rectangle{r.x, r.y, r.w, r.h}, d)
		CarpetStore(screen, rectangle{r.x + r.w, r.y, r.w, r.h}, d)
		CarpetStore(screen, rectangle{r.x + 2*r.w, r.y, r.w, r.h}, d)
		CarpetStore(screen, rectangle{r.x, r.y + r.h, r.w, r.h}, d)
		CarpetStore(screen, rectangle{r.x + 2*r.w, r.y + r.h, r.w, r.h}, d)
		CarpetStore(screen, rectangle{r.x, r.y + 2*r.h, r.w, r.h}, d)
		CarpetStore(screen, rectangle{r.x + r.w, r.y + 2*r.h, r.w, r.h}, d)
		CarpetStore(screen, rectangle{r.x + 2*r.w, r.y + 2*r.h, r.w, r.h}, d)
	}
	return
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{230, 230, 250, 255})
	CarpetStore(screen, rectangle{0, 0, screenWidth, screenHeight}, 0)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
