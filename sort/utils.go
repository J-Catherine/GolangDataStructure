package sort

import (
	"testing"
)


// 定义可排序接口
type Sortable interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

type SortableIndexing interface {
	Sortable
	Get(int) interface{}
	Set(int, interface{})
}

func myTest(t *testing.T, actualValue interface{}, exceptedValue interface{})  {
	if actualValue != exceptedValue{
		t.Errorf("Got %v expected %v", actualValue, exceptedValue)
	}
}

