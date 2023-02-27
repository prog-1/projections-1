package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Vec struct {
	X, Y, Z float64
}

func Add(a, b Vec) Vec {
	return Vec{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func Sub(a, b Vec) Vec {
	return Vec{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func Divide(v Vec, a float64) Vec {
	return Vec{v.X / a, v.Y / a, v.Z / a}
}

func Mod(a Vec) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func Normalize(v Vec) Vec {
    return Vec{
		v.X / Mod(v),
        v.Y / Mod(v),
        v.Z / Mod(v),
	}
}

type Rotator struct {
	X, Y, Z float64
}

func (v *Vec) RotateZ(rad float64) {
	x := v.X*math.Cos(rad) - v.Y*math.Sin(rad)
	y := v.X*math.Sin(rad) + v.Y*math.Cos(rad)
	v.X, v.Y = x, y
}

func (v *Vec) RotateY(rad float64) {
	x := v.X*math.Cos(rad) + v.Z*math.Sin(rad)
	z := v.X*-math.Sin(rad) + v.Z*math.Cos(rad)
	v.X, v.Z = x, z
}

func (v *Vec) RotateX(rad float64) {
	y := v.Y*math.Cos(rad) - v.Z*math.Sin(rad)
	z := v.Y*math.Sin(rad) + v.Z*math.Cos(rad)
	v.Y, v.Z = y, z
}

func (v *Vec) Rotate(rad Rotator) {
	v.RotateZ(rad.Z)
	v.RotateY(rad.Y)
	v.RotateX(rad.X)
}

func CentralProjection(v Vec, k float64) Vec {
	return Vec{
		-(v.X / v.Z) * k,
		-(v.Y / v.Z) * k,
		-1,
	}
}

func DrawLine(screen *ebiten.Image, a, b Vec, clr color.Color) {
	halfWidth, halfHeight := float64(screenWidth/2), float64(screenHeight/2)
	k := float64(250)
	a = CentralProjection(a, k)
	b = CentralProjection(b, k)
	ebitenutil.DrawLine(screen, a.X+halfWidth, a.Y+halfHeight, b.X+halfWidth, b.Y+halfHeight, color.RGBA{255, 102, 204, 255})
}

type Rect struct {
	A, B, C, D Vec
}

func (r *Rect) Draw(screen *ebiten.Image, clr color.Color) {
	DrawLine(screen, r.A, r.B, clr)
	DrawLine(screen, r.B, r.C, clr)
	DrawLine(screen, r.C, r.D, clr)
	DrawLine(screen, r.D, r.A, clr)
}

type Cube struct {
	p [8]Vec
}

func (c *Cube) Rotate(screen *ebiten.Image, r Rotator) {
	ctr := Add(Divide(Sub(c.p[6], c.p[0]), 2), c.p[0])
	for i := range c.p {
		c.p[i] = Sub(c.p[i], ctr)
		c.p[i].Rotate(r)
		c.p[i] = Add(c.p[i], ctr)
	}
}

func (r *Cube) Draw(screen *ebiten.Image, clr color.Color) {
	// Near plane
  	DrawLine(screen, r.p[0], r.p[1], color.White)
  	DrawLine(screen, r.p[1], r.p[2], color.White)
  	DrawLine(screen, r.p[2], r.p[3], color.White)
  	DrawLine(screen, r.p[3], r.p[0], color.White)

	// Far plane
	DrawLine(screen, r.p[4], r.p[5], clr)
	DrawLine(screen, r.p[5], r.p[6], clr)
	DrawLine(screen, r.p[6], r.p[7], clr)
	DrawLine(screen, r.p[7], r.p[4], clr)

	//Left plane
	DrawLine(screen, r.p[4], r.p[5], clr)
	DrawLine(screen, r.p[5], r.p[1], clr)
	DrawLine(screen, r.p[1], r.p[0], clr)
	DrawLine(screen, r.p[0], r.p[4], clr)

	// Top plane
	DrawLine(screen, r.p[1], r.p[5], clr)
	DrawLine(screen, r.p[5], r.p[6], clr)
	DrawLine(screen, r.p[6], r.p[2], clr)
	DrawLine(screen, r.p[2], r.p[1], clr)

	// Right plane
	DrawLine(screen, r.p[3], r.p[2], clr)
	DrawLine(screen, r.p[2], r.p[6], clr)
	DrawLine(screen, r.p[6], r.p[7], clr)
	DrawLine(screen, r.p[7], r.p[3], clr)
	
	// Bottom plane
	DrawLine(screen, r.p[4], r.p[0], clr)
	DrawLine(screen, r.p[0], r.p[3], clr)
	DrawLine(screen, r.p[3], r.p[7], clr)
	DrawLine(screen, r.p[4], r.p[4], clr)
}

type game struct {
	c            []Cube
	screenBuffer *ebiten.Image
}

func NewGame() *game {
	return &game{
		[]Cube{
			{
				[8]Vec{
					{-200, -200, 200}, // NearBottomLeft
					{-200, 200, 200},  // NearTopLeft
					{200, 200, 200},   // NearTopRight
					{200, -200, 200},  // NearBottomRightS

					{-200, -200, 600}, // FarBottomLeft
					{-200, 200, 600},  // FarTopLeft
					{200, 200, 600},   // FarTopRight
					{200, -200, 600},  // FarBottomRight
				},
			},
		},
		ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	for i := range g.c {
		g.c[i].Rotate(g.screenBuffer, Rotator{math.Pi / 180, math.Pi / 180, math.Pi / 180})
	}
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	for i := range g.c {
		g.c[i].Draw(screen, color.RGBA{255, 102, 204, 255})
	}
	screen.DrawImage(g.screenBuffer, &ebiten.DrawImageOptions{})
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
