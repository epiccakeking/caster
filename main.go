package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TracerGame struct {
	BlockMap
	x, y, theta float64
	sX, sY      int
}

var blockColors = []color.NRGBA{
	{R: 255, B: 255, A: 255}, // In case Air gets rendered somehow
	{R: 128, G: 128, B: 128, A: 255},
	{R: 255, G: 128, B: 128, A: 255},
}

func (t *TracerGame) Draw(s *ebiten.Image) {
	for screenX := 0; screenX < t.sX; screenX++ {
		b, d := t.Trace(t.x, t.y, t.theta-float64(screenX-t.sX/2)/float64(t.sX))
		c := blockColors[b]
		c.A = uint8(255 / (1 + d/10))
		ebitenutil.DrawRect(s, float64(screenX), float64(t.sY)/2-float64(t.sX)/d/2, 1, float64(t.sX)/d, c)
	}
	ebitenutil.DebugPrint(s, fmt.Sprintf("Theta: %v\nFPS %v", t.theta, ebiten.CurrentFPS()))
}
func (t *TracerGame) Layout(w, h int) (int, int) {
	t.sX = w
	t.sY = h
	return w, h
}
func (t *TracerGame) Update() (err error) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		t.x += math.Cos(t.theta) / 60
		t.y -= math.Sin(t.theta) / 60
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		t.x -= math.Cos(t.theta) / 60
		t.y += math.Sin(t.theta) / 60
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		t.theta += .02
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		t.theta -= .02
	}
	return
}
func main() {
	ebiten.RunGame(&TracerGame{
		BlockMap: BlockMap{
			MapData: &RasterMap{
				Blocks: [][]Block{
					{Air, Brick, Air, Brick},
					{Air, Air, Air, Brick},
					{Air, Air, Air, Brick},
					{Air, Air, Air, Brick},
					{Air, Brick, Air, Brick},
				},
			},
		},
		x: 2, y: 2, theta: 0,
	})
}
