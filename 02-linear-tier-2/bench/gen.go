//go:build ignore
// +build ignore

package main

// gen.go — generates test data files used by the benchmark.
// Run once: go run bench/gen/gen.go
// Produces:
//   bench/data/medium.txt  (~50 000 integers, ~30% zeros)
//   bench/data/large.txt   (~1 000 000 integers, ~30% zeros)

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func generate(n int, zeroRatio float64, seed int64) string {
	r := rand.New(rand.NewSource(seed))
	parts := make([]string, n)
	for i := range parts {
		if r.Float64() < zeroRatio {
			parts[i] = "0"
		} else {
			parts[i] = strconv.Itoa(r.Intn(9998) + 1) // 1..9998
		}
	}
	return strings.Join(parts, "\n")
}

func write(path string, content string) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		panic(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("wrote %s\n", path)
}

func main() {
	write("data/medium.txt", generate(50_000, 0.30, 42))
	write("data/large.txt", generate(1_000_000, 0.30, 99))
}
