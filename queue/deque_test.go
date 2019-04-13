package queue

import "testing"


func TestNewDeque(t *testing.T) {
	d1 := NewDeque(0)
	myTest(t,d1.Empty(),true)
	d3:=NewDeque(4)
	myTest(t, d3.Empty(),true)
	myTest(t, d3.Size(),0)

}

func TestDeque_Sample(t *testing.T) {
	d1 := NewDeque(10)
	for i:=0; i<11; i++{
		d1.PushBack(i)
	}
	myTest(t,d1.Empty(),false)
	myTest(t,d1.Size(),10)
	myTest(t,d1.Back(),9)
	myTest(t,d1.PopBack(),9)
}