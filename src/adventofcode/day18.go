package adventofcode

import (
	"fmt"
	//		"time"
)

func ExecuteDay18(inputfile string) {

	lines := readFileAsLines(inputfile)
	grid := ParseDay18LinesToGrid(lines)

	for times := 0; times < 100; times++ {
		grid = Toggle(grid)
		//		PrintGrid(grid)
		//		fmt.Println("")
		//		time.Sleep(time.Second)
	}

	onCount := CountLightsOn(grid) // From Day 6
	fmt.Printf("Total Lights on after 100 iterations: %d", onCount)

}

func PrintGrid(grid [][]bool) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func ParseDay18LinesToGrid(lines []string) [][]bool {
	grid := make([][]bool, len(lines))

	for x, line := range lines {
		grid[x] = make([]bool, len(lines))
		for y, ch := range line {
			var val bool
			switch ch {
			case '.':
				val = false
			case '#':
				val = true
			default:
				panic(fmt.Sprintf("Unexpected Char: %v", ch))
			}
			grid[x][y] = val
		}
	}

	return grid
}

func Toggle(grid [][]bool) [][]bool {
	newGrid := make([][]bool, len(grid))

	for x := 0; x < len(grid); x++ {
		newGrid[x] = make([]bool, len(grid))
		for y := 0; y < len(grid[x]); y++ {
			newGrid[x][y] = NewLightValueFor(grid, x, y)
		}
	}

	return newGrid
}

func NewLightValueFor(grid [][]bool, x int, y int) bool {

	// Part 2 addition:
	if InGridCorner(x, y, len(grid)) {
		return true
	}

	onNeighbors := 0
	for xoff := -1; xoff < 2; xoff++ {
		for yoff := -1; yoff < 2; yoff++ {
			if xoff == 0 && yoff == 0 {
				continue
			}
			if SafeGetGridValue(grid, x+xoff, y+yoff) {
				onNeighbors++
			}
		}
	}

	return NewLightValue(grid[x][y], onNeighbors)
}

func InGridCorner(x int, y int, size int) bool {
	last := size - 1
	if (x == 0 && y == 0) || (x == 0 && y == last) || (x == last && y == 0) || (x == last && y == last) {
		return true
	}
	return false
}

func NewLightValue(oldValue bool, onNeighbors int) bool {
	if oldValue {
		if onNeighbors == 2 || onNeighbors == 3 {
			return true
		}
	} else {
		if onNeighbors == 3 {
			return true
		}
	}
	return false
}

func SafeGetGridValue(grid [][]bool, x int, y int) bool {

	//Part 2 addition:
	if InGridCorner(x, y, len(grid)) {
		return true
	}

	var result bool
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid) {
		result = false
	} else {
		result = grid[x][y]
	}
	//	fmt.Printf("SafeGetGridValue at %d, %d: %v\n", x, y, result)
	return result
}
