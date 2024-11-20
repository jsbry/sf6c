package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func arrowDown(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(imgArrow, op)
}

func arrowLeft(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	w := imgArrow.Bounds().Dx()
	op.GeoM.Rotate(90 * math.Pi / 180)
	op.GeoM.Translate(x+float64(w), y)
	screen.DrawImage(imgArrow, op)
}

func arrowRight(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	h := imgArrow.Bounds().Dy()
	op.GeoM.Rotate(270 * math.Pi / 180)
	op.GeoM.Translate(x, y+float64(h))
	screen.DrawImage(imgArrow, op)
}

func arrowDownLeft(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	w, h := imgArrow.Bounds().Dx(), imgArrow.Bounds().Dy()
	op.GeoM.Rotate(45 * math.Pi / 180)
	op.GeoM.Translate(x+float64(w/2), y-float64(h/4))
	screen.DrawImage(imgArrow, op)
}
func arrowDownRight(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	w, h := imgArrow.Bounds().Dx(), imgArrow.Bounds().Dy()
	op.GeoM.Rotate(315 * math.Pi / 180)
	op.GeoM.Translate(x-float64(w/4), y+float64(h/2))
	screen.DrawImage(imgArrow, op)
}
