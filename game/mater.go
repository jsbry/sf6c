package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SquareWidth  = 9
	SquareHeight = 15
)

func (g *Game) frameRateProgress(screen *ebiten.Image) {
	x := 10
	y := ScreenHeight - 30
	// 背景
	for xx := x + SquareWidth - 1; xx <= x+(FrameRate*(SquareWidth+1))+SquareWidth-1; xx++ {
		for yy := y - 1; yy < y+SquareHeight+1; yy++ {
			screen.Set(xx, yy, color.Black)
		}
	}

	x = 10
	// ベース
	for f := 1; f <= FrameRate; f++ {
		square(screen, x+(SquareWidth*f), y, color.RGBA{255, 255, 255, 255})
		x++
	}

	x = 10
	// 経過
	for f := 1; f <= g.fps; f++ {
		square(screen, x+(SquareWidth*f), y, color.RGBA{255, 0, 255, 255})
		x++
	}
}

func square(screen *ebiten.Image, x int, y int, c color.RGBA) {
	for xx := x; xx < x+SquareWidth; xx++ {
		for yy := y; yy < y+SquareHeight; yy++ {
			screen.Set(xx, yy, c)
		}
	}
}
