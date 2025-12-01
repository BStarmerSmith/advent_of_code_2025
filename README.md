# Advent of Code 2025

## Project Structure

```
advent_of_code_2025/
├── day_1/
│   ├── main.go
│   ├── input.txt
│   └── README.md
├── day_2/
│   ├── main.go
│   └── input.txt
├── ...
├── utils/
│   └── file_reader.go
├── template/
│   ├── main.go
│   └── input.txt
├── go.mod
└── README.md
```

## Running Solutions

### Option 1: Run Individual Days (Recommended)
Each day has its own `main.go` file. To run a specific day:

```bash
# From the project root
cd day_1
go run .

# Or from project root in one command
cd day_1 && go run .
```

**Pros:**
- Simple and straightforward
- Each day is independent
- Easy to test and debug individual solutions
- Can run multiple days in parallel in different terminals
- Natural for Advent of Code workflow

**Cons:**
- Need to navigate to each directory

### Option 2: Alternative - Central Runner
If you prefer a central command, create a `main.go` in the root:

```go
// main.go in project root
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run . <day>")
        return
    }
    
    day := os.Args[1]
    cmd := exec.Command("go", "run", ".")
    cmd.Dir = filepath.Join(".", fmt.Sprintf("day_%s", day))
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}
```

Then run with: `go run . 1`

## Creating New Days

1. Copy the template directory:
   ```bash
   cp -r template day_2
   ```

2. Update the day number in the print statement

3. Add your puzzle input to `input.txt`

4. Implement your solution

## Utility Functions

The `utils` package contains common functions:
- `ReadFileLines(filename string) []string` - Reads file lines into a string slice

Add more utilities as needed for:
- Number parsing
- Grid operations
- Graph algorithms
- etc.

## Tips

1. **Input Files**: Always name them `input.txt` for consistency
2. **Testing**: Consider adding `_test.go` files for complex solutions
3. **Benchmarking**: Use Go's built-in benchmarking for optimization challenges
4. **Part 2**: Structure your code to easily extend for part 2

## Quick Start for New Day

```bash
# From project root
cp -r template day_X
cd day_X
# Paste your input into input.txt
go run .
```
