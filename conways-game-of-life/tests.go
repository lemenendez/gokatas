package main

import "fmt"

type testCase struct {
	name     string
	state0   [][]bool
	size     int
	steps    int
	want     [][]bool // centered on the board — use for non-moving patterns
	wantFull [][]bool // exact full board — use when the pattern shifts position
}

func manualCases() []testCase {
	return []testCase{
		{
			name: "Single live cell dies (underpopulation)",
			state0: [][]bool{
				{true},
			},
			size:  5,
			steps: 1,
			want:  [][]bool{{false}},
		},
		{
			name: "Two adjacent live cells die (1 neighbor each)",
			state0: [][]bool{
				{true, true},
			},
			size:  5,
			steps: 1,
			want: [][]bool{
				{false, false},
			},
		},
		{
			name: "2x2 block stays stable",
			state0: [][]bool{
				{true, true},
				{true, true},
			},
			size:  6,
			steps: 1,
			want: [][]bool{
				{true, true},
				{true, true},
			},
		},
		{
			name: "Blinker step 1: horizontal -> vertical",
			state0: [][]bool{
				{true, true, true},
			},
			size:  7,
			steps: 1,
			want: [][]bool{
				{true},
				{true},
				{true},
			},
		},
		{
			name: "Blinker step 2: returns to horizontal",
			state0: [][]bool{
				{true, true, true},
			},
			size:  7,
			steps: 2,
			want: [][]bool{
				{true, true, true},
			},
		},
		{
			// Glider: moves +1 row +1 col every 4 generations (SE direction).
			// After 4 steps it has the same shape but shifted — so we use wantFull.
			name: "Glider moves one step SE after 4 generations",
			state0: [][]bool{
				{false, true, false},
				{false, false, true},
				{true, true, true},
			},
			size:  10,
			steps: 4,
			// On a 10x10 board the glider starts centered at rows 4-6 cols 4-6.
			// After 4 steps it occupies rows 5-7 cols 5-7.
			wantFull: fullBoard(10, [][2]int{{5, 6}, {6, 7}, {7, 5}, {7, 6}, {7, 7}}),
		},
	}
}

func printCaseIDs(cases []testCase) {
	for i, tc := range cases {
		fmt.Printf("  %d: %s\n", i, tc.name)
	}
}

// fullBoard builds a size×size board with only the listed cells set to true.
func fullBoard(size int, cells [][2]int) [][]bool {
	b := make([][]bool, size)
	for i := range size {
		b[i] = make([]bool, size)
	}
	for _, c := range cells {
		b[c[0]][c[1]] = true
	}
	return b
}

func runManualCases(cases []testCase) {
	for i, tc := range cases {
		fmt.Println("----------------------------------------")
		fmt.Printf("Case %d: %s\n", i, tc.name)
		g := NewGame(tc.state0, tc.size)

		var want [][]bool
		if tc.wantFull != nil {
			want = tc.wantFull
		} else {
			want = centeredBoard(tc.want, tc.size)
		}

		fmt.Println("Initial:")
		printBoard(g.Get())

		for range tc.steps {
			g.Next()
		}

		fmt.Printf("After %d generation(s):\n", tc.steps)
		printBoard(g.Get())

		if g.BoardEqual(want) {
			fmt.Println("Result: PASS")
		} else {
			fmt.Println("Result: FAIL")
			fmt.Println("Expected:")
			printBoard(want)
		}
		fmt.Println()
	}
}

func printBoard(board [][]bool) {
	for _, row := range board {
		for _, cell := range row {
			if cell {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

func centeredBoard(state [][]bool, size int) [][]bool {
	board := make([][]bool, size)
	for i := range size {
		board[i] = make([]bool, size)
	}

	w := size / 2
	h := size / 2
	width := patternWidth(state)
	x := w - len(state)/2
	for i := 0; i < len(state); i++ {
		y := h - width/2
		for j := 0; j < len(state[i]); j++ {
			board[x][y] = state[i][j]
			y++
		}
		x++
	}

	return board
}

func patternWidth(state [][]bool) int {
	max := 0
	for i := range state {
		if len(state[i]) > max {
			max = len(state[i])
		}
	}
	return max
}
