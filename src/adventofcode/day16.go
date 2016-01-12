package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

/*
children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1
*/

type Aunt struct {
	Number     int64
	Attributes map[string]int64
}

func (a *Aunt) String() string {
	return fmt.Sprintf("[%d] : %v", a.Number, a.Attributes)
}

func (a *Aunt) Matches(properties map[string]int64) bool {
	for k, v := range a.Attributes {
		if properties[k] != v {
			return false
		}
	}
	return true
}

func (a *Aunt) MatchesPart2(properties map[string]int64) bool {
	for k, v := range a.Attributes {
		switch k {
		case "trees", "cats":
			if properties[k] >= v {
				return false
			}
		case "pomeranians", "goldfish":
			if properties[k] <= v {
				return false
			}
		default:
			if properties[k] != v {
				return false
			}
		}
	}
	return true
}

func ExecuteDay16(inputfile string) {

	lines := readFileAsLines(inputfile)

	aunts := make([]*Aunt, len(lines))

	for idx, line := range lines {
		aunts[idx] = ProcessAuntLine(line)
		fmt.Println(aunts[idx])
	}

	properties := map[string]int64{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	// Part 1
	for _, aunt := range aunts {
		if aunt.Matches(properties) {
			fmt.Printf("Part 1 Matching: %v\n", aunt)
		}
	}

	// Part 2
	for _, aunt := range aunts {
		if aunt.MatchesPart2(properties) {
			fmt.Printf("Part 2 Matching: %v\n", aunt)
		}
	}

}

func ProcessAuntLine(line string) *Aunt {

	words := strings.Split(line, ":")
	number, err := strconv.ParseInt(strings.Split(words[0], " ")[1], 10, 64)
	checkError(err)

	attrib1 := strings.TrimSpace(words[1])
	words2 := strings.Split(words[2], ",")
	attrib1val, err := strconv.ParseInt(strings.TrimSpace(words2[0]), 10, 64)
	checkError(err)
	attrib2 := strings.TrimSpace(words2[1])

	words3 := strings.Split(words[3], ",")
	attrib2val, err := strconv.ParseInt(strings.TrimSpace(words3[0]), 10, 64)
	checkError(err)
	attrib3 := strings.TrimSpace(words3[1])

	attrib3val, err := strconv.ParseInt(strings.TrimSpace(words[4]), 10, 64)
	checkError(err)

	attributes := make(map[string]int64)
	attributes[attrib1] = attrib1val
	attributes[attrib2] = attrib2val
	attributes[attrib3] = attrib3val

	return &Aunt{
		Number:     number,
		Attributes: attributes,
	}
}
