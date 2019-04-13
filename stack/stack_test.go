package stack

import "testing"

func test(t *testing.T, actualValue interface{}, exceptedValue interface{})  {
	if actualValue != exceptedValue{
		t.Errorf("Got %v expected %v", actualValue, exceptedValue)
	}
}


func TestNewStack(t *testing.T)  {
	s1 := NewStack(0)
	test(t,s1.Empty(),true)
	s2:=NewStack(4)
	test(t, s2.Empty(),true)
	test(t, s2.Size(),0)
}

func TestStack_Push(t *testing.T) {
	s1 := NewStack(10)
	for i:=0; i<11; i++{
		s1.Push(i)
	}
	test(t,s1.Empty(),false)
	test(t,s1.Size(),10)
	test(t,s1.Top(),9)
	test(t,s1.Pop(),9)
}