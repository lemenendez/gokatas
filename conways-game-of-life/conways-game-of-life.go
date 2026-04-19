package main

import (
	"flag"
	"fmt"
	"time"
)

// LifeGame holds the state of a Conway's Game of Life simulation.
// The board is a fixed square grid with toroidal (wrap-around) edges,
// meaning cells on one side are neighbors of cells on the opposite side.
type LifeGame struct {
	board [][]bool
	next  [][]bool
	size  int
}

// normalize wraps x and y into valid board coordinates using toroidal indexing,
// so out-of-bounds values wrap around to the opposite edge.
func (g *LifeGame) normalize(x, y int) (int, int) {
	return ((x % g.size) + g.size) % g.size, ((y % g.size) + g.size) % g.size
}

// Next advances the board by one generation applying the four Conway rules:
//   - Underpopulation: a live cell with fewer than 2 live neighbors dies.
//   - Survival:        a live cell with 2 or 3 live neighbors lives on.
//   - Overpopulation:  a live cell with more than 3 live neighbors dies.
//   - Reproduction:    a dead cell with exactly 3 live neighbors becomes alive.
func (g *LifeGame) Next() {
	for i, row := range g.board {
		copy(g.next[i], row)
	}

	for i, row := range g.board {
		for j, cell := range row {
			alive := 0
			// visit the neighbors
			for r := i - 1; r <= i+1; r++ {
				for q := j - 1; q <= j+1; q++ {
					rx, ry := g.normalize(r, q)
					// avoid current cell in the counts
					if r == i && q == j {
						continue
					}
					if g.board[rx][ry] {
						alive++
					}
				}
			}
			if cell && alive < 2 {
				// kill because underpopulation
				g.next[i][j] = false
			} else if cell && alive > 3 {
				// kill because overpopulation
				g.next[i][j] = false
			} else if !cell && alive == 3 {
				// reproduction
				g.next[i][j] = true
			}
		}
	}
	for i := range g.size {
		copy(g.board[i], g.next[i])
	}
}

// Get returns the current board state as a 2D slice of booleans,
// where true means the cell is alive.
func (g *LifeGame) Get() [][]bool {
	return g.board
}

// BoardEqual reports whether the current board matches other cell by cell.
func (g *LifeGame) BoardEqual(other [][]bool) bool {
	if len(g.board) != len(other) {
		return false
	}

	for i := range g.board {
		if len(g.board[i]) != len(other[i]) {
			return false
		}
		for j := range g.board[i] {
			if g.board[i][j] != other[i][j] {
				return false
			}
		}
	}

	return true
}

// NewGame creates a LifeGame with a square board of the given size.
// The initial pattern state0 is centered on the board.
func NewGame(state0 [][]bool, size int) *LifeGame {
	init := make([][]bool, size)
	for i := range size {
		init[i] = make([]bool, size)
	}

	// set initial state
	w := size / 2
	h := size / 2
	width := patternWidth(state0)
	x := w - len(state0)/2
	for i := 0; i < len(state0); i++ {
		y := h - width/2
		for j := 0; j < len(state0[i]); j++ {
			init[x][y] = state0[i][j]
			y++
		}
		x++
	}

	next := make([][]bool, size)
	for i := range size {
		next[i] = make([]bool, size)
		copy(next[i], init[i])
	}

	return &LifeGame{
		board: init,
		next:  next,
		size:  size,
	}
}

// Start runs the simulation for up to max duration, advancing one generation
// every cycle and rendering each frame with the given Printer.
func (g *LifeGame) Start(p Printer, max time.Duration, cycle time.Duration) {
	ticker := time.NewTicker(cycle)
	done := make(chan bool)
	p.Print(*g)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				g.Next()
				p.Print(*g)
				_ = t
			}
		}
	}()

	time.Sleep(max)
	ticker.Stop()
	done <- true
}

func main() {
	cases := manualCases()

	runID := flag.Int("run", -1, "run one case by index in simulation mode")
	max := flag.Duration("max", 5*time.Second, "simulation duration")
	tick := flag.Duration("tick", 500*time.Millisecond, "simulation tick")
	flag.Parse()

	if *runID >= 0 {
		if *runID >= len(cases) {
			fmt.Printf("Invalid -run value %d. Available ids: 0..%d\n", *runID, len(cases)-1)
			printCaseIDs(cases)
			return
		}

		tc := cases[*runID]
		fmt.Printf("Running case %d: %s\n", *runID, tc.name)
		g := NewGame(tc.state0, tc.size)
		g.Start(&BoardPrinter{}, *max, *tick)
		return
	}

	runManualCases(cases)

	fmt.Println("Tip: run one animated case with -run <id>, e.g. -run 3")
}
