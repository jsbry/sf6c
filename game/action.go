package game

import (
	"fmt"
	"image/color"
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
	g.move = Move{}
	g.attack = Attack{}

	if g.gamepadIDs == nil {
		g.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	}
	g.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(g.gamepadIDsBuf[:0])
	for _, id := range g.gamepadIDsBuf {
		// log.Printf("gamepad connected: id: %d, SDL ID: %s", id, ebiten.GamepadSDLID(id))
		g.gamepadIDs[id] = struct{}{}
	}
	for id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			// log.Printf("gamepad disconnected: id: %d", id)
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
			// if inpututil.IsGamepadButtonJustPressed(id, b) {
			// 	log.Printf("button pressed: id: %d, button: %d", id, b)
			// }
			// if inpututil.IsGamepadButtonJustReleased(id, b) {
			// 	log.Printf("button released: id: %d, button: %d", id, b)
			// }
		}

		if ebiten.IsStandardGamepadLayoutAvailable(id) {
			for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
				// Log button events.
				if inpututil.IsStandardGamepadButtonJustPressed(id, b) {
					g.keys = append(g.keys, b)
					// log.Printf("standard button pressed: id: %d, button: %d", id, b)
				}
				if inpututil.IsStandardGamepadButtonJustReleased(id, b) {
					i := 0
					for _, v := range g.keys {
						if v != b {
							g.keys[i] = v
							i++
						}
					}
					g.keys = g.keys[:i]
					// log.Printf("standard button released: id: %d, button: %d", id, b)
				}
			}
		}
	}

	for _, k := range g.keys {
		// move
		if k == ebiten.StandardGamepadButtonLeftLeft {
			g.move.Left = true
		}
		if k == ebiten.StandardGamepadButtonLeftBottom {
			g.move.Down = true
		}
		if k == ebiten.StandardGamepadButtonLeftRight {
			g.move.Right = true
		}
		if g.move.Left && g.move.Down {
			g.move.DownLeft = true
		}
		if g.move.Right && g.move.Down {
			g.move.DownRight = true
		}

		// punch
		if k == ebiten.StandardGamepadButtonRightLeft {
			g.attack.LP = true
		}
		if k == ebiten.StandardGamepadButtonRightTop {
			g.attack.MP = true
		}
		if k == ebiten.StandardGamepadButtonFrontTopRight {
			g.attack.HP = true
		}

		// kick
		if k == ebiten.StandardGamepadButtonRightBottom {
			g.attack.LK = true
		}
		if k == ebiten.StandardGamepadButtonRightRight {
			g.attack.MK = true
		}
		if k == ebiten.StandardGamepadButtonFrontBottomRight {
			g.attack.HK = true
		}

		// ex
		if k == ebiten.StandardGamepadButtonLeftStick {
			g.attack.LP = true
			g.attack.LK = true
		}
		if k == ebiten.StandardGamepadButtonRightStick {
			g.attack.HP = true
			g.attack.HK = true
		}
		if k == ebiten.StandardGamepadButtonFrontTopLeft {
			g.attack.MP = true
			g.attack.MK = true
		}
		if k == ebiten.StandardGamepadButtonFrontBottomLeft {
			g.attack.LK = true
			g.attack.MK = true
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
