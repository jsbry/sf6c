package game

import (
	"fmt"
	"image/color"
	"reflect"
	"strings"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	historyLen = 24
)

func (g *Game) setAction() {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.move = Move{}
	g.attack = Attack{}
	for _, k := range g.keys {
		if k.String() == "A" {
			g.move.Left = true
		}
		if k.String() == "S" {
			g.move.Down = true
		}
		if k.String() == "D" {
			g.move.Right = true
		}
		if g.move.Left && g.move.Down {
			g.move.DownLeft = true
		}
		if g.move.Right && g.move.Down {
			g.move.DownRight = true
		}

		if k.String() == "U" {
			g.attack.Low = true
		}
		if k.String() == "I" {
			g.attack.Middle = true
		}
		if k.String() == "O" {
			g.attack.Hi = true
		}
	}
}

func (g *Game) setHistory() {
	a := &Action{
		m: g.move,
		a: g.attack,
		f: 1,
		c: color.RGBA{0, 0, 0, 255},
	}

	if g.preAction == nil {
		g.actionHistory = append([]*Action{a}, g.actionHistory[0:]...)
		g.preAction = a
	} else if reflect.DeepEqual(g.preAction.m, a.m) && reflect.DeepEqual(g.preAction.a, a.a) {
		if g.preAction.f < 99 {
			g.preAction.f++
		}
	} else {
		g.actionHistory = append([]*Action{a}, g.actionHistory[0:]...)
		if len(g.actionHistory) >= historyLen {
			g.actionHistory = g.actionHistory[:historyLen]
		}
		g.preAction = a
	}
}

func (g *Game) drawHistory(screen *ebiten.Image, perfectTimingKeys []int) {
	for i, h := range g.actionHistory {
		pressedKeys := []int{KeyNeutral}
		keyStrs := []string{"N"}
		if h.m.DownLeft {
			keyStrs = []string{"↙"}
			pressedKeys = []int{KeyDownLeft}
		} else if h.m.DownRight {
			keyStrs = []string{"↘"}
			pressedKeys = []int{KeyDownRight}
		} else if h.m.Left {
			keyStrs = []string{"←"}
			pressedKeys = []int{KeyLeft}
		} else if h.m.Down {
			keyStrs = []string{"↓"}
			pressedKeys = []int{KeyDown}
		} else if h.m.Right {
			keyStrs = []string{"→"}
			pressedKeys = []int{KeyRight}
		}

		if h.a.Low {
			keyStrs = append(keyStrs, "弱")
			pressedKeys = append(pressedKeys, KeyLow)
		}
		if h.a.Middle {
			keyStrs = append(keyStrs, "中")
			pressedKeys = append(pressedKeys, KeyMiddle)
		}
		if h.a.Hi {
			keyStrs = append(keyStrs, "強")
			pressedKeys = append(pressedKeys, KeyHi)
		}
		if h.a.DP {
			keyStrs = append(keyStrs, "DP")
			pressedKeys = append(pressedKeys, KeyDP)
		}
		if h.a.DI {
			keyStrs = append(keyStrs, "DI")
			pressedKeys = append(pressedKeys, KeyDI)
		}
		if h.a.Auto {
			keyStrs = append(keyStrs, "AUTO")
			pressedKeys = append(pressedKeys, KeyAuto)
		}

		if i == 0 && reflect.DeepEqual(perfectTimingKeys, pressedKeys) {
			h.c = color.RGBA{255, 0, 0, 255}
		}
		text.Draw(screen, fmt.Sprintf("%02d  %s", h.f, strings.Join(keyStrs, "+")), bitmapfont.Face, 600, 40+(i*20), h.c)
	}
}
