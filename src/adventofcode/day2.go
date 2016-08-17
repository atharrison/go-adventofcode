package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func CalculatePaperAndRibbon(inputfile string) {
	f, err := os.Open(inputfile)
	checkError(err)
	reader := bufio.NewReader(f)

	var inputs [][]int64
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		vals := strings.Split(line, "x")
		//		fmt.Printf("Line vals: %v\n", vals)
		dimensions := make([]int64, 3)
		for idx, item := range vals {
			if idx == 2 {
				item = strings.Split(item, "\n")[0]
			}
			intItem, err := strconv.Atoi(item)
			checkError(err)
			dimensions[idx] = int64(intItem)
		}
		inputs = append(inputs, dimensions)
	}

	//	fmt.Printf("Inputs: %v\n", inputs)
	CalculateWrappingPaper(inputs)
	CalculateRibbon(inputs)
}

func CalculateRibbon(inputs [][]int64) {
	total := int64(0)
	for idx := 0; idx < len(inputs); idx++ {
		items := inputs[idx]
		fmt.Printf("items: %v\n", items)

		cubic := items[0] * items[1] * items[2]

		small, loc := smallest(items)
		remaining, _ := remove(items, loc)
		secondSmallest, _ := smallest(remaining)
		fmt.Printf("Items: %v, smallest: %v, second: %v, remaining: %v\n", items, small, secondSmallest, remaining)

		length := small + small + secondSmallest + secondSmallest
		itemRibbon := length + cubic

		total += itemRibbon
		fmt.Printf("length: %v, bow: %v, running total: %v\n", length, cubic, total)
	}
	fmt.Printf("Total Ribbon: %v\n", total)
}

func CalculateWrappingPaper(inputs [][]int64) {
	var total int64

	for idx := 0; idx < len(inputs); idx++ {
		//	for idx := 0; idx < 3; idx++ {
		item := inputs[idx]
		areas := make([]int64, 3)
		areas[0] = 2 * item[0] * item[1]
		areas[1] = 2 * item[1] * item[2]
		areas[2] = 2 * item[0] * item[2]

		smallestArea, _ := smallest(areas)
		dimensionTotal := areas[0] + areas[1] + areas[2] + smallestArea/2
		total = total + dimensionTotal
		fmt.Printf("Dimensions: %v, Areas: %v, smallest: %v, thisTotal: %v, Running Total: %v\n", item, areas, smallestArea, dimensionTotal, total)
	}

	fmt.Printf("Total square feet: %v\n", total)

}

func smallest(items []int64) (int64, int) {
	result := int64(math.MaxInt64)
	location := -1
	//	fmt.Printf("result start: %v, items: %v\n", result, items)
	for idx := 0; idx < len(items); idx++ {
		//		fmt.Printf("checking %v against %v...\n", items[idx], result)
		if result > int64(items[idx]) {
			result = int64(items[idx])
			location = idx
			//			fmt.Printf("result is now %v\n", result)
		}
	}
	//	fmt.Printf("returning %v\n", result)
	return result, location
}

func remove(a []int64, i int) ([]int64, int64) {
	deleted := a[i]
	a[i] = a[len(a)-1]
	a[len(a)-1] = 0
	a = a[:len(a)-1]

	return a, deleted
}
