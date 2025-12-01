# Makefile for Advent of Code 2025

# Run a specific day (e.g., make day1)
day%:
	@cd day_$* && go run .

# Run using the central runner (e.g., make run DAY=1)
run:
	@go run aoc.go $(DAY)

# Create a new day from template (e.g., make new DAY=2)
new:
	@if [ -z "$(DAY)" ]; then echo "Please specify DAY=X"; exit 1; fi
	@if [ -d "day_$(DAY)" ]; then echo "day_$(DAY) already exists!"; exit 1; fi
	@cp -r template day_$(DAY)
	@sed -i '' 's/Day X/Day $(DAY)/g' day_$(DAY)/main.go 2>/dev/null || sed -i 's/Day X/Day $(DAY)/g' day_$(DAY)/main.go
	@echo "Created day_$(DAY)"
	@echo "Add your puzzle input to day_$(DAY)/input.txt"

# Run tests for a specific day
test%:
	@cd day_$* && go test -v

# Run all tests
test-all:
	@go test ./...

# Clean build artifacts
clean:
	@find . -name "*.test" -delete
	@find . -name "*.out" -delete

.PHONY: run new test-all clean
