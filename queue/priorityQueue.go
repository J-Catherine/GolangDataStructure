package queue

import (
	"fmt"
)

// heap
type PriorityQueue struct {
	data   []int
	size   int
	maxLen int
	flag   bool // 1为最大堆，0为最小堆
}

func NewPriorityQueue(len int, flag bool) *PriorityQueue {
	heap := &PriorityQueue{make([]int, len), 0, len, flag}
	for i := 0; i < len; i++ {
		heap.data[i] = -1e10
	}
	return heap
}

func (h *PriorityQueue) Push(value int) {
	if h.size >= h.maxLen {
		fmt.Println("Push failure!! Heap is full!!")
	} else {
		f := -1
		if h.flag {
			f = 1
		}
		now := h.size
		h.data[now] = value
		h.size += 1
		h.up(now,f)
	}
}

func (h *PriorityQueue) up(i, f int) {
	for j := int((i-1)/2); j >= 0; i, j = j, int((j-1)/2) {
		if f*h.data[i] > f*h.data[j] {
			h.data[i], h.data[j] = h.data[j], h.data[i]
		} else {
			break
		}
	}
}

func (h *PriorityQueue) down(i int, f int) {
	for i, j := 0, 1; j < h.size; i, j = j, j*2+1 {
		if j+1 < h.size && h.data[j]*f < h.data[j+1]*f {
			j+= 1
		}
		if h.data[i]*f < h.data[j]*f {
			h.data[i], h.data[j] = h.data[j], h.data[i]
		}
		break
	}
}

func (h *PriorityQueue) Pop() int {
	if h.size <= 0 {
		fmt.Println("Get size failure!! Heap is empty!!")
		return -1e10
	} else {
		f := -1
		if h.flag {
			f = 1
		}
		h.size -= 1
		h.data[0], h.data[h.size] = h.data[h.size], h.data[0]
		h.down(0,f)
	}
	return h.data[h.size+1]
}

func (h *PriorityQueue) Top() int {
	if h.size > 0 {
		return h.data[0]
	} else {
		fmt.Println("Get max or min failure!! Heap is empty!!")
		return -1e10
	}
}

func (h *PriorityQueue) Empty() bool{
	if h.size >0{
		return false
	} else {
		return true
	}
}

func (h *PriorityQueue) Size() int{
	return h.size
}
