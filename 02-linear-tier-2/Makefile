.PHONY: bench gen clean

# Generate data files (only needed once, or after clean)
gen:
	go run bench/gen.go

# Run the benchmark (generates data first if missing)
bench: bench/data/medium.txt bench/data/large.txt
	go run bench/bench.go

# Auto-generate data if the files don't exist yet
bench/data/medium.txt bench/data/large.txt:
	$(MAKE) gen

clean:
	rm -f bench/data/medium.txt bench/data/large.txt
