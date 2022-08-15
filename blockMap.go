package main

import (
	"math"
)

type BlockMap struct {
	MapData
}

// Provides a function to check what block is at x,y. Outside of bounds should return some non-air block.
type MapData interface {
	At(x, y int) Block
}

type Block uint8

const (
	Air Block = iota
	Stone
	Brick
)

type RasterMap struct {
	Blocks [][]Block
}

func (m *RasterMap) At(x, y int) Block {
	if 0 <= y && y < len(m.Blocks) {
		if a := m.Blocks[y]; 0 <= x && x < len(a) {
			return a[x]
		}
	}
	return Stone
}

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (b *BlockMap) Trace(x, y, theta float64) (block Block, distance float64) {
	// Split into integer and float parts
	var (
		dX = math.Cos(theta)
		dY = -math.Sin(theta)
	)

	var (
		pX = int(math.Floor(x))
		pY = int(math.Floor(y))
	)

	const renderDistance = 50
	for {
		if distance > renderDistance {
			return Air, renderDistance
		}
		// If we are in a wall, return
		if block = b.At(pX, pY); block != Air {
			return
		}

		if dX > 0 {
			pX += 1
		}
		if dY > 0 {
			pY += 1
		}
		var nearest float64
		nearest = (float64(pX) - x) / dX
		if nearestY := (float64(pY) - y) / dY; dY != 0 && nearestY < nearest {
			nearest = nearestY
			if dY < 0 {
				pY -= 1
			} else {
				pY += 1
			}
		} else {
			if dX < 0 {
				pX -= 1
			} else {
				pX += 1
			}
		}
		if dX > 0 {
			pX -= 1
		}
		if dY > 0 {
			pY -= 1
		}
		distance = nearest
	}
	return
}
