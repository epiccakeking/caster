package main

import "math"

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
		pX     = int(math.Floor(x))
		pY     = int(math.Floor(y))
		blockX = x - float64(pX)
		blockY = y - float64(pY)
		dX     = math.Cos(theta)
		dY     = math.Sin(theta)
	)

	for {
		// If we are in a wall, return
		if block = b.At(pX, pY); block != Air {
			return
		}

		// Calculate the "wall" of the current block that we hit (divide by zero isn't handled because that just leads to infinity which is fine)
		var nearest float64
		var nearestDirection Direction
		if dX < 0 {
			nearest = blockX / -dX
			nearestDirection = Left
		} else {
			nearest = (1 - blockX) / dX
			nearestDirection = Right
		}

		if dY < 0 {
			if n := blockY / -dY; n < nearest {
				nearest = n
				nearestDirection = Up
			}
		} else {
			if n := (1 - blockY) / dY; n < nearest {
				nearest = n
				nearestDirection = Down
			}
		}
		// Update coordinates
		switch nearestDirection {
		case Up:
			pY -= 1
			blockY = 1
			blockX -= dX * nearest
		case Down:
			pY += 1
			blockY = 0
			blockX += dX * nearest
		case Left:
			pX -= 1
			blockX = 1
			blockY -= dY * nearest
		case Right:
			pX += 1
			blockX = 0
			blockY += dY * nearest
		}
		distance += nearest
	}
}
