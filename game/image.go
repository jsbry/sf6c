package game

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed resource/arrow.png resource/low.png resource/middle.png resource/hi.png resource/dp.png resource/di.png resource/di.png resource/auto.png
var embedImg embed.FS

var imgArrow, imgLow, imgMiddle, imgHi, imgDP, imgDI, imgAuto *ebiten.Image

func init() {
	initImage(&imgArrow, "resource/arrow.png")
	initImage(&imgLow, "resource/low.png")
	initImage(&imgMiddle, "resource/middle.png")
	initImage(&imgHi, "resource/hi.png")
	initImage(&imgDP, "resource/dp.png")
	initImage(&imgDI, "resource/di.png")
	initImage(&imgAuto, "resource/auto.png")
}

func initImage(i **ebiten.Image, p string) {
	b, err := embedImg.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	*i = ebiten.NewImageFromImage(img)
}
