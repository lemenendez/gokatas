//go:build ignore
// +build ignore

// bench.go — compares SC: O(n) vs SC: O(1) approaches on real data files.
//
// HOW TO RUN:
//   1. Generate data (only once):   go run bench/gen/gen.go
//   2. Run benchmark:               go run bench/bench.go
//
// WHAT YOU WILL SEE:
//   For each dataset size, both algorithms run on an identical copy of the data.
//   The table shows:
//     - elapsed wall-clock time
//     - HeapAlloc: live heap bytes at the moment of measurement
//     - TotalAlloc: cumulative bytes ever allocated (shows extra O(n) cost clearly)
//
// WHY TotalAlloc matters here:
//   collectNonZeros    — allocates a brand-new result slice → TotalAlloc grows by ~size of input
//   compactInPlaceNoFill — rewrites the original slice in place → TotalAlloc stays flat

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// ---------- algorithms under test ----------

// SC: O(n) — builds a new slice, leaving the original untouched.
func collectNonZeros(a []int) []int {
	result := make([]int, 0, len(a)) // <-- extra allocation proportional to len(a)
	for i := range a {
		if a[i] != 0 {
			result = append(result, a[i])
		}
	}
	return result
}

// SC: O(1) — rewrites a in place, returns logical length.
func compactInPlaceNoFill(a []int) int {
	write := 0
	for i := range a {
		if a[i] != 0 {
			a[write] = a[i]
			write++
		}
	}
	return write
}

// ---------- helpers ----------

func readInts(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var result []int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("bad value %q: %w", line, err)
		}
		result = append(result, n)
	}
	return result, sc.Err()
}

func memStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}

func printStats(label string, elapsed time.Duration, before, after runtime.MemStats) {
	fmt.Printf("  %-30s  time=%-12s  HeapAlloc=%6d KB  TotalAlloc=%6d KB\n",
		label,
		elapsed.Round(time.Microsecond).String(),
		(after.HeapAlloc-before.HeapAlloc)/1024,
		(after.TotalAlloc-before.TotalAlloc)/1024,
	)
}

// ---------- benchmark runner ----------

func run(label, path string) {
	data, err := readInts(path)
	if err != nil {
		fmt.Printf("[%s] error reading file: %v\n", label, err)
		return
	}
	fmt.Printf("\n=== %s  (%d elements) ===\n", label, len(data))

	// --- collectNonZeros (SC: O(n)) ---
	// Work on a copy so both algos see the same input.
	a := make([]int, len(data))
	copy(a, data)
	runtime.GC()
	before := memStats()
	t0 := time.Now()
	result := collectNonZeros(a)
	elapsed := time.Since(t0)
	after := memStats()
	_ = result
	printStats("collectNonZeros  [SC: O(n)]", elapsed, before, after)

	// --- compactInPlaceNoFill (SC: O(1)) ---
	b := make([]int, len(data))
	copy(b, data)
	runtime.GC()
	before = memStats()
	t0 = time.Now()
	n := compactInPlaceNoFill(b)
	elapsed = time.Since(t0)
	after = memStats()
	_ = b[:n]
	printStats("compactInPlaceNoFill [SC: O(1)]", elapsed, before, after)
}

func main() {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("  Benchmark: SC O(n) vs SC O(1) — removing zeros from an array")
	fmt.Println("  TotalAlloc shows cumulative bytes allocated during each call.")
	fmt.Println("  The O(n) variant allocates a whole extra slice; O(1) does not.")
	fmt.Println("------------------------------------------------------------------")

	run("MEDIUM (50 000)", "data/medium.txt")
	run("LARGE  (1 000 000)", "data/large.txt")

	fmt.Println()
}
