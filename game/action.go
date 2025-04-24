package game

import (
	"fmt"
	"image/color"
	"log"
	"reflect"
	"strconv"
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
	if g.gamepadIDs == nil {
		g.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	}
	g.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(g.gamepadIDsBuf[:0])
	for _, id := range g.gamepadIDsBuf {
		log.Printf("gamepad connected: id: %d, SDL ID: %s", id, ebiten.GamepadSDLID(id))
		g.gamepadIDs[id] = struct{}{}
	}
	for id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			log.Printf("gamepad disconnected: id: %d", id)
			delete(g.gamepadIDs, id)
		}
	}
	pressedButtons := map[ebiten.GamepadID][]string{}

	for id := range g.gamepadIDs {
		maxButton := ebiten.GamepadButton(ebiten.GamepadButtonCount(id))
		for b := ebiten.GamepadButton(0); b < maxButton; b++ {
			if ebiten.IsGamepadButtonPressed(id, b) {
				pressedButtons[id] = append(pressedButtons[id], strconv.Itoa(int(b)))
			}

			// Log button events.
			if inpututil.IsGamepadButtonJustPressed(id, b) {
				log.Printf("button pressed: id: %d, button: %d", id, b)
			}
			if inpututil.IsGamepadButtonJustReleased(id, b) {
				log.Printf("button released: id: %d, button: %d", id, b)
			}
		}

		if ebiten.IsStandardGamepadLayoutAvailable(id) {
			for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
				// Log button events.
				if inpututil.IsStandardGamepadButtonJustPressed(id, b) {
					log.Printf("standard button pressed: id: %d, button: %d", id, b)
				}
				if inpututil.IsStandardGamepadButtonJustReleased(id, b) {
					log.Printf("standard button released: id: %d, button: %d", id, b)
				}
			}
		}
	}

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
			g.attack.LP = true
		}
		if k.String() == "I" {
			g.attack.MP = true
		}
		if k.String() == "O" {
			g.attack.HP = true
		}
		if k.String() == "J" {
			g.attack.LK = true
		}
		if k.String() == "K" {
			g.attack.MK = true
		}
		if k.String() == "L" {
			g.attack.HK = true
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

		if h.a.LP {
			keyStrs = append(keyStrs, "弱P")
			pressedKeys = append(pressedKeys, KeyLP)
		}
		if h.a.MP {
			keyStrs = append(keyStrs, "中P")
			pressedKeys = append(pressedKeys, KeyMP)
		}
		if h.a.HP {
			keyStrs = append(keyStrs, "強P")
			pressedKeys = append(pressedKeys, KeyHP)
		}
		if h.a.LK {
			keyStrs = append(keyStrs, "弱K")
			pressedKeys = append(pressedKeys, KeyLP)
		}
		if h.a.MK {
			keyStrs = append(keyStrs, "中K")
			pressedKeys = append(pressedKeys, KeyMP)
		}
		if h.a.HK {
			keyStrs = append(keyStrs, "強K")
			pressedKeys = append(pressedKeys, KeyHP)
		}

		if i == 0 && reflect.DeepEqual(perfectTimingKeys, pressedKeys) {
			h.c = color.RGBA{255, 0, 0, 255}
		}
		text.Draw(screen, fmt.Sprintf("%02d  %s", h.f, strings.Join(keyStrs, "+")), bitmapfont.Face, 600, 40+(i*20), h.c)
	}
}
