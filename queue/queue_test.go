package queue

import (
	"testing"
)

func test(t *testing.T, actualValue interface{}, exceptedValue interface{})  {
	if actualValue != exceptedValue{
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestNewQueue(t *testing.T) {
	q1 := NewQueue(0)
	test(t,q1.Empty(),true)
	//q2:=NewQueue(-1)
	//if q2 != nil {
	//	t.Errorf("Got %v expected %v", q2, true)
	//}
	q3:=NewQueue(4)
	test(t, q3.Empty(),true)
	test(t, q3.Size(),0)

}

func TestQueue_Back(t *testing.T) {

}
func TestQueue_Push(t *testing.T) {
	q1 := NewQueue(4)
	q1.Push(1)
	//q1.Push(2)

	test(t,q1.Size(),1)

}

func TestQueue_Pop(t *testing.T) {
	//q1 := NewQueue()
}

func TestNewFile(t *testing.T) {

}

func TestQueue_Empty(t *testing.T) {

}

func TestQueue_Front(t *testing.T) {

}

func TestQueue_Size(t *testing.T) {

}
