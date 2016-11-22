package adventofcode

import (
	"fmt"
	"math"
)

func ExecuteDay20(input int64) {

	fmt.Printf("Checking against %v\n", input)

	var presents int64
	house := 1

	limits := make(map[int]int)

	count := 1.0
	for presents < input {

		//Part 1
		//presents = calculatePresentsForHouse(house)

		//Part 2
		presents = calculatePresentsForHouseWithLimits(house, limits)
		count = count + 1
		if math.Mod(count, 1000.0) == 0 {
			fmt.Printf("House %v received %v presents\n", house, presents)
		}
		house = house + 1
	}
	fmt.Printf("Final: House %v received %v presents\n", house-1, presents)

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

func calculatePresentsForHouseWithLimits(house int, limits map[int]int) int64 {
	elf := house
	var presents int64
	for elf > 0 {
		if limits[elf] < 50 && math.Mod(float64(house), float64(elf)) == 0 {
			presents = presents + int64(11*elf)
			limits[elf] = limits[elf] + 1
		} else if limits[elf] == 50 {
			fmt.Printf("Elf %v retired\n", elf)
			limits[elf] = 51
		}
		elf = elf - 1
	}

	return presents
}
