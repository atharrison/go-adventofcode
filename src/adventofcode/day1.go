package adventofcode

import (
	"fmt"
	"io/ioutil"
)

func CalculateFloor(inputfile string) {
	contents, err := ioutil.ReadFile(inputfile)
	if err != nil {
		panic(err)
	}
	floor := 0
	//	fmt.Printf("Contents: %v\n", string(contents))
	//	fmt.Printf("Content size: %v\n", len(contents))
	for idx := 0; idx < len(contents); idx++ {
		//		fmt.Print(contents[idx])
		if string(contents[idx]) == "(" {
			//			fmt.Print("U")
			floor++
		} else if string(contents[idx]) == ")" {
			//			fmt.Print("D")
			floor--
		} else {
			//			fmt.Print("?")
		}
	}
	fmt.Printf("Final Floor: %v\n", floor)

	// Part 2
	floor = 0
	for idx := 0; idx < len(contents); idx++ {
		if string(contents[idx]) == "(" {
			floor++
		} else if string(contents[idx]) == ")" {
			floor--
		} else {
			fmt.Print("?")
		}
		if floor == -1 {
			fmt.Printf("\nFirst entered basement at %v\n", idx+1)
			return
		}
	}
}
