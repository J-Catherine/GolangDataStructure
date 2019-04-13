package queue

import "testing"

func myTest(t *testing.T, actualValue interface{}, exceptedValue interface{})  {
	if actualValue != exceptedValue{
		t.Errorf("Got %v expected %v", actualValue, exceptedValue)
	}
}