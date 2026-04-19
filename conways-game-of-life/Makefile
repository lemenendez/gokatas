.PHONY: run sim sim-case list help

GO111MODULE ?= off
RUN ?= 0
MAX ?= 5s
TICK ?= 500ms

run:
	GO111MODULE=$(GO111MODULE) go run .

sim:
	GO111MODULE=$(GO111MODULE) go run . -run $(RUN) -max $(MAX) -tick $(TICK)

sim-case: sim

list:
	GO111MODULE=$(GO111MODULE) go run . | grep '^Case '

help:
	@echo "Targets:"
	@echo "  make run                          # run all manual checks"
	@echo "  make sim RUN=3 MAX=2s TICK=250ms # animate one case"
	@echo "  make list                         # list available case ids"