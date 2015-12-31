package adventofcode

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func ExecuteDay6(inputfile string) {
	lines := readFileAsLines(inputfile)

	CalculateLightsOn(lines)
	CalculateBrightness(lines)
}

func CalculateLightsOn(lines []string) {
	instructions := make([]*Instruction, len(lines))
	for idx, line := range lines {
		instructions[idx] = ParseInstruction(line)
	}

	grid := make([][]bool, 1000)
	ProcessLightInstructions(instructions, grid)
	totalOn := CountLightsOn(grid)
	fmt.Printf("Total lights on: %v\n", totalOn)
}

func CalculateBrightness(lines []string) {
	instructions := make([]*Instruction, len(lines))
	for idx, line := range lines {
		instructions[idx] = ParseInstruction(line)
	}

	grid := make([][]int, 1000)
	ProcessLightBrightnessInstructions(instructions, grid)
	totalOn := CountBrightnessOf(grid)
	fmt.Printf("Total brightness: %v\n", totalOn)
}

const FUNC_OFF = 0
const FUNC_ON = 1
const FUNC_TOGGLE = 2

const FOUR_NUMBERS_REGEX = "[a-zA-Z ]+ (?P<x1>[0-9]{1,3}),(?P<y1>[0-9]{1,3}) [a-zA-Z ]+ (?P<x2>[0-9]{1,3}),(?P<y2>[0-9]{1,3})"

type Instruction struct {
	Function int
	X1       int
	Y1       int
	X2       int
	Y2       int
}

func ParseInstruction(line string) *Instruction {

	instr := &Instruction{}
	if strings.Contains(line, "turn on") {
		instr.Function = FUNC_ON
	} else if strings.Contains(line, "turn off") {
		instr.Function = FUNC_OFF
	} else if strings.Contains(line, "toggle") {
		instr.Function = FUNC_TOGGLE
	}

	//	matchResults := map[string]string{}
	matcher := regexp.MustCompile(FOUR_NUMBERS_REGEX)
	names := matcher.SubexpNames()
	results := matcher.FindAllStringSubmatch(line, -1)[0]

	md := map[string]string{}
	for i, n := range results {
		md[names[i]] = n
	}
	instr.X1, _ = strconv.Atoi(md["x1"])
	instr.Y1, _ = strconv.Atoi(md["y1"])
	instr.X2, _ = strconv.Atoi(md["x2"])
	instr.Y2, _ = strconv.Atoi(md["y2"])
	return instr
}

func ProcessLightInstructions(instructions []*Instruction, grid [][]bool) {

	//finish initializing grid:
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, 1000)
	}

	//Process each instruction
	for _, instr := range instructions {
		ApplyInstructionToGrid(instr, grid)
	}
}

func ProcessLightBrightnessInstructions(instructions []*Instruction, grid [][]int) {
	//finish initializing grid:
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 1000)
	}

	//Process each instruction
	for _, instr := range instructions {
		ApplyBrightnessInstructionToGrid(instr, grid)
	}
}

func ApplyInstructionToGrid(instr *Instruction, grid [][]bool) {

	for x := instr.X1; x <= instr.X2; x++ {
		for y := instr.Y1; y <= instr.Y2; y++ {
			switch instr.Function {
			case FUNC_OFF:
				grid[x][y] = false
			case FUNC_ON:
				grid[x][y] = true
			case FUNC_TOGGLE:
				grid[x][y] = !grid[x][y]
			}
		}
	}
}

func ApplyBrightnessInstructionToGrid(instr *Instruction, grid [][]int) {

	for x := instr.X1; x <= instr.X2; x++ {
		for y := instr.Y1; y <= instr.Y2; y++ {
			switch instr.Function {
			case FUNC_OFF:
				grid[x][y] = int(math.Max(float64(grid[x][y]-1), 0.0))
			case FUNC_ON:
				grid[x][y]++
			case FUNC_TOGGLE:
				grid[x][y] += 2
			}
		}
	}

}

func CountLightsOn(grid [][]bool) int {
	total := 0
	for _, rows := range grid {
		//		fmt.Printf("Rows: %v\n", rows)
		for _, pixel := range rows {
			if pixel {
				total++
			}
		}
	}
	return total
}

func CountBrightnessOf(grid [][]int) int {
	total := 0
	for _, rows := range grid {
		//		fmt.Printf("Rows: %v\n", rows)
		for _, pixel := range rows {
			total += pixel
		}
	}
	return total
}
