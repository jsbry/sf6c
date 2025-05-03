// Package game sf6c
package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/bitmapfont/v3"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 620

	FrameRate        = 60
	MaxTrainingFrame = 300

	InputKeyX = 590
	InputKeyY = 50
)

type Game struct {
	keys          []ebiten.StandardGamepadButton //[]ebiten.Key
	fps           int
	move          Move
	attack        Attack
	trainingFrame int
	chart         []*Chart
	preAction     *Action
	actionHistory []*Action
	notes         []*Note

	gamepadIDsBuf []ebiten.GamepadID
	gamepadIDs    map[ebiten.GamepadID]struct{}
}

type Move struct {
	Up        bool
	Down      bool
	Left      bool
	Right     bool
	DownLeft  bool
	DownRight bool
}

type Attack struct {
	LP bool
	MP bool
	HP bool
	LK bool
	MK bool
	HK bool
}

type Action struct {
	m Move
	a Attack
	f int
	c color.Color
}

func NewGame() *Game {
	g := &Game{
		fps: 0,
	}
	g.setChart()
	return g
}

func (g *Game) Update() error {
	g.setAction()
	g.setHistory()

	g.fps++
	if FrameRate < g.fps {
		g.fps = 0
	}
	g.trainingFrame++

	for _, n := range g.notes {
		n.Y++
	}
	if g.trainingFrame > MaxTrainingFrame {
		g.trainingFrame = 0
		g.chart = []*Chart{}
		g.setChart()
	}

	for _, n := range g.notes {
		n.Y += moveDistance
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{169, 169, 169, 255})

	text.Draw(screen, fmt.Sprintf("fps:%02d training:%03d time:%s", g.fps, g.trainingFrame, time.Now().Format("15:04:05")), bitmapfont.Face, 435, 14, color.Black)

	g.perfect(screen)
	perfectTimingKeys := g.flow(screen)

	g.drawHistory(screen, perfectTimingKeys)

	g.frameRateProgress(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
