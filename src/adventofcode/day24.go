package adventofcode

import (
	"fmt"
	"math"
	"strconv"
	//"math/rand"
)

var minBucketSize int
var maxBucketSize int

var validCombos []map[int][]int

func RunDay24(inputfile string) {

	src := readFileAsLines(inputfile)
	fmt.Printf("Input: %v\n", src)

	data := []int{}
	minBucketSize = 0
	total := 0
	//Convert src to []int
	//Determine min/max bucket size:
	for _, v := range src {
		intVal, _ := strconv.Atoi(v)
		data = append(data, intVal)
		if minBucketSize < intVal {
			minBucketSize = intVal
		}
		total += intVal
	}
	maxBucketSize = total / 3
	fmt.Printf("Total: %v, BucketSize: %v\n", total, maxBucketSize)

	combinations := findBalancedCombinations(data)
	fmt.Printf("\nCombinations:\n%v\n", combinations)

	printBestQECombos(combinations)

}

func printBestQECombos(combinations []map[int][]int) {
	smallestQE := 0
	smallestCount := 0
	var smallestOutput string
	for _, c := range combinations {
		l, qe, out := getQuantumEntanglement(c)
		//fmt.Printf("Combo: %v, %v, %v\n", l, qe, out)
		if smallestCount == 0 || smallestCount >= l {
			if smallestCount > l || (smallestQE == 0 || smallestQE > qe) {
				smallestCount = l
				smallestQE = qe
				smallestOutput = out
			}
		}
	}
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nTotal: %v\nBest: %v, %v, %v\n", len(combinations), smallestQE, smallestCount, smallestOutput)

}

func findBalancedCombinations(data []int) []map[int][]int {

	validCombos = []map[int][]int{}
	buckets := make(map[int][]int)
	buckets[0] = []int{}
	buckets[1] = []int{}
	buckets[2] = []int{}
	recursiveFindBalancedCombinations(data, buckets)
	return validCombos
}

func recursiveFindBalancedCombinations(remaining []int, buckets map[int][]int) {

	//fmt.Printf("Remaining(%v): %v\t Buckets: %v, Combos: %v\n", len(remaining), remaining, buckets, validCombos)
	if len(remaining) == 0 {
		//fmt.Println("Checking full buckets...")
		allSame := true

		for _, bucket := range buckets {
			bucketSize := SumIntSlice(bucket)
			//fmt.Printf("Bucket %v(%v) == %v?\n", i, bucketSize, maxBucketSize )

			if bucketSize != maxBucketSize {
				allSame = false
				break
			}
		}
		if allSame {
			smallestLen, qe, out := getQuantumEntanglement(buckets)
			fmt.Printf("\nFound Combo: %v, %v, %v\n", smallestLen, qe, out)

			//Store valid combos in global var, so we don't have to pass it around.
			validCombos = append(validCombos, buckets)

			//Print as we go, so we can try valid guesses before the program finishes
			printBestQECombos(validCombos)

			return
		}

		return //Base case, no more items to store
	}
	next := remaining[0]
	remaining = remaining[1:]

	for i, bucket := range buckets {
		if next+SumIntSlice(bucket) <= maxBucketSize {
			//fmt.Printf("Putting %v in bucket %v\n", next, i)
			newBuckets := copyBuckets(buckets)

			newBucket := make([]int, len(bucket))
			copy(newBucket, bucket)
			newBucket = append(newBucket, next)
			newBuckets[i] = newBucket
			fmt.Printf("Spawning recursive find to bucket %v with %v                  \r", i, newBuckets)

			//Tail-recursion to prevent memory issues
			recursiveFindBalancedCombinations(remaining, newBuckets)
		}
	}
}

func copyBuckets(buckets map[int][]int) map[int][]int {
	newBuckets := make(map[int][]int)
	for k, bucket := range buckets {
		newBucket := make([]int, len(bucket))
		copy(newBucket, bucket)
		newBuckets[k] = newBucket
	}
	return newBuckets
}

func validCombo(buckets map[int][]int) bool {

	bucket0 := SumIntSlice(buckets[0])
	bucket1 := SumIntSlice(buckets[1])
	bucket2 := SumIntSlice(buckets[2])

	//fmt.Printf("Sums %v %v %v for Bucket %v\r", bucket0, bucket1, bucket2, buckets)
	if bucket0 == bucket1 && bucket1 == bucket2 && bucket0 == maxBucketSize {
		return true
	}
	return false
}

func getQuantumEntanglement(buckets map[int][]int) (int, int, string) {

	var output string
	var smallestLength int
	var qe int
	for _, bucket := range buckets {
		product := ProductIntSlice(bucket)
		output = output + fmt.Sprintf("\t%v\t (QE= %v);", bucket, product)
		if smallestLength == 0 || smallestLength > len(bucket) {
			smallestLength = len(bucket)
			qe = product
		}
	}
	output += "\n"
	return smallestLength, qe, output
}

//========================================================================================================
// This function would probably find it, but it is way too slow.

func findBalancedCombinationsExhaustive(data []string) [][]int {
	combinations := [][]int{}
	intData := []int{}
	for _, item := range data {
		val, _ := strconv.ParseInt(item, 10, 64)
		intData = append(intData, int(val))
	}
	fmt.Printf("IntData (%v): %v\n", len(intData), intData)

	//Need a base 3 number, to enumerate combinations of 3 buckets of items.
	//Min size is 3^(len(data)+1)
	//Max size is 3^(len(data)+2)
	//Ignore the last digit; Min size of one extra digit gives me the right length of padded zeros
	minNumber := math.Pow(3, float64(len(data)))
	maxNumber := math.Pow(3, float64(len(data)+1)) - 1
	base3Start := ConvertToBase(int64(minNumber), 3)
	base3End := ConvertToBase(int64(maxNumber), 3)
	fmt.Printf("Max: %v, Base 3S (%v): %v\n", minNumber, len(base3Start), base3Start)
	fmt.Printf("Max: %v, Base 3E: (%v)%v\n", maxNumber, len(base3End), base3End)

	validBuckets := []map[int][]int{}
	for combo := int(maxNumber); combo >= int(minNumber); combo-- {
		base3 := ConvertToBase(int64(combo), 3)
		//fmt.Printf("Base3: %v\n", base3)

		buckets := make(map[int][]int)

		for idx, bucket := range base3 {
			if idx == len(base3)-1 {
				continue
			}
			if _, ok := buckets[bucket]; !ok {
				buckets[bucket] = []int{}
			}
			buckets[bucket] = append(buckets[bucket], intData[idx])
		}
		if validCombo(buckets) {
			//fmt.Printf("\nValid: %v\n", buckets)
			validBuckets = append(validBuckets, buckets)
			l, qe, output := getQuantumEntanglement(buckets)
			fmt.Println("%v\tValid Group: %v, %v, %v\n", combo, l, qe, output)

		}
		//fmt.Printf("Buckets: %v\n", buckets)
	}

	return combinations
}
