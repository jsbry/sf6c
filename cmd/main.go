package main

import (
	"embed"
	"encoding/json"
	"log"
	"os"
	"regexp"
	"sf6c/game"
	"strconv"
	"strings"
)

//go:embed kd.txt
var keyDisplayLog embed.FS

var InputMap map[string]int

func init() {
	InputMap = map[string]int{
		"n":  game.KeyNeutral,
		"a":  game.KeyLeft,
		"as": game.KeyDownLeft,
		"s":  game.KeyDown,
		"sd": game.KeyDownRight,
		"d":  game.KeyRight,
		"lp": game.KeyLP,
		"mp": game.KeyMP,
		"hp": game.KeyHP,
		"lk": game.KeyLK,
		"mk": game.KeyMK,
		"hk": game.KeyHK,
	}
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("run err: %v", err)
	}
}

type C struct {
	Key   int `json:"k"`
	Frame int `json:"f"`
	Hold  int `json:"h"`
}

func run() error {
	b, err := keyDisplayLog.ReadFile("kd.txt")
	if err != nil {
		log.Fatal(err)
	}

	var combo []C
	var preFrame int
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		ss := strings.Split(line, ",")

		frame := 0
		n := len(ss)
		for i := 0; i < n; i++ {
			if isDigitsOnly(ss[i]) {
				frame, _ = strconv.Atoi(ss[i])
			}
		}

		for i := 0; i < n; i++ {
			if !isDigitsOnly(ss[i]) {
				c := C{
					Key:   InputMap[ss[i]],
					Frame: preFrame,
					Hold:  frame,
				}
				combo = append(combo, c)
			}
		}
		preFrame += frame
	}

	b, err = json.Marshal(combo)
	if err != nil {
		return err
	}
	f, err := os.Create("combo.json")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func isDigitsOnly(s string) bool {
	matched, _ := regexp.MatchString(`^[0-9]+$`, s)
	return matched
}
