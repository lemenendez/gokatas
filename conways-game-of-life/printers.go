package main

import (
	"fmt"
	"runtime"
	"strings"
)

// Printer renders a LifeGame frame.
type Printer interface {
	Print(LifeGame)
}

// BoardPrinter renders the board in the terminal without clearing the screen,
// which eliminates flicker by overwriting lines in place.
type BoardPrinter struct{}

func (bp BoardPrinter) Print(g LifeGame) {
	var sb strings.Builder

	// Move cursor to top-left without erasing — avoids the flicker caused by
	// a full clear (\033[2J) before every frame.
	sb.WriteString("\033[H")

	for _, row := range g.Get() {
		for _, cell := range row {
			if cell {
				sb.WriteString("⬛")
			} else {
				sb.WriteString("⬜")
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("RAM: %s\n", memUsage()))

	fmt.Print(sb.String())
}

// memUsage returns a human-readable string of current heap memory in use.
func memUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("%d KB", m.Alloc/1024)
}
