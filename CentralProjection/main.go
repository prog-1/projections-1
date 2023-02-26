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

type Rotator struct {
	X, Y, Z float64
}

func (v *Vec) Rotate(rad Rotator) {
	// Rotation around Z
	v.X = v.X*math.Cos(rad.Z) - v.Y*math.Sin(rad.Z)
	v.Y = v.X*math.Sin(rad.Z) + v.Y*math.Cos(rad.Z)

	// Rotation around Y
	v.X = v.X*math.Cos(rad.Y) - v.Z*math.Sin(rad.Y)
	v.Z = v.X*math.Sin(rad.Y) + v.Z*math.Cos(rad.Y)

	// Rotation around X
	v.Y = v.Y*math.Cos(rad.X) - v.Z*math.Sin(rad.X)
	v.Z = -v.Y*math.Sin(rad.X) + v.Z*math.Cos(rad.X)
}

func CentralProjection(v Vec, k float64) Vec {
	return Vec{
		-(v.X / v.Z) * k,
		-(v.Y / v.Z) * k,
		-k,
	}
}

func DrawLine(screen *ebiten.Image, a, b Vec, clr color.Color) {
	k := float64(-250)
	a = CentralProjection(a, k)
	b = CentralProjection(b, k)
	ebitenutil.DrawLine(screen, a.X, a.Y, b.X, b.Y, color.RGBA{255, 102, 204, 255})
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
	// Drawing near plane
	DrawLine(screen, r.p[0], r.p[1], clr)
	DrawLine(screen, r.p[1], r.p[2], clr)
	DrawLine(screen, r.p[2], r.p[3], clr)
	DrawLine(screen, r.p[3], r.p[0], clr)

	// Drawing far plane
	DrawLine(screen, r.p[4], r.p[5], clr)
	DrawLine(screen, r.p[5], r.p[6], clr)
	DrawLine(screen, r.p[6], r.p[7], clr)
	DrawLine(screen, r.p[7], r.p[4], clr)

	// Drawing connections between planes
	DrawLine(screen, r.p[0], r.p[4], clr)
	DrawLine(screen, r.p[1], r.p[5], clr)
	DrawLine(screen, r.p[2], r.p[6], clr)
	DrawLine(screen, r.p[3], r.p[7], clr)
}

type game struct {
	c            []Cube
	screenBuffer *ebiten.Image
}

func NewGame() *game {
	halfWidth, halfHeight := float64(screenWidth/2), float64(screenHeight/2)
	return &game{
		[]Cube{
			{
				[8]Vec{
					{-200 + halfWidth, -200 + halfHeight, 200}, // NearBottomLeft
					{-200 + halfWidth, 200 + halfHeight, 200},  // NearTopLeft
					{200 + halfWidth, 200 + halfHeight, 200},   // NearTopRight
					{200 + halfWidth, -200 + halfHeight, 200},  // NearBottomRightS

					{-200 + halfWidth, -200 + halfHeight, 250}, // FarBottomLeft
					{-200 + halfWidth, 200 + halfHeight, 250},  // FarTopLeft
					{200 + halfWidth, 200 + halfHeight, 250},   // FarTopRight
					{200 + halfWidth, -200 + halfHeight, 250},  // FarBottomRight
				},
			},
			// {
			// 	[8]Vec{
			// 		{-400 + halfWidth, -200 + halfHeight, 200}, // NearBottomLeft
			// 		{-400 + halfWidth, 200 + halfHeight, 200},  // NearTopLeft
			// 		{halfWidth, 200 + halfHeight, 200},         // NearTopRight
			// 		{halfWidth, -200 + halfHeight, 200},        // NearBottomRight

			// 		{-400 + halfWidth, -200 + halfHeight, 600}, // FarBottomLeft
			// 		{-400 + halfWidth, 200 + halfHeight, 600},  // FarTopLeft
			// 		{halfWidth, 200 + halfHeight, 600},         // FarTopRight
			// 		{halfWidth, -200 + halfHeight, 600},        // FarBottomRight
			// 	},
			// },
		},
		ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	for i := range g.c {
		g.c[i].Rotate(g.screenBuffer, Rotator{0, 0, math.Pi / 180})
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
