package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q1 := NewQueue(0)
	myTest(t,q1.Empty(),true)
	q3:=NewQueue(4)
	myTest(t, q3.Empty(),true)
	myTest(t, q3.Size(),0)

}


func TestQueue_Sample(t *testing.T) {
	q1 := NewQueue(4)
	q1.Push(1)
	q1.Push(2)

	myTest(t,q1.Size(),2)

	myTest(t,q1.Pop(),1)
	myTest(t,q1.Front(),2)
	myTest(t,q1.Pop(),2)
	myTest(t,q1.Pop(),-10000000000)
}
