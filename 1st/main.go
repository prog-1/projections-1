package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	x1, y1, x2, y2 float64
}

func recusion(screen *ebiten.Image, x1, y1, x2, y2 float64, depth int) {
	if depth == 0 {
		return
	}
	x2 = (x2 - x1) / 3
	y2 = (y2 - y1) / 3
	ebitenutil.DrawRect(screen, x1+x2, y1+y2, x2, y2, color.White)
	recusion(screen, x1, y1, x1+x2, y1+y2, depth-1)
	recusion(screen, x1+x2, y1, x1+x2*2, y1+y2, depth-1)
	recusion(screen, x1+x2*2, y1, x1+x2*3, y1+y2, depth-1)
	recusion(screen, x1, y1+y2, x1+x2, y1+y2*2, depth-1)
	recusion(screen, x1+x2*2, y1+y2, x1+x2*3, y1+y2*2, depth-1)
	recusion(screen, x1, y1+y2*2, x1+x2, y1+y2*3, depth-1)
	recusion(screen, x1+x2, y1+y2*2, x1+x2*2, y1+y2*3, depth-1)
	recusion(screen, x1+x2*2, y1+y2*2, x1+x2*3, y1+y2*3, depth-1)

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, g.x1, g.y1, g.x2-g.x1, g.y2-g.y1, color.RGBA{255, 0, 0, 255})
	recusion(screen, g.x1, g.y1, g.x2, g.y2, 7)
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{100, 100, 400, 400}); err != nil {
		log.Fatal(err)
	}
}
