package stack

import "testing"


func TestNewMStack(t *testing.T)  {
	ms1 := NewMStack(0,false)
	test(t,ms1.Empty(),true)
	ms2:=NewStack(4)
	test(t, ms2.Empty(),true)
	test(t, ms2.Size(),0)
}

func TestMStack_Push(t *testing.T) {
	ms1 := NewMStack(10,true)
	for i:=0; i<11; i++{
		ms1.Push(i)
	}
	test(t,ms1.Empty(),false)
	test(t,ms1.Size(),10)
	test(t,ms1.Top(),9)
	test(t,ms1.Pop(),9)

	ms2 := NewMStack(10,false)
	for i:=10; i>0; i--{
		ms2.Push(i)
	}
	test(t,ms2.Empty(),false)
	test(t,ms2.Size(),10)
	test(t,ms2.Top(),1)
	test(t,ms2.Pop(),1)
}