package queue

import "fmt"

type Queue struct {
	data []int
	size int
	head int
	end int
	maxLen int
}

func NewQueue(len int) *Queue {
	if len<0{
		fmt.Println("NewQueue failure!! Len is an illegal input! ")
		return nil
	}
	q := &Queue{make([]int, len), 0, 0, 0,len}
	return q
}

func (q *Queue) Push(value int) {
	if q.size >= q.maxLen{
		fmt.Println("Push failure!! Queue is full!!")
	}else{
		q.data[q.end] = value
		q.end++
		q.size++
	}
}

func (q *Queue) Pop() int{
	if q.size > 0{
		tmp := q.data[q.head]
		q.size--
		if q.size == 0{
			q.head,q.end = 0,0
		} else{
			q.head++
		}
		return tmp
	} else{
		fmt.Println("Pop failure!! Queue is empty!!")
		return -1e10
	}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Empty() bool {
	if q.size == 0{
		return true
	} else {
		return false
	}
}

func (q *Queue) Front() int{
	if q.size == 0{
		fmt.Println("Front failure!! Queue is empty!!")
		return -1e10
	} else{
		return q.data[q.head]
	}
}
func (q *Queue) Back()  int{
	if q.end == 0{
		fmt.Println("Back failure!! Queue is empty!!")
		return -1e10
	} else{
		return q.data[q.end]
	}
}

