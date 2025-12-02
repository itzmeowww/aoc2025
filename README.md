# Advent of Code 2025

Solutions implemented in Go!

## Project Structure

```text
aoc2025/
├── 00/          # Template directory
├── 01/          # Day 1 solution
├── 02/          # Day 2 solution
└── ...
```

Each day's directory contains:

- `part1/main.go` - Solution for part 1
- `part2/main.go` - Solution for part 2
- `data/inp.txt` - Puzzle input

## How to Use

### Setup for a New Day

1. Copy the template directory (replace `XX` with the day number, e.g., `03`):

   ```bash
   cp -r 00 XX
   cd XX
   ```

2. Add your puzzle input to `data/inp.txt`

3. Implement your solution in `part1/main.go` and `part2/main.go`

### Running Solutions

From within a day's directory (e.g., `01/`):

```bash
# Run part 1
go run part1/main.go

# Run part 2
go run part2/main.go
```