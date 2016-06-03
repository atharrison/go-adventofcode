package adventofcode

import (
	"fmt"
	"strings"
)

func ExecuteDay19(inputfile string) {

	lines := readFileAsLines(inputfile)

	replacements := parseReplacements(lines[0 : len(lines)-1])
	fmt.Printf("Replacements: %v\n", replacements)

	molecule := lines[len(lines)-1]
	fmt.Printf("Molecule: %v\n", molecule)

	results := calibrateMolecule(molecule, replacements)
	fmt.Printf("Results: %v\n", results)
	fmt.Printf("Distinct Result count: %v", len(results))
}

type Replacement struct {
	Match   string
	Replace string
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
		//		fmt.Printf("Splitting %v once\n", currentMolecule)
		nextSplit := strings.SplitN(currentMolecule, match, 2)
		if len(nextSplit) != 2 {
			return results
		}

		fmt.Printf("nextSplit: %v\n", nextSplit)
		results = append(results, []string{fmt.Sprintf("%v%v", currentPrefix, nextSplit[0]), nextSplit[1]})
		//		fmt.Printf("Results so far: %v\n", results)
		if len(nextSplit[1]) == 0 {
			break
		}
		currentPrefix = fmt.Sprintf("%v%v%v", currentPrefix, nextSplit[0], match)
		currentMolecule = nextSplit[1]
	}
	return results
}
