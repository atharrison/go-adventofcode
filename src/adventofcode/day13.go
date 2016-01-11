package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
	//	"time"
	"sort"
)

func ExecuteDay13(inputfile string) {
	lines := readFileAsLines(inputfile)

	happiness := ProcessSeatingLines(lines)

	for _, h := range happiness {
		fmt.Println(h)
	}

	happinessByPerson := ProcessHappinessByPerson(happiness)
	fmt.Println(happinessByPerson)

	best, chart := FindBestSeatingArrangement(happinessByPerson)
	fmt.Printf("Best: [%d]\t%v\n", best, chart)
}

type SeatingHappiness struct {
	Name      string
	Neighbor  string
	Happiness int64
}

func NewSeatingHappiness(line string) *SeatingHappiness {
	words := strings.Split(line, " ")

	happy, err := strconv.ParseInt(words[3], 10, 64)
	checkError(err)

	if words[2] == "lose" {
		happy = -happy
	}

	return &SeatingHappiness{
		Name:      words[0],
		Happiness: happy,
		Neighbor:  words[10][:len(words[10])-1],
	}

}

func (s SeatingHappiness) String() string {
	return fmt.Sprintf("%s -> %s = %d", s.Name, s.Neighbor, s.Happiness)
}

func ProcessSeatingLines(lines []string) []*SeatingHappiness {

	happiness := make([]*SeatingHappiness, len(lines))
	for idx, line := range lines {
		happiness[idx] = NewSeatingHappiness(line)
	}
	return happiness
}

func ProcessHappinessByPerson(happiness []*SeatingHappiness) map[string]map[string]int64 {

	personCount := 9 //Part 2 there are 9
	happinessMap := make(map[string]map[string]int64, personCount)

	happinessMap["Andrew"] = make(map[string]int64)

	for _, h := range happiness {
		name, ok := happinessMap[h.Name]
		if !ok {
			name = make(map[string]int64)
			happinessMap[h.Name] = name
			happinessMap["Andrew"][h.Name] = int64(0)
		}
		name[h.Neighbor] = h.Happiness
	}

	return happinessMap
}

func FindBestSeatingArrangement(personMap map[string]map[string]int64) (int64, []string) {

	names := keysFromMap(personMap)
	count := len(names)

	seatingChart := make([]string, count)

	var bestTotal int64
	bestSeatingChart := make([]string, count)

	// This likely would be much cleaner with recursion.
	for a := 0; a < len(names); a++ {
		name, remainder := PluckItemFromArray(names, a)
		seatingChart[0] = name
		for b := 0; b < len(remainder); b++ {
			name, remainder := PluckItemFromArray(remainder, b)
			seatingChart[1] = name
			for c := 0; c < len(remainder); c++ {
				name, remainder := PluckItemFromArray(remainder, c)
				seatingChart[2] = name
				for d := 0; d < len(remainder); d++ {
					name, remainder := PluckItemFromArray(remainder, d)
					seatingChart[3] = name
					for e := 0; e < len(remainder); e++ {
						name, remainder := PluckItemFromArray(remainder, e)
						seatingChart[4] = name
						for f := 0; f < len(remainder); f++ {
							name, remainder := PluckItemFromArray(remainder, f)
							seatingChart[5] = name
							for g := 0; g < len(remainder); g++ {
								name, remainder := PluckItemFromArray(remainder, g)
								seatingChart[6] = name
								for g := 0; g < len(remainder); g++ {
									name, remainder := PluckItemFromArray(remainder, g)
									// Yep, recursion would have made part 2 EASY
									seatingChart[7] = name
									seatingChart[8] = remainder[0]

									total := CalculateHappiness(seatingChart, personMap)
									if bestTotal < total {
										copy(bestSeatingChart, seatingChart)
										bestTotal = total
										fmt.Printf("[%d] New Best with %v\n", total, seatingChart)
									}
									//								time.Sleep(time.Second)

								}
							}
						}
					}
				}
			}
		}
	}

	return bestTotal, bestSeatingChart
}

func CalculateHappiness(seatingChart []string, personMap map[string]map[string]int64) int64 {

	var total int64

	for idx, name := range seatingChart {
		var neighbor string
		if idx == len(seatingChart)-1 {
			neighbor = seatingChart[0]
		} else {
			neighbor = seatingChart[idx+1]
		}
		happiness := personMap[name][neighbor]
		total += happiness
		//		fmt.Printf("[%d]\t%s next to %s adds %d\n", total, name, neighbor, happiness)

		happiness = personMap[neighbor][name]
		total += happiness
		//		fmt.Printf("[%d]\t%s next to %s adds %d\n", total, neighbor, name, happiness)
	}
	fmt.Printf("[%d]\t%v\n", total, seatingChart)

	return total
}

func PluckItemFromArray(items []string, ptr int) (string, []string) {
	item := items[ptr]
	remainder := make([]string, len(items))
	copy(remainder, items)
	remainder = append(remainder[:ptr], remainder[ptr+1:]...)

	//	fmt.Printf("From %v, taking item %d as %s, remainder %v\n", items, ptr, item, remainder)
	return item, remainder
}

func keysFromMap(m map[string]map[string]int64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
