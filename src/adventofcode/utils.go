package adventofcode

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"github.com/oleiade/lane"
	"github.com/sajari/fuzzy"
	//	"fmt"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFileAsLines(inputfile string) []string {
	f, err := os.Open(inputfile)
	checkError(err)
	reader := bufio.NewReader(f)

	var inputs []string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		inputs = append(inputs, line[0:len(line)-1])
	}
	return inputs
}

func readFileAsString(inputfile string) string {
	contents, err := ioutil.ReadFile(inputfile)
	if err != nil {
		panic(err)
	}
	return string(contents)
}

func Levenshtein(item1 string, item2 string) int {
	return fuzzy.Levenshtein(&item1, &item2)
}

type QueueNode struct {
	Value interface{}
}

//type Queue []*QueueNode

func NewLanePriorityQueue(queueType lane.PQType) *lane.PQueue {
	//	return &Queue{}
	return lane.NewPQueue(queueType)
}

//func (q *Queue) Push(n *QueueNode) {
//	*q = append(*q, n)
//}
//
//func (q *Queue) Pop() (n *QueueNode) {
//	n = (*q)[0]
//	*q = (*q)[1:]
//	return
//}
//
//func (q *Queue) Len() int {
//	return len(*q)
//}

//// NewQueue returns a new queue with the given initial size.
//func NewQueue(size int) *Queue {
//	return &Queue{
//		nodes: make([]*QueueNode, size),
//		size:  size,
//	}
//}
//
//// Queue is a basic FIFO queue based on a circular list that resizes as needed.
//type Queue struct {
//	nodes []*QueueNode
//	size  int
//	head  int
//	tail  int
//	count int
//}
//
//// Push adds a node to the queue.
//func (q *Queue) Push(n *QueueNode) {
//	if q.head == q.tail && q.count > 0 {
//		nodes := make([]*QueueNode, len(q.nodes)+q.size)
//		copy(nodes, q.nodes[q.head:])
//		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
//		q.head = 0
//		q.tail = len(q.nodes)
//		q.nodes = nodes
//	}
//	q.nodes[q.tail] = n
//	q.tail = (q.tail + 1) % len(q.nodes)
//	q.count++
//}
//
//// Pop removes and returns a node from the queue in first to last order.
//func (q *Queue) Pop() *QueueNode {
//	if q.count == 0 {
//		return nil
//	}
//	node := q.nodes[q.head]
//	q.head = (q.head + 1) % len(q.nodes)
//	q.count--
//	return node
//}

type PriorityQueue struct {
	queues     map[int][]*QueueNode
	queueType  QueueType
	length     int64
	priorities *IntSet
}

type QueueType int

const (
	MAXPQ QueueType = iota
	MINPQ
)

func NewPriorityQueue(queueType QueueType) *PriorityQueue {
	return &PriorityQueue{
		queues:     make(map[int][]*QueueNode),
		length:     0,
		queueType:  queueType,
		priorities: NewIntSet(),
	}
}

func (q *PriorityQueue) Push(n *QueueNode, priority int) {
	//	fmt.Printf("Pushing %v at priority %v\n", n, priority)
	q.priorities.Add(priority)

	queue := q.queues[priority]
	//	*q = append(*q, n)
	q.queues[priority] = append(queue, n)
}

func (q *PriorityQueue) Pop() *QueueNode {
	var priority int
	if q.queueType == MAXPQ {
		priority = q.priorities.Max
	} else {
		priority = q.priorities.Min
	}
	//	fmt.Printf("Popping from priority %v\n", priority)
	queue := q.queues[priority]

	//	n = (*q)[0]
	item := queue[0]
	//	*q = (*q)[1:]
	q.queues[priority] = queue[1:]
	if len(q.queues[priority]) == 0 {
		//		fmt.Printf("Removing priority %v\n", priority)
		q.priorities.Remove(priority)
	}
	return item
}

func (q *PriorityQueue) Size() int64 {
	var size int64
	for _, queue := range q.queues {
		size = size + int64(len(queue))
	}
	return size
}

type IntSet struct {
	theSet map[int]bool
	Max    int
	Min    int
}

func NewIntSet() *IntSet {
	return &IntSet{
		theSet: make(map[int]bool),
		Max:    0,
		Min:    -1,
	}
}

func (set *IntSet) Add(i int) bool {
	_, found := set.theSet[i]
	set.theSet[i] = true
	if set.Max < i {
		set.Max = i
	}
	if set.Min > i || set.Min == -1 {
		set.Min = i
	}
	return !found //False if it existed already
}

func (set *IntSet) Get(i int) bool {
	_, found := set.theSet[i]
	return found //True if it existed already
}

func (set *IntSet) Remove(i int) {
	delete(set.theSet, i)
	// Reset min/max

	//	fmt.Printf("Removed %v from %v\n", i, set.theSet)
	if len(set.theSet) == 0 {
		set.Min = -1
		set.Max = 0
	} else {
		for k, _ := range set.theSet {
			if set.Max < k {
				set.Max = k
			}
			if set.Min > k || set.Min == -1 {
				set.Min = k
			}
		}
	}
	//	fmt.Printf("Min/Max: [%v, %v]\n", set.Min, set.Max)
}
