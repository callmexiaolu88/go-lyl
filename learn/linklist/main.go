package main

import (
	"fmt"
	"math/rand"
	"time"
)

const count = 1000

var midindex = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(count - 100)
var sets = []int{2, 3, 4, 5, 6, 7, 8, 9, 8824, 5650, 9860}

type Node struct {
	data int
	next *Node
}

func main() {
	fmt.Println("-----start-----")
	var head Node
	var mid, tmp *Node
	tmp = &head
	for i := 0; i < count; i++ {
		tmp = createLinklist(tmp)
		if i == midindex {
			mid = tmp
		}
	}
	tmp.next = mid
	for _, set := range sets {
		search(&head, set)
	}

}

func createLinklist(head *Node) *Node {
	if head != nil {
		head.next = &Node{
			data: head.data + 1,
		}
		return head.next
	}
	panic("head != nil")
}

func search(head *Node, set int) {
	index := 0
	first := head.next
	second := head.next
	for {
		index = index + 1
		for index := 0; index < set; index++ {
			first = first.next
		}
		second = second.next
		if first == second {
			fmt.Println("index:", index, "midindex:", midindex, "count:", count, "set:", set, "first:", first.data, "second:", second.data)
			break
		}
	}
}
