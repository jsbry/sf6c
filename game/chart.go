package game

import (
	"embed"
	"encoding/json"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed combo/combo.json
var embedChart embed.FS

var chartCombo1 []Chart

const (
	KeyNeutral   = 0
	KeyLeft      = 4
	KeyDownLeft  = 1
	KeyDown      = 2
	KeyDownRight = 3
	KeyRight     = 6

	KeyLP = 11
	KeyMP = 12
	KeyHP = 13
	KeyLK = 14
	KeyMK = 15
	KeyHK = 16

	KeyLeftX      = 0
	KeyDownLeftX  = 50
	KeyDownX      = 100
	KeyDownRightX = 150
	KeyRightX     = 200

	KeyLPX = 250
	KeyMPX = 300
	KeyHPX = 350
	KeyLKX = 400
	KeyMKX = 450
	KeyHKX = 500

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
	initChart(&chartCombo1, "combo/combo.json")
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
				note := &Note{
					Key: c.Key,
					Y:   noteStart - (moveDistance * f),
				}
				notes = append(notes, note)
				if c.Hold > 0 {
					for h := 0; h < c.Hold; h++ {
						note := &Note{
							Key: c.Key,
							Y:   noteStart - (moveDistance * f) - h*moveDistance,
						}
						notes = append(notes, note)
					}
				}
			}
		}
	}

	// 反転
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
		case KeyLP:
			atk.LP = true
			attackLP(screen, float64(KeyLPX), float64(n.Y))
		case KeyMP:
			atk.MP = true
			attackMP(screen, float64(KeyMPX), float64(n.Y))
		case KeyHP:
			atk.HP = true
			attackHP(screen, float64(KeyHPX), float64(n.Y))
		case KeyLK:
			atk.LK = true
			attackLK(screen, float64(KeyLKX), float64(n.Y))
		case KeyMK:
			atk.MK = true
			attackMK(screen, float64(KeyMKX), float64(n.Y))
		case KeyHK:
			atk.HK = true
			attackHK(screen, float64(KeyHKX), float64(n.Y))
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
	attackLP(screen, float64(KeyLPX), float64(perfectTimingY))
	attackMP(screen, float64(KeyMPX), float64(perfectTimingY))
	attackHP(screen, float64(KeyHPX), float64(perfectTimingY))
	attackLK(screen, float64(KeyLKX), float64(perfectTimingY))
	attackMK(screen, float64(KeyMKX), float64(perfectTimingY))
	attackHK(screen, float64(KeyHKX), float64(perfectTimingY))
}
