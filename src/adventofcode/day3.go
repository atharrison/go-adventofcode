package adventofcode

import (
	"fmt"
	"io/ioutil"
)

const (
	UP    = "^"
	DOWN  = "v"
	LEFT  = "<"
	RIGHT = ">"
)

func ExecuteDay3(inputfile string) {
	buffer, err := ioutil.ReadFile(inputfile)
	if err != nil {
		panic(err)
	}
	contents := make([]string, len(buffer))
	for idx, ch := range buffer {
		contents[idx] = string(ch)
	}

	fmt.Printf("Contents: %v\n", contents)
	fmt.Printf("Size: %v\n", len(contents))

	//	CalculatePresentDelivery(contents)
	CalculatePresentDeliveryWithRobot(contents)
}

func CalculatePresentDelivery(contents []string) {

	xcoord := 0
	ycoord := 0

	counter := make(map[string]int)

	key := coordsToKey(xcoord, ycoord)
	counter[key] = 1

	for _, movement := range contents {
		//		if idx > 3 {
		//			break
		//		}
		if movement == UP {
			ycoord++
		} else if movement == DOWN {
			ycoord--
		} else if movement == LEFT {
			xcoord--
		} else if movement == RIGHT {
			xcoord++
			//		} else {
			//			panic(fmt.Sprintf("WHAT IS THIS? [%v]", movement))
		}
		key = coordsToKey(xcoord, ycoord)
		counter[key] = counter[key] + 1
	}
	fmt.Printf("Counter: %v\n", counter)
	fmt.Printf("Size of counter: %v\n", len(counter))
}

func coordsToKey(x int, y int) string {
	return fmt.Sprintf("%vx%v", x, y)
}

func CalculatePresentDeliveryWithRobot(contents []string) {

	xcoord := 0
	ycoord := 0
	counter := make(map[string]int)

	rxcoord := 0
	rycoord := 0
	//	rcounter := make(map[string]int)

	key := coordsToKey(xcoord, ycoord)
	rkey := coordsToKey(xcoord, ycoord)
	counter[key] = 1
	//	rcounter[rkey] = 1

	robot := false
	for _, movement := range contents {

		if robot {
			if movement == UP {
				rycoord++
			} else if movement == DOWN {
				rycoord--
			} else if movement == LEFT {
				rxcoord--
			} else if movement == RIGHT {
				rxcoord++
			} else {
				continue
			}
			rkey = coordsToKey(rxcoord, rycoord)
			counter[rkey] = counter[rkey] + 1
			//			fmt.Printf("Robot %v\n", movement)
		} else {
			if movement == UP {
				ycoord++
			} else if movement == DOWN {
				ycoord--
			} else if movement == LEFT {
				xcoord--
			} else if movement == RIGHT {
				xcoord++
				//		} else {
				//			panic(fmt.Sprintf("WHAT IS THIS? [%v]", movement))
			} else {
				continue
			}
			key = coordsToKey(xcoord, ycoord)
			counter[key] = counter[key] + 1
			//			fmt.Printf("Santa %v\n", movement)
		}
		robot = !robot
	}
	//	fmt.Printf("Counter: %v\n", counter)
	fmt.Printf("Size of counter: %v\n", len(counter))
	//	fmt.Printf("RCounter: %v\n", counter)
	fmt.Printf("Total: %v\n", len(counter))
}
