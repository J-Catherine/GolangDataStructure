package stack

import "fmt"

type Stack struct {
	data []int
	size int
	maxLen int
}

func NewStack(len int)  *Stack{
	if len<0{
		fmt.Println("New a stack failure!! Len is an illegal input! ")
		return nil
	}
	s := &Stack{make([]int,len),0,len}
	for i:= 0; i<len; i++{
		s.data[i] = -1e10
	}
	return s
}

func (s *Stack) Push(value int) {
	if s.size >= s.maxLen{
		fmt.Println("Push failure!! Stack is full!!")
	} else{
		s.data[s.size] = value
		s.size+=1
	}
}

func (s *Stack) Pop() int{
	if s.size <=0{
		fmt.Println("Pop failure!! Stack is empty!!")
		return -1e10
	} else {
		s.size -= 1
		tmp := s.data[s.size]
		return tmp
	}
}

func (s *Stack) Empty() bool{
	if s.size <= 0{
		return true
	} else {
		return false
	}
}

func (s *Stack) Size() int{
	return s.size
}

func (s *Stack) Top() int {
	if s.size <= 0 {
		fmt.Println("Get top failure!! Stack is empty!!")
		return -1e10
	} else {
		return s.data[s.size-1]
	}
}