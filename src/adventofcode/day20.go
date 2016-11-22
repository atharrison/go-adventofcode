package adventofcode

import (
	"fmt"
	"math"
)

func ExecuteDay20(input int64) {

	fmt.Printf("Checking against %v\n", input)

	var presents int64
	house := 1
	for presents < input {
		presents = calculatePresentsForHouse(house)
		fmt.Printf("House %v received %v presents\n", house, presents)
		house = house + 1
	}

}

func calculatePresentsForHouse(house int) int64 {
	elf := house
	var presents int64
	for elf > 0 {
		if math.Mod(float64(house), float64(elf)) == 0 {
			presents = presents + int64(10*elf)
		}
		elf = elf - 1
	}

	return presents
}
