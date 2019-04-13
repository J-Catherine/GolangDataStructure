package queue

import "fmt"

type Deque struct {
	data []int
	size int
	head int
	tail int
	maxLen int
}

func NewDeque(len int) *Deque {
	if len<0{
		fmt.Println("New a deque failure!! Len is an illegal input! ")
		return nil
	}
	d := &Deque{make([]int, len), 0, 0, -1,len}
	for i := 0; i< len; i++ {
		d.data[i] = -1e10
	}
	return d
}


func (d *Deque) PushBack(value int) {
	if d.size >= d.maxLen{
		fmt.Println("Push back failure!! Deque is full!!")
	}else{
		d.tail = (d.tail+1)%d.maxLen
		d.data[d.tail] = value
		d.size+=1
	}
}
func (d *Deque) PushFront(value int) {
	if d.size >= d.maxLen{
		fmt.Println("Push front failure!! Deque is full!!")
	}else{
		d.head = (d.head-1)%d.maxLen
		d.data[d.head] = value
		d.size+=1
	}
}

func (d *Deque) PopBack() int{
	if d.size <= 0 {
		fmt.Println("Pop back failure!! Deque is empty!!")
		return -1e10
	}else{
		tmp := d.data[d.tail]
		d.data[d.tail] = -1e10
		d.tail = (d.tail-1)%d.maxLen
		d.size-=1
		return tmp
	}
}

func (d *Deque) PopFront() int{
	if d.size <= 0 {
		fmt.Println("Pop front failure!! Deque is empty!!")
		return 1e10
	}else{
		tmp := d.data[d.head]
		d.data[d.head] = -1e10
		d.head = (d.head+1)%d.maxLen
		d.size-=1
		return tmp
	}
}

func (d *Deque) Back() int{
	if d.size <= 0 {
		fmt.Println("Get back failure!! Deque is empty!!")
	}
	return d.data[d.tail]
}

func (d *Deque) Front() int{
	if d.size <= 0 {
		fmt.Println("Get front failure!! Deque is empty!!")
	}
	return d.data[d.head]
}

func (d *Deque) Size() int{
	return d.size
}

func (d *Deque) Empty() bool{
	if d.size == 0{
		return true
	} else {
		return false
	}
}