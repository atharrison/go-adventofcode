package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
	//	"encoding/json"
	"errors"
)

func ExecuteDay7(inputfile string) {
	lines := readFileAsLines(inputfile)

	gates := make([]*LogicGate, len(lines))
	for idx, line := range lines {
		gates[idx] = ParseLogicGate(line, idx)
		//		jsonbytes, _ := json.Marshal(gates[idx])
		//		fmt.Printf("Gate %v: %v\n", idx, string(jsonbytes))
	}

	wires := ProcessGates(gates)
	if true {

		for wirename, wire := range wires {
			fmt.Printf("Wire %v:\t%v\tSource: %v\n", wirename, wire.Signal, wire.SourceGate)
		}
	}
	if wires["a"] != nil {
		wire := wires["a"]
		fmt.Printf("\n--->Wire %v:\t%v\tSource: %v\n", wire.Identifier, wire.Signal, wire.SourceGate)
	}
}

func ProcessGates(gates []*LogicGate) map[string]*LogicWire {
	wires := make(map[string]*LogicWire, len(gates)*3)

	//Create all Wires
	for _, gate := range gates {
		for _, name := range WireNamesInGate(gate) {
			if wires[name] == nil {
				//Create Wire reference
				wires[name] = &LogicWire{
					Identifier: name,
					SignalSet:  false,
				}
			}
		}
		wires[gate.OutputWire].SourceGate = gate.Id // Assign Source Gate for Wire
	}

OuterLoop:
	for {
		allSignalsSet := true
		//Process Gates in order:
		for _, gate := range gates {
			if !wires[gate.OutputWire].SignalSet {
				var newSignal uint16
				var err error
				switch gate.Operator {
				case OP_NONE:
					newSignal, err = GetDirectSignal(gate.TopInput, wires)
				case OP_AND:
					newSignal, err = GetAndSignal(gate.TopInput, gate.BottomInput, wires)
				case OP_OR:
					newSignal, err = GetOrSignal(gate.TopInput, gate.BottomInput, wires)
				case OP_LSHIFT:
					newSignal, err = GetLShiftSignal(gate.TopInput, gate.BottomInput.Value, wires)
				case OP_RSHIFT:
					newSignal, err = GetRShiftSignal(gate.TopInput, gate.BottomInput.Value, wires)
				case OP_NOT:
					newSignal, err = GetNotSignal(gate.TopInput, wires)
				}
				if err != nil {
					continue
				}
				wires[gate.OutputWire].Signal = newSignal
				wires[gate.OutputWire].SignalSet = true
				allSignalsSet = false
			}
		}
		if allSignalsSet {
			break OuterLoop
		}
	}

	return wires
}

func GetNotSignal(input *LogicInput, wires map[string]*LogicWire) (uint16, error) {
	signal, err := GetDirectSignal(input, wires)
	return ^signal, err
}

func GetLShiftSignal(input *LogicInput, shift uint16, wires map[string]*LogicWire) (uint16, error) {
	signal, err := GetDirectSignal(input, wires)
	return signal << uint(shift), err
}

func GetRShiftSignal(input *LogicInput, shift uint16, wires map[string]*LogicWire) (uint16, error) {
	signal, err := GetDirectSignal(input, wires)
	return signal >> uint(shift), err
}

func GetOrSignal(top *LogicInput, bottom *LogicInput, wires map[string]*LogicWire) (uint16, error) {
	topSignal, err := GetDirectSignal(top, wires)
	if err != nil {
		return 0, err
	}
	bottomSignal, err := GetDirectSignal(bottom, wires)
	if err != nil {
		return 0, err
	}
	return topSignal | bottomSignal, nil
}

func GetAndSignal(top *LogicInput, bottom *LogicInput, wires map[string]*LogicWire) (uint16, error) {
	topSignal, err := GetDirectSignal(top, wires)
	if err != nil {
		return 0, err
	}
	bottomSignal, err := GetDirectSignal(bottom, wires)
	if err != nil {
		return 0, err
	}
	return topSignal & bottomSignal, nil
}

func GetDirectSignal(input *LogicInput, wires map[string]*LogicWire) (uint16, error) {
	switch input.Type {
	case IN_VALUE:
		return input.Value, nil
	case IN_WIRE:
		if wires[input.Wire].SignalSet {
			return wires[input.Wire].Signal, nil
		} else {
			return 0, errors.New("Unset Wire")
		}
	default:
		panic(fmt.Sprintf("Unexpected type in LogicInput: %v", input.Type))
	}
}

func WireNamesInGate(gate *LogicGate) []string {
	names := make([]string, 0)

	if gate.TopInput.Type == IN_WIRE {
		names = append(names, gate.TopInput.Wire)
	}
	if gate.BottomInput != nil && gate.BottomInput.Type == IN_WIRE {
		names = append(names, gate.BottomInput.Wire)
	}

	names = append(names, gate.OutputWire)

	return names[0:]

}

const OP_NONE = "NOOP"
const OP_AND = "AND"
const OP_OR = "OR"
const OP_NOT = "NOT"
const OP_LSHIFT = "LSHIFT"
const OP_RSHIFT = "RSHIFT"

const IN_VALUE = 0
const IN_GATE = 1
const IN_WIRE = 2

type LogicGate struct {
	Id          int
	Operator    string
	TopInput    *LogicInput
	BottomInput *LogicInput // optional when Operator is unary
	OutputWire  string
}

type LogicInput struct {
	Type  int // IN_VALUE, IN_GATE, IN_WIRE
	Value uint16
	Wire  string
	//	Gate int
}

type LogicWire struct {
	Identifier string
	Signal     uint16
	SourceGate int
	SignalSet  bool
}

func ParseLogicGate(line string, gateid int) *LogicGate {

	var gate *LogicGate
	parts := strings.Split(line, " ")

	if parts[0] == OP_NOT {
		gate = NewNotGate(gateid, parts[1], parts[3])
	} else if len(parts) == 3 {
		if intVal, err := strconv.Atoi(parts[0]); err == nil {
			gate = NewInValueGate(gateid, uint16(intVal), parts[2])
		} else {
			gate = NewInWireGate(gateid, parts[0], parts[2])
		}
	} else if len(parts) == 5 {
		switch parts[1] {
		case OP_AND,
			OP_OR,
			OP_LSHIFT,
			OP_RSHIFT:
			gate = NewBinaryGate(gateid, parts[0], parts[1], parts[2], parts[4])
		default:
			panic(fmt.Sprintf("Could not handle %v\n", line))
		}
	} else {
		panic(fmt.Sprintf("Did not properly handle line %v with parts %v\n", line, parts))
	}

	return gate
}

func NewBinaryGate(gateid int, left string, op string, right string, outwire string) *LogicGate {
	return &LogicGate{
		Id:          gateid,
		TopInput:    NewValWireLogicInput(left),
		BottomInput: NewValWireLogicInput(right),
		Operator:    op,
		OutputWire:  outwire,
	}
}

func NewInWireGate(gateid int, inwire string, outwire string) *LogicGate {
	return &LogicGate{
		Id:         gateid,
		Operator:   OP_NONE,
		TopInput:   NewWireLogicInput(inwire),
		OutputWire: outwire,
	}
}

func NewInValueGate(gateid int, inVal uint16, outwire string) *LogicGate {
	return &LogicGate{
		Id:         gateid,
		Operator:   OP_NONE,
		TopInput:   NewValLogicInput(inVal),
		OutputWire: outwire,
	}
}

func NewNotGate(gateid int, input string, outwire string) *LogicGate {
	return &LogicGate{
		Id:         gateid,
		Operator:   OP_NOT,
		TopInput:   NewValWireLogicInput(input),
		OutputWire: outwire,
	}
}

func NewValWireLogicInput(input string) *LogicInput {
	if intVal, err := strconv.Atoi(input); err == nil {
		return NewValLogicInput(uint16(intVal))
	} else {
		return NewWireLogicInput(input)
	}
}

func NewValLogicInput(input uint16) *LogicInput {
	return &LogicInput{
		Type:  IN_VALUE,
		Value: input,
	}
}

func NewWireLogicInput(input string) *LogicInput {
	return &LogicInput{
		Type: IN_WIRE,
		Wire: input,
	}
}
