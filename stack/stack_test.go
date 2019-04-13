package stack

import "testing"

func TestNewStack(t *testing.T)  {
	s1 := NewStack(0)
	myTest(t,s1.Empty(),true)
	s2:=NewStack(4)
	myTest(t, s2.Empty(),true)
	myTest(t, s2.Size(),0)
}

func TestStack_Sample(t *testing.T) {
	s1 := NewStack(10)
	for i:=0; i<11; i++{
		s1.Push(i)
	}
	myTest(t,s1.Empty(),false)
	myTest(t,s1.Size(),10)
	myTest(t,s1.Top(),9)
	myTest(t,s1.Pop(),9)
}
