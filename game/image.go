package game

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed resource/c/up.png resource/c/left.png resource/c/right.png resource/c/down.png
//go:embed resource/c/down-left.png resource/c/down-right.png
//go:embed resource/c/low-punch.png resource/c/middle-punch.png resource/c/hi-punch.png
//go:embed resource/c/low-kick.png resource/c/middle-kick.png resource/c/hi-kick.png
var embedImg embed.FS

var (
	imgUp        *ebiten.Image
	imgLeft      *ebiten.Image
	imgRight     *ebiten.Image
	imgDown      *ebiten.Image
	imgDownLeft  *ebiten.Image
	imgDownRight *ebiten.Image
	imgLP        *ebiten.Image
	imgMP        *ebiten.Image
	imgHP        *ebiten.Image
	imgLK        *ebiten.Image
	imgMK        *ebiten.Image
	imgHK        *ebiten.Image
)

func init() {
	initImage(&imgUp, "resource/c/up.png")
	initImage(&imgLeft, "resource/c/left.png")
	initImage(&imgRight, "resource/c/right.png")
	initImage(&imgDown, "resource/c/down.png")
	initImage(&imgDownLeft, "resource/c/down-left.png")
	initImage(&imgDownRight, "resource/c/down-right.png")
	initImage(&imgLP, "resource/c/low-punch.png")
	initImage(&imgMP, "resource/c/middle-punch.png")
	initImage(&imgHP, "resource/c/hi-punch.png")
	initImage(&imgLK, "resource/c/low-kick.png")
	initImage(&imgMK, "resource/c/middle-kick.png")
	initImage(&imgHK, "resource/c/hi-kick.png")
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
