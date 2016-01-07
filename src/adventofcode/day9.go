package adventofcode

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var idCounter = -1

func ExecuteDay9(inputfile string) {
	lines := readFileAsLines(inputfile)

	nodeGraph := make(map[string]*Node)
	edges := make([]*Edge, len(lines))
	for idx, line := range lines {
		edge := ParseTownDistancePair(line)
		edges[idx] = edge
		fmt.Printf("%d\t%v\n", idx, edges[idx])

		if node, ok := nodeGraph[edge.town1]; ok {
			node.AddEdge(edge)
		} else {
			node := NewNode(edge.town1, edge)
			nodeGraph[edge.town1] = node
		}

		if node, ok := nodeGraph[edge.town2]; ok {
			node.AddEdge(edge)
		} else {
			node := NewNode(edge.town2, edge)
			nodeGraph[edge.town2] = node
		}
	}

	fmt.Printf("Size of nodeGraph: %d\n", len(nodeGraph))
	for _, v := range nodeGraph {
		fmt.Printf("Node %s\n", v.String())
	}

	timeTable := CreateTimeTable(nodeGraph)
	for row, col := range timeTable {
		fmt.Printf("%d: %v\n", row, col)
	}
	//	fmt.Printf("TimeTable:\n%v", timeTable)

	shortest, path, longest, longestPath := BruteForceWalk(timeTable)
	fmt.Printf("Part 1: Shortest: %v via %v\n", shortest, path)
	fmt.Printf("Part 2: Longest: %v via %v\n", longest, longestPath)

}

func CreateTimeTable(nodeGraph map[string]*Node) [][]int {
	timeTable := make([][]int, len(nodeGraph))
	for i := 0; i < len(nodeGraph); i++ {
		timeTable[i] = make([]int, len(nodeGraph))
	}

	for _, node := range nodeGraph {
		for _, edge := range node.edges {
			var othertown string
			if edge.town1 == node.name {
				othertown = edge.town2
			} else {
				othertown = edge.town1
			}
			town2id := nodeGraph[othertown].id
			//			fmt.Printf("TT %d->%d=%d\n", node.id, town2id, edge.distance)
			timeTable[node.id][town2id] = edge.distance
			timeTable[town2id][node.id] = edge.distance
		}
	}

	return timeTable
}

type Edge struct {
	town1    string `json:"town1"`
	town2    string `json:"town2"`
	distance int    `json:"distance"`
}

func (e *Edge) String() string {
	return fmt.Sprintf("%v<-%v->%v", e.town1, e.distance, e.town2)
}

type Node struct {
	id        int     `json:"id"`
	name      string  `json:"name"`
	edges     []*Edge `json:"edge"`
	edgeCount int     `json:"edgeCount"`
}

func (n *Node) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("{")
	for idx, edge := range n.edges[:] {
		buffer.WriteString(edge.String())
		if idx < len(n.edges)-1 {
			buffer.WriteString(",")
		}
	}
	return fmt.Sprintf("%d{%s: %s}", n.id, n.name, buffer.String())
}

func NewNode(name string, firstEdge *Edge) *Node {
	fmt.Printf("New Node %v with %v\n", name, firstEdge)
	edges := make([]*Edge, 7)
	edges[0] = firstEdge
	idCounter++
	return &Node{
		id:        idCounter,
		name:      name,
		edges:     edges[:],
		edgeCount: 1,
	}
}

func (n *Node) AddEdge(edge *Edge) {
	//	fmt.Printf("Adding %v to %v\n", edge, n)
	n.edges[n.edgeCount] = edge
	n.edgeCount++
}

func ParseTownDistancePair(line string) *Edge {
	split := strings.Split(line, " ")
	dist, _ := strconv.ParseInt(split[4], 10, 0)
	return &Edge{
		town1:    split[0],
		town2:    split[2],
		distance: int(dist),
	}
}

func BruteForceWalk(timeTable [][]int) (int, []int, int, []int) {

	shortest := math.MaxInt64
	longest := 0
	path := make([]int, 8)
	longestPath := make([]int, 8)

	count := len(timeTable)
	for a := 0; a < count; a++ {
		for b := 0; b < count; b++ {
			for c := 0; c < count; c++ {
				for d := 0; d < count; d++ {
					for e := 0; e < count; e++ {
						for f := 0; f < count; f++ {
							for g := 0; g < count; g++ {
								for h := 0; h < count; h++ {
									legSet := make(map[int]bool)
									legSet[a] = true
									legSet[b] = true
									legSet[c] = true
									legSet[d] = true
									legSet[e] = true
									legSet[f] = true
									legSet[g] = true
									legSet[h] = true
									if len(legSet) != count {
										continue
									}

									leg1 := timeTable[a][b]
									leg2 := timeTable[b][c]
									leg3 := timeTable[c][d]
									leg4 := timeTable[d][e]
									leg5 := timeTable[e][f]
									leg6 := timeTable[f][g]
									leg7 := timeTable[g][h]

									if leg1 == 0 ||
										leg2 == 0 ||
										leg3 == 0 ||
										leg4 == 0 ||
										leg5 == 0 ||
										leg6 == 0 ||
										leg7 == 0 {
										continue
									}

									distance := leg1 + leg2 + leg3 + leg4 + leg5 + leg6 + leg7
									if distance < shortest {
										shortest = distance
										path = []int{a, b, c, d, e, f, g}
										fmt.Printf("New shortest: %v via %v\n", shortest, path)
									}
									if distance > longest {
										longest = distance
										longestPath = []int{a, b, c, d, e, f, g}
										fmt.Printf("New longest: %v via %v\n", longest, longestPath)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return shortest, path, longest, longestPath
}
