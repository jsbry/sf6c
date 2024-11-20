package game

import (
	"embed"
	"encoding/json"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed combo/c1.json
var embedChart embed.FS

var chartCombo1 []Chart

const (
	KeyNeutral   = 0
	KeyLeft      = 4
	KeyDownLeft  = 1
	KeyDown      = 2
	KeyDownRight = 3
	KeyRight     = 6

	KeyLow    = 11
	KeyMiddle = 12
	KeyHi     = 13

	KeyDP   = 21
	KeyDI   = 22
	KeyAuto = 23

	KeyLeftX      = 0
	KeyDownLeftX  = 50
	KeyDownX      = 100
	KeyDownRightX = 150
	KeyRightX     = 200

	KeyLowX    = 250
	KeyMiddleX = 300
	KeyHiX     = 350

	KeyDPX   = 400
	KeyDIX   = 450
	KeyAutoX = 500

	noteStart      = -40
	moveDistance   = 5
	perfectTimingY = 530
)

type Chart struct {
	Key   int `json:"k"`
	Frame int `json:"f"`
	Hold  int `json:"h"`
}

type Note struct {
	Key int
	Y   int
}

func init() {
	initChart(&chartCombo1, "combo/c1.json")
}

func initChart(c *[]Chart, p string) {
	b, err := embedChart.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(b, c)
}

func (g *Game) setChart() {
	notes := []*Note{}

	for f := 0; f <= MaxTrainingFrame; f++ {
		for _, c := range chartCombo1 {
			if f == c.Frame {
				notes = append(notes, &Note{Key: c.Key, Y: noteStart - (moveDistance * f)})
				if c.Hold > 0 {
					for h := 0; h < c.Hold; h++ {
						for d := 0; d < moveDistance; d++ {
							notes = append(notes, &Note{Key: c.Key, Y: noteStart - (moveDistance * f) - h*d})
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(notes)/2; i++ {
		notes[i], notes[len(notes)-i-1] = notes[len(notes)-i-1], notes[i]
	}
	g.notes = notes
}

func (g *Game) flow(screen *ebiten.Image) []int {
	perfectTimingKeys := []int{}
	for _, n := range g.notes {
		m := Move{}
		atk := Attack{}
		switch n.Key {
		case KeyLeft:
			m.Left = true
			arrowLeft(screen, float64(KeyLeftX), float64(n.Y))
		case KeyDownLeft:
			m.DownLeft = true
			arrowDownLeft(screen, float64(KeyDownLeftX), float64(n.Y))
		case KeyDown:
			m.Down = true
			arrowDown(screen, float64(KeyDownX), float64(n.Y))
		case KeyDownRight:
			m.DownRight = true
			arrowDownRight(screen, float64(KeyDownRightX), float64(n.Y))
		case KeyRight:
			m.Right = true
			arrowRight(screen, float64(KeyRightX), float64(n.Y))
		case KeyLow:
			atk.Low = true
			attackLow(screen, float64(KeyLowX), float64(n.Y))
		case KeyMiddle:
			atk.Middle = true
			attackMiddle(screen, float64(KeyMiddleX), float64(n.Y))
		case KeyHi:
			atk.Hi = true
			attackHi(screen, float64(KeyHiX), float64(n.Y))
		case KeyDP:
			atk.DP = true
			attackDP(screen, float64(KeyDPX), float64(n.Y))
		case KeyDI:
			atk.DI = true
			attackDI(screen, float64(KeyDIX), float64(n.Y))
		case KeyAuto:
			atk.DI = true
			attackAuto(screen, float64(KeyAutoX), float64(n.Y))
		default:
		}
		if n.Y == perfectTimingY {
			perfectTimingKeys = append(perfectTimingKeys, n.Key)
		}
	}

	return perfectTimingKeys
}

func (g *Game) perfect(screen *ebiten.Image) {
	arrowLeft(screen, float64(KeyLeftX), float64(perfectTimingY))
	arrowDownLeft(screen, float64(KeyDownLeftX), float64(perfectTimingY))
	arrowDown(screen, float64(KeyDownX), float64(perfectTimingY))
	arrowDownRight(screen, float64(KeyDownRightX), float64(perfectTimingY))
	arrowRight(screen, float64(KeyRightX), float64(perfectTimingY))
	attackLow(screen, float64(KeyLowX), float64(perfectTimingY))
	attackMiddle(screen, float64(KeyMiddleX), float64(perfectTimingY))
	attackHi(screen, float64(KeyHiX), float64(perfectTimingY))
	attackDP(screen, float64(KeyDPX), float64(perfectTimingY))
	attackDI(screen, float64(KeyDIX), float64(perfectTimingY))
	attackAuto(screen, float64(KeyAutoX), float64(perfectTimingY))
}
