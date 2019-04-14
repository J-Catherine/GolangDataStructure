package list

import (
	"testing"
)

type Node int

// 需要使用不同类型时重新定义Node

func myTest(t *testing.T, actualValue interface{}, exceptedValue interface{}) {
	if actualValue != exceptedValue {
		t.Errorf("Got %v expected %v", actualValue, exceptedValue)
	}
}
