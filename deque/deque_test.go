package deque

import "testing"

func test(t *testing.T, actualValue interface{}, exceptedValue interface{})  {
	if actualValue != exceptedValue{
		t.Errorf("Got %v expected %v", actualValue, exceptedValue)
	}
}

func TestNewDeque(t *testing.T) {
	d1 := NewDeque(0)
	test(t,d1.Empty(),true)
	d3:=NewDeque(4)
	test(t, d3.Empty(),true)
	test(t, d3.Size(),0)

}

func TestDeque_PushBack(t *testing.T) {
	d1 := NewDeque(10)
	for i:=0; i<11; i++{
		d1.PushBack(i)
	}
	test(t,d1.Empty(),false)
	test(t,d1.Size(),10)
	test(t,d1.Back(),9)
	test(t,d1.PopBack(),9)
}