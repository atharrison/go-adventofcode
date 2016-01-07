package adventofcode

import (
	"bytes"
	"fmt"
)

func ExecuteDay10(input string) {

	// Part1: 40 times
	for times := 0; times < 40; times++ {

		input = LookAndSay(input)
	}

	//	fmt.Println(input)
	fmt.Printf("\nLength: %d\n", len(input))

}

func LookAndSay(input string) string {

	digit := string(input[0])

	var outputBuf bytes.Buffer
	count := 1
	for idx := 1; idx < len(input); idx++ {
		nextDigit := string(input[idx])
		switch nextDigit {
		case digit:
			count++
		default:
			outputBuf.WriteString(fmt.Sprintf("%d%v", count, digit))
			count = 1
			digit = nextDigit
		}
	}
	outputBuf.WriteString(fmt.Sprintf("%d%v", count, digit))

	//	fmt.Printf("%v -> %v\n", input, outputBuf.String())
	return outputBuf.String()
}
