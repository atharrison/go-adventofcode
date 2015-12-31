package adventofcode

import (
	"bufio"
	"io"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFileAsLines(inputfile string) []string {
	f, err := os.Open(inputfile)
	checkError(err)
	reader := bufio.NewReader(f)

	var inputs []string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		inputs = append(inputs, line[0:len(line)-1])
	}
	return inputs
}
