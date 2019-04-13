package stack

import "fmt"

// 单调栈
type MinStack struct {
	diff *Stack
	min int
}


func NewMinStack(len int)  *MinStack{
	if len<0{
		fmt.Println("New a min_stack failure!! Len is an illegal input! ")
		return nil
	}
	minStack := &MinStack{NewStack(len), 1e10}
	return minStack
}

func (minStack *MinStack) Push(value int) {
	sd := minStack.diff
	if sd.size >= sd.maxLen{
		fmt.Println("Push failure!! Min_stack is full!!")
	} else{
		if sd.size == 0{
			minStack.min = value
		}
		sd.data[sd.size] = value - minStack.min
		sd.size+=1

		if sd.data[sd.size-1] < 0{
			minStack.min = value
		}
	}
}

func (minStack *MinStack) Pop() int{
	sd := minStack.diff
	if sd.size <=0{
		fmt.Println("Pop failure!! Min_stack stack is empty!!")
		return -1e10
	} else {
		sd.size -= 1
		tmp := 0
		if sd.data[sd.size] < 0{
			tmp = minStack.min
			if sd.size == 0{
				minStack.min = 1e10
			} else{
				minStack.min -= sd.data[sd.size]
			}
			return tmp
		} else {
			tmp = minStack.min
			if sd.size == 0{
				minStack.min = 1e10
			}
			return tmp+sd.data[sd.size]
		}
	}
}

func (minStack *MinStack) Empty() bool{
	sd := minStack.diff
	if sd.size <= 0{
		return true
	} else {
		return false
	}
}

func (minStack *MinStack) Size() int{
	return minStack.diff.size
}

func (minStack *MinStack) Top() int {
	sd := minStack.diff
	if sd.size <= 0 {
		fmt.Println("Get top failure!! Stack is empty!!")
		return -1e10
	} else {
		if sd.data[sd.size-1] < 0{
			return minStack.min
		} else {
			return minStack.min + sd.data[sd.size-1]
		}
	}
}

func (minStack *MinStack) Min() int{
	return minStack.min
}