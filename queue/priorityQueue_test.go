package queue

import "testing"

func TestNewPriorityQueue(t *testing.T) {
	h1 := NewPriorityQueue(4,true)
	h2:= NewPriorityQueue(5,false)
	myTest(t,h1.Empty(),true)
	myTest(t,h2.Size(),0)
}

func TestPriorityQueue_Sample(t *testing.T) {
	h1 := NewPriorityQueue(4,true)
	h2:= NewPriorityQueue(5,false)
	h1.Push(2)
	h1.Push(5)
	h1.Push(3)
	myTest(t,h1.Top(),5)
	h1.Pop()
	myTest(t,h1.Top(),3)
	h1.Push(8)
	myTest(t,h1.Top(),8)
	h2.Push(2)
	h2.Push(5)
	h2.Push(8)
	myTest(t,h2.Top(),2)
	h2.Push(0)
	h2.Pop()
	myTest(t,h2.Top(),2)
	h2.Push(1)
	myTest(t,h2.Top(),1)
}
