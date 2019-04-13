package stack

import "testing"


func TestNewMStack(t *testing.T)  {
	ms1 := NewMStack(0,false)
	myTest(t,ms1.Empty(),true)
	ms2:=NewStack(4)
	myTest(t, ms2.Empty(),true)
	myTest(t, ms2.Size(),0)
}

func TestMStack_Sample(t *testing.T) {
	ms1 := NewMStack(10,true)
	for i:=0; i<11; i++{
		ms1.Push(i)
	}
	myTest(t,ms1.Empty(),false)
	myTest(t,ms1.Size(),10)
	myTest(t,ms1.Top(),9)
	myTest(t,ms1.Pop(),9)

	ms2 := NewMStack(10,false)
	for i:=10; i>0; i--{
		ms2.Push(i)
	}
	myTest(t,ms2.Empty(),false)
	myTest(t,ms2.Size(),10)
	myTest(t,ms2.Top(),1)
	myTest(t,ms2.Pop(),1)
}