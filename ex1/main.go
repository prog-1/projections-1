package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 999
	screenHeight = 999
)

type Game struct {
	img *ebiten.Image
}

var col = color.RGBA{0xff, 0xff, 0xff, 0xff}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	return nil
}

func carpet(img *ebiten.Image, w, h, mx, my float64, cnt int) {
	w = w / 3
	h = h / 3
	ebitenutil.DrawRect(img, w+mx, h+my, w-1, h-1, col)
	cnt++
	if cnt > 6 {
		return
	} else {
		carpet(img, w, h, mx, my, cnt)
		carpet(img, w, h, mx+w, my, cnt)
		carpet(img, w, h, mx+w*2, my, cnt)
		carpet(img, w, h, mx, my+h, cnt)
		carpet(img, w, h, mx+w*2, my+h, cnt)
		carpet(img, w, h, mx, my+h*2, cnt)
		carpet(img, w, h, mx+w, my+h*2, cnt)
		carpet(img, w, h, mx+w*2, my+h*2, cnt)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.img, nil)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	img := ebiten.NewImage(screenWidth, screenHeight)
	carpet(img, screenWidth, screenHeight, 0, 0, 0)
	if err := ebiten.RunGame(&Game{img}); err != nil {
		log.Fatal(err)
	}
}
