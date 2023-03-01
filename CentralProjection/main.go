package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1280
	screenHeight = 720
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

func Multiply(v Vec, a float64) Vec {
	return Vec{v.X * a, v.Y * a, v.Z * a}
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

func Cross(a, b Vec) Vec {
	return Vec{
		a.Y*b.Z - b.Y*a.Z,
		a.Z*b.X - b.Z*a.X,
		a.X*b.Y - b.X*a.Y,
	}
}

func Dot(a, b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
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
	ebitenutil.DrawLine(screen, a.X+halfWidth, -a.Y+halfHeight, b.X+halfWidth, -b.Y+halfHeight, clr)
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

func (c *Cube) Draw(screen *ebiten.Image, clr color.Color) {

	// da, db stands for diagonal a and b(diagonal starting and ending points)
	DrawNormal := func(da, db, v, w Vec, col color.Color) {
		ctr := Add(Divide(Sub(da, db), 2), db)
		DrawLine(screen, ctr, Add(Multiply(Normalize(Cross(v, w)), 200), ctr), col)
	}

	for _, f := range [][]int{
		{0, 1, 2, 3}, // Near
		{7, 6, 5, 4}, // Far
		{4, 5, 1, 0}, // Left
		{1, 5, 6, 2}, // Top
		{3, 2, 6, 7}, // Right
		{4, 0, 3, 7}, // Bottom
	} {
		DrawNormal(c.p[f[2]], c.p[f[0]], Sub(c.p[f[1]], c.p[f[2]]), Sub(c.p[f[2]], c.p[f[3]]), color.RGBA{255, 0, 0, 255})
		if cr := Cross(Sub(c.p[f[1]], c.p[f[2]]), Sub(c.p[f[2]], c.p[f[3]])); Dot(Vec{0, 0, 1}, cr) < 0 {
			DrawLine(screen, c.p[f[0]], c.p[f[1]], clr)
			DrawLine(screen, c.p[f[1]], c.p[f[2]], clr)
			DrawLine(screen, c.p[f[2]], c.p[f[3]], clr)
			DrawLine(screen, c.p[f[3]], c.p[f[0]], clr)
		}
	}
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
					{-200, -200, 400}, // NearBottomLeft
					{-200, 200, 400},  // NearTopLeft
					{200, 200, 400},   // NearTopRight
					{200, -200, 400},  // NearBottomRight

					{-200, -200, 800}, // FarBottomLeft
					{-200, 200, 800},  // FarTopLeft
					{200, 200, 800},   // FarTopRight
					{200, -200, 800},  // FarBottomRight
				},
			},
		},
		ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	for i := range g.c {
		g.c[i].Rotate(g.screenBuffer, Rotator{math.Pi / 180 / 2, math.Pi / 180 / 2, math.Pi / 180 / 2})
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

	for i := range g.c {
		g.c[i].Rotate(g.screenBuffer, Rotator{0, math.Pi, 0})
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
