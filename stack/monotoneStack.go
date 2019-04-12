package stack

import "fmt"

// 单调栈
type MStack struct {
	data []int
	top int
	size int
	maxLen int
	flag bool  // 单调增：true   单调减：false
}


func NewMStack(len int, flag bool)  *MStack{
	if len<0{
		fmt.Println("New a monotone stack failure!! Len is an illegal input! ")
		return nil
	}
	ms := &MStack{make([]int,len),-1,0,len, flag}
	for i:= 0; i<len; i++{
		ms.data[i] = -1e10
	}
	return ms
}

func (s *MStack) Push(value int) {
	if s.size >= s.maxLen{
		fmt.Println("Push failure!! Monotone stack is full!!")
	} else{
		f:= -1
		if s.flag {
			f = 1
		}
		for i:= s.top; i >= 0 && s.data[i]*f > value*f; i--{
			s.Pop()
		}
		s.top += 1
		s.data[s.top] = value
		s.size += 1
	}
}

func (s *MStack) Pop() int{
	if s.size <=0{
		fmt.Println("Pop failure!! Monotone stack is empty!!")
		return -1e10
	} else {
		tmp := s.data[s.top]
		s.top -= 1
		s.size -= 1
		return tmp
	}
}

func (s *MStack) Empty() bool{
	if s.size <= 0{
		return true
	} else {
		return false
	}
}

func (s *MStack) Size() int{
	return s.size
}

func (s *MStack) Top() int {
	if s.size <= 0 {
		fmt.Println("Get top failure!! Stack is empty!!")
		return -1e10
	} else {
		return s.data[s.top]
	}
}