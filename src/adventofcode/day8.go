package adventofcode

import (
	"bytes"
	"fmt"
)

func ExecuteDay8(inputfile string) {
	lines := readFileAsLines(inputfile)

	literalCount := 0
	charCount := 0
	encodedLines := make([]string, len(lines))
	for idx, line := range lines {
		encodedLines[idx] = EncodeQuotedStringLine(line)
		moreLiteral, moreChar := ProcessQuotedStringLine(line)
		fmt.Printf("%d\tLiterals: %d\tChars: %d\tLine: %v\n", idx, moreLiteral, moreChar, line)
		literalCount += moreLiteral
		charCount += moreChar
	}
	subtractand := literalCount - charCount
	fmt.Printf("Part 1: Literal Count: %d\tChar Count: %d\tTotal Space: %d\n", literalCount, charCount, subtractand)

	// Part 2
	literalCount = 0
	charCount = 0
	for idx, line := range encodedLines {
		moreLiteral, moreChar := ProcessQuotedStringLine(line)
		fmt.Printf("%d\tLiterals: %d\tChars: %d\tLine: %v\n", idx, moreLiteral, moreChar, line)
		literalCount += moreLiteral
		charCount += moreChar
	}
	subtractand = literalCount - charCount
	fmt.Printf("Part 2: Literal Count: %d\tChar Count: %d\tTotal Space: %d\n", literalCount, charCount, subtractand)

}

func EncodeQuotedStringLine(line string) string {
	var buffer bytes.Buffer

	buffer.WriteString("\"")
	for idx := 0; idx < len(line); idx++ {
		nextChar := line[idx]
		switch string(nextChar) {
		case "\\", "\"":
			buffer.WriteString("\\")
			buffer.WriteString(string(nextChar))
		default:
			buffer.WriteString(string(nextChar))
		}
	}
	buffer.WriteString("\"")

	return buffer.String()
}

func ProcessQuotedStringLine(line string) (literals int, chars int) {
	if line[0] != '"' || line[len(line)-1] != '"' {
		panic(fmt.Sprintf("Line does not conform! [%v]", line))
	}

	literals = 2
	chars = 0

	//	lastChar := ""
	for idx := 1; idx < len(line)-1; idx++ {
		nextChar := line[idx]
		switch string(nextChar) {
		case "\\":
			idx++
			nextChar = line[idx]
			switch string(nextChar) {
			case "x":
				idx += 2
				literals += 4
				chars++
			case "\\", "\"":
				literals += 2
				chars++
			}
		default:
			literals++
			chars++
		}
	}

	return
}
