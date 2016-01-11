package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

type GeneralObject struct {
	Item interface{}
}

func ExecuteDay12Part2(inputfile string) {
	contents := readFileAsString(inputfile)

	// TODO
	fmt.Println(contents)

}

func ExecuteDay12Part1(inputfile string) {
	contents := readFileAsString(inputfile)

	//	fmt.Println(contents)
	sum := int64(0)

	//Such a hack:
	data := strings.Split(contents, "[")
	for _, data2 := range data {
		data3 := strings.Split(data2, "]")
		for _, data4 := range data3 {
			data5 := strings.Split(data4, "{")
			for _, data6 := range data5 {
				data7 := strings.Split(data6, "}")
				for _, data8 := range data7 {
					data9 := strings.Split(data8, ":")
					for _, data10 := range data9 {
						data11 := strings.Split(data10, ",")
						for _, item := range data11 {
							val, err := strconv.ParseInt(string(item), 10, 0)
							if err == nil {
								sum += val
								//								fmt.Printf("%s\tVal: %d\t Sum: %d\n", item, val, sum)
							} else {
								//								fmt.Printf("%s\tNaN\n", item)
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("Sum: %d", sum)
}
