package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ExecuteDay5(inputfile string) {
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

	//	fmt.Printf("Inputs: %v\n", inputs)
	//	CalculateNaughtyNice(inputs)
	CalculateNaughtyNicePart2(inputs)
}

func CalculateNaughtyNice(inputs []string) {

	niceCount := 0
	for _, item := range inputs {
		nice := NoForbidden(item) && ContainsDoubleLetter(item) && Has3Vowels(item)
		if nice {
			niceCount++
		}
	}
	fmt.Printf("Nice: %v\n", niceCount)
}

func CalculateNaughtyNicePart2(inputs []string) {

	niceCount := 0
	for idx, item := range inputs {
		anypairtwice := AnyPairTwice(item)
		letterrepeat := LetterRepeatSeparateByOne(item)
		nice := anypairtwice && letterrepeat
		fmt.Printf("%v\t%v\t%v\t %v\t (%v\t %v)\n", idx, niceCount, item, nice, anypairtwice, letterrepeat)
		if nice {
			niceCount++
		}
	}
	fmt.Printf("Nice: %v\n", niceCount)
}

func AnyPairTwice(item string) bool {

	//	pairCounter := make(map[string]int)

	for idx := 0; idx < len(item)-1; idx++ {

		pair := string(item[idx]) + string(item[idx+1])
		if strings.Contains(item[idx+2:], pair) {
			return true
		}

		// The following section misses one pair. Not sure why.
		//
		//		pair := fmt.Sprintf("%v%v", string(item[idx]), string(item[idx+1]))
		//		if pairCounter[pair] == 1 {
		//			prevpair := fmt.Sprintf("%v%v", string(item[idx-1]), string(item[idx]))
		//			if prevpair != pair {
		//				return true
		//			}
		//		} else {
		//			pairCounter[pair] = 1
		////			fmt.Printf("%v Found [%v]\n", item, pair)
		//		}
	}

	return false
}

func LetterRepeatSeparateByOne(item string) bool {

	for idx := 0; idx < len(item)-2; idx++ {
		if string(item[idx]) == string(item[idx+2]) {
			return true
		}
	}
	return false
}

func NoForbidden(item string) bool {

	for idx := 0; idx < len(item)-1; idx++ {
		substr := fmt.Sprintf("%v%v", string(item[idx]), string(item[idx+1]))
		switch substr {
		case
			"ab",
			"cd",
			"pq",
			"xy":
			return false
		}
	}

	return true
}

func ContainsDoubleLetter(item string) bool {
	for idx, _ := range item {
		if idx == len(item)-1 {
			return false
		} else if item[idx+1] == item[idx] {
			return true
		}
	}
	return false
}

func Has3Vowels(item string) bool {
	vowelCount := 0
	for _, ch := range item {
		switch ch {
		case
			'a',
			'e',
			'i',
			'o',
			'u':
			vowelCount++

		}
	}
	return vowelCount >= 3
}
