package list

import "fmt"

type NodeSpace struct {
	data   []Node
	next   []int
	maxLen int
	head   int
	start  int
	size   int
}

func NewNodeSpace(maxLen int) *NodeSpace {
	ret := NodeSpace{}
	ret.data = make([]Node, maxLen+1)
	ret.next = make([]int, maxLen+1)
	ret.maxLen = maxLen
	ret.start = 1
	return &ret
}

func (space *NodeSpace) NewNode(value Node) int {
	ret := 0
	if space.start <= space.maxLen {
		ret = space.start
		space.start += 1
	} else if space.head != 0 {
		ret = space.head
		space.head = space.next[space.head]
	}
	if ret != 0 {
		space.data[ret] = value
	} else {
		fmt.Println("NodeSpace is full!")
	}
	space.size++
	return ret
}

func (space *NodeSpace) ReleaseNode(index int) {
	space.next[index] = space.head
	space.head = index
	space.size--
}

func (space *NodeSpace) Size(index int) int {
	return space.size
}

func (space *NodeSpace) Next(index int) int {
	return space.next[index]
}

func (space *NodeSpace) SetNext(index, next int) {
	space.next[index] = next
}

func (space *NodeSpace) Get(index int) Node {
	return space.data[index]
}

func (space *NodeSpace) Set(index int, data Node) {
	space.data[index] = data
}

type NodeList int

type List struct {
	space *NodeSpace
	Head  int
	size  int
}

func NewList(data *NodeSpace) *List {
	return &List{data, 0, 0}
}

func NewListCluster(size int, sp *NodeSpace) []List {
	cluster := make([]List, size)
	for i := 0; i != size; i++ {
		cluster[i].SetNodeSpace(sp)
	}
	return cluster
}

func (list *List) SetNodeSpace(space *NodeSpace) {
	list.space = space
}

func (list *List) Insert(value Node, after int) int {
	var index int
	if after == 0 {
		index = list.space.NewNode(value)
		list.space.SetNext(index, list.Head)
		list.Head = index
	} else {
		index = list.space.NewNode(value)
		list.space.SetNext(index, list.space.Next(after))
		list.space.SetNext(after, index)
	}
	list.size++
	return index
}

func (list *List) Delete(after int) int {
	var del, ret int
	if after == 0 {
		del = list.Head
		list.Head = list.space.Next(list.Head)
		ret = list.Head
	} else {
		del = list.space.Next(after)
		list.space.SetNext(after, list.space.Next(del))
		ret = list.space.Next(after)
	}
	if del != 0 {
		list.space.ReleaseNode(del)
		list.size--
	}
	return ret
}

func (list *List) Size() int {
	return list.size
}
