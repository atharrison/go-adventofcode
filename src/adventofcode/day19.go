package adventofcode

import (
	"fmt"
	"strings"
	//	"github.com/oleiade/lane"
)

func ExecuteDay19(inputfile string) {

	lines := readFileAsLines(inputfile)

	replacements := parseReplacements(lines[0 : len(lines)-1])
	fmt.Printf("Replacements: %v\n", replacements)

	molecule := lines[len(lines)-1]
	fmt.Printf("Molecule: %v\n", molecule)

	part1 := false
	part2 := true
	if part1 {
		results := calibrateMolecule(molecule, replacements)
		fmt.Printf("Results: %v\n", results)
		fmt.Printf("Distinct Result count: %v", len(results))
	}
	fmt.Println("----------")
	if part2 {
		//		stepCount := fabricateMolecule("e", molecule, replacements)
		stepCount := deconstructMolecule(molecule, "e", replacements, 0)
		fmt.Printf("Final Step Count: %v\n", stepCount)
	}
}

type Replacement struct {
	Match   string
	Replace string
}

type Fabrication struct {
	currentMolecule string
	stepCount       int
	distance        int
}

func deconstructMolecule(molecule string, endMolecule string, replacementSet []Replacement, stepCount int) int {

	if len(molecule) == len(endMolecule) && molecule == endMolecule {
		fmt.Printf("Found a result, after %v steps\n", stepCount)
		return stepCount
	} else if len(molecule) > 1 {
		best := -1
		for _, replacement := range replacementSet {
			splits := splitsForMoleculeAndMatch(molecule, replacement.Replace)
			if len(splits) > 0 {
				fmt.Printf("Splits for %v against %v: %v\n", molecule, replacement, splits)
				for _, split := range splits {
					prefix := split[0]
					suffix := split[1]
					result := fmt.Sprintf("%s%s%s",
						prefix,
						replacement.Match,
						suffix)
					fmt.Printf("Reduced %v to %v\n", molecule, result)
					deconstructedStepCount := deconstructMolecule(result, endMolecule, replacementSet, stepCount+1)

					if best < 0 || deconstructedStepCount < best {
						best = deconstructedStepCount
					} else {
						fmt.Printf("Too deep: %v\n", deconstructedStepCount)
					}
				}
			}
		}

		fmt.Printf("Returning best of %v\n", best)
		return best
	} else {
		//Degenerate
		fmt.Printf("Failed to find a result on this path.")
		return -1
	}
}

func fabricateMolecule(startMolecule string, target string, replacementSet []Replacement) int {

	checks := 0
	partialFabrications := NewPriorityQueue(MINPQ)
	fabrication := Fabrication{currentMolecule: startMolecule,
		stepCount: 0,
		distance:  len(startMolecule),
	}
	partialFabrications.Push(&QueueNode{Value: fabrication}, fabrication.distance)

	//	found := false
	stepCount := 0
	fmt.Printf("Fabricating %v\n", target)
	best := -1
	//	for !found {
	for partialFabrications.Size() > 0 {
		queueItem := partialFabrications.Pop()
		nextFabrication := queueItem.Value.(Fabrication)
		//		fmt.Printf("Checking for match? %v\n", nextFabrication.currentMolecule)
		//		if len(nextFabrication.currentMolecule) == len(target) &&
		//		   nextFabrication.currentMolecule == target {
		if nextFabrication.distance == 0 {
			stepCount = nextFabrication.stepCount
			fmt.Printf("Found %v\nAfter %v steps.\n", startMolecule, stepCount)
			//			found = true
			if best < 0 || best > stepCount {
				best = stepCount
			}
		} else {
			checks = checks + 1
			//			fmt.Printf("Inserting replacements...\n")
			printed := false
			for _, replacement := range replacementSet {
				splits := splitsForMoleculeAndMatch(nextFabrication.currentMolecule, replacement.Match)
				//				fmt.Printf("Splits for %v against %v: %v\n", nextFabrication.currentMolecule, replacement, splits)
				for _, split := range splits {
					prefix := split[0]
					suffix := split[1]
					result := fmt.Sprintf("%s%s%s",
						prefix,
						replacement.Replace,
						suffix)
					queuedFabrication := Fabrication{currentMolecule: result,
						stepCount: nextFabrication.stepCount + 1,
						distance:  Levenshtein(result, target)}
					if best < 0 || queuedFabrication.stepCount < best {
						partialFabrications.Push(&QueueNode{Value: queuedFabrication}, queuedFabrication.distance)
					}

					if checks%100 == 0 && !printed {
						fmt.Printf("[%v, %v], Queueing %v from [%v|%v|%v] at stepCount %v (distance: %v)\n", checks, partialFabrications.Size(), result, prefix, replacement.Replace, suffix, nextFabrication.stepCount+1, nextFabrication.distance)
						printed = true
					}
				}
			}
		}
	}

	return stepCount
}

func parseReplacements(replacementLines []string) []Replacement {

	replacements := make([]Replacement, 0)
	for _, line := range replacementLines {
		split := strings.Split(line, " => ")
		if len(split) == 2 {
			fmt.Printf("Processing Replacement %v into [%v] ~> [%v]\n", line, split[0], split[1])
			replacements = append(replacements, Replacement{Match: split[0], Replace: split[1]})
		}
	}
	return replacements
}

func calibrateMolecule(molecule string, replacements []Replacement) map[string]int {

	results := make(map[string]int)

	for _, r := range replacements {
		//		splits := strings.Split(molecule, r.Match)
		splits := splitsForMoleculeAndMatch(molecule, r.Match)
		fmt.Printf("Split %v into %v\n", molecule, splits)
		for idx, split := range splits {
			prefix := split[0]
			suffix := split[1]
			result := fmt.Sprintf("%s%s%s",
				prefix,
				r.Replace,
				suffix)
			fmt.Printf("Created %v from [%v|%v|%v], at idx %v\n", result, prefix, r.Replace, suffix, idx)
			results[result] = results[result] + 1
		}

		//		result := strings.Replace(molecule, r.Match, r.Replace, 1)
		//		results[result] = results[result] + 1
	}

	return results
}

func splitsForMoleculeAndMatch(molecule string, match string) [][]string {
	var results [][]string
	currentPrefix := ""
	currentMolecule := molecule
	for {
		//				fmt.Printf("Splitting %v once\n", currentMolecule)
		nextSplit := strings.SplitN(currentMolecule, match, 2)
		if len(nextSplit) != 2 {
			return results
		}

		//		fmt.Printf("nextSplit: %v\n", nextSplit)
		results = append(results, []string{fmt.Sprintf("%v%v", currentPrefix, nextSplit[0]), nextSplit[1]})
		//				fmt.Printf("Results so far: %v\n", results)
		if len(nextSplit[1]) == 0 {
			break
		}
		currentPrefix = fmt.Sprintf("%v%v%v", currentPrefix, nextSplit[0], match)
		currentMolecule = nextSplit[1]
	}
	return results
}
