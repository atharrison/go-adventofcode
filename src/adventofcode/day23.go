package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

type TuringMachine struct {
	Registers []int
	Tape      []Instruction
	Cursor    int
}

func NewTuringMachine() *TuringMachine {
	return &TuringMachine{
		Registers: make([]int, 2),
		Tape:      []Instruction{},
		Cursor:    0,
	}
}

type Instruction struct {
	Label    string
	Register int
	Offset   int
}

func NewInstruction(line string) Instruction {
	tokens := strings.Split(line, " ")
	//fmt.Printf("Tokens: %v\n", tokens)
	switch tokens[0] {
	case "hlf", "tpl", "inc":
		return Instruction{Label: tokens[0], Register: RegisterForToken(tokens[1])}
	case "jmp":
		offset, _ := strconv.ParseInt(tokens[1], 10, 64)
		return Instruction{Label: tokens[0], Offset: int(offset)}
	case "jie", "jio":
		offset, _ := strconv.ParseInt(tokens[2], 10, 64)
		instr := Instruction{
			Label:    tokens[0],
			Register: RegisterForToken(tokens[1][0:1]),
			Offset:   int(offset),
		}
		return instr
	}
	fmt.Printf("Could not decipher line [%v]\n", line)
	return Instruction{}
}

var REGISTER_OFFSET = int('a')

func RegisterForToken(token string) int {
	reg := int([]rune(token)[0])
	fmt.Printf("Converting %v to Reg %v\n", token, reg)
	return reg - REGISTER_OFFSET

}

func (tm *TuringMachine) AddInstruction(instr Instruction) {
	tm.Tape = append(tm.Tape, instr)
}

func (tm *TuringMachine) ExecuteTape() {
	fmt.Println("Executing Tape...")
	for tm.Cursor < len(tm.Tape) {
		tm.ProcessNextInstruction()
	}
	fmt.Printf("Complete.\n Registers: %v\n", tm.Registers)
}

func (tm *TuringMachine) ProcessNextInstruction() {
	offset := tm.ProcessInstruction(tm.Tape[tm.Cursor])
	tm.Cursor += offset
}

func (tm *TuringMachine) ProcessInstruction(instr Instruction) int {
	fmt.Printf("Registers: %v, Cursor %v, Instr: %v\n", tm.Registers, tm.Cursor, tm.Tape[tm.Cursor])
	switch instr.Label {
	case "hlf":
		fmt.Printf("Half Register %v\n", instr.Register)
		tm.Registers[instr.Register] = tm.Registers[instr.Register] / 2
		return 1
	case "tpl":
		fmt.Printf("Triple Register %v\n", instr.Register)
		tm.Registers[instr.Register] = tm.Registers[instr.Register] * 3
		return 1
	case "inc":
		fmt.Printf("Incr Register %v\n", instr.Register)
		tm.Registers[instr.Register] = tm.Registers[instr.Register] + 1
		return 1
	case "jmp":
		fmt.Printf("Jump to Instr %v\n", tm.Cursor+instr.Offset)
		return instr.Offset
	case "jie":
		// Jump if EVEN
		if tm.Registers[instr.Register]%2 == 0 {
			fmt.Printf("JumpE to Instr %v\n", tm.Cursor+instr.Offset)
			return instr.Offset
		} else {
			fmt.Printf("JumpE NOOP\n")
			return 1
		}
	case "jio":
		// Jump If ONE
		if tm.Registers[instr.Register] == 1 {
			fmt.Printf("JumpO to Instr %v\n", tm.Cursor+instr.Offset)
			return instr.Offset
		} else {
			fmt.Printf("JumpO NOOP\n")
			return 1
		}
	}
	fmt.Printf("Bug in ProcessInstruction! Can't handle [%v]\n", instr.Label)
	return 999
}

func RunDay23(inputfile string) {
	lines := readFileAsLines(inputfile)

	tm := NewTuringMachine()
	for num, line := range lines {
		//fmt.Printf("Processing line %v...", line)
		instr := NewInstruction(line)
		fmt.Printf("%v\t %v\n", num, instr)
		tm.AddInstruction(instr)
	}

	tm.ExecuteTape()
}
