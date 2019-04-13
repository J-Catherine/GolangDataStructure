package stack

import (
	"testing"
	"math/rand"
	"time"
)

func TestMinStack_Sample(t *testing.T) {
	s1 := NewMinStack(5)
	myTest(t,s1.Empty(),true)
	s1.Push(2)
	s1.Push(1)
	s1.Push(4)
	myTest(t,s1.Top(),4)
	myTest(t,s1.Min(),1)
	myTest(t,s1.Pop(),4)
	myTest(t,s1.Pop(),1)
	myTest(t,s1.Min(),2)
	myTest(t,s1.Pop(),2)
	myTest(t,s1.Min(),int(1e10))
}

func TestMinStack_Large(t *testing.T) {
	rand.Seed(time.Now().Unix())
	s1 := NewMinStack(1000)
	s2 := make([]int,1000)
	s2Size := 0
	for i:=0;i!=10;i++{
		randOne:= rand.Intn(1e10-1)
		if i%3==1{
			s2Size -= 1
			s1.Pop()
		}else {
			s1.Push(randOne)
			s2[s2Size] = randOne
			s2Size += 1
		}
		min := int(1e10)
		for k:=0;k!=s2Size;k++ {
			if s2[k] < min {
				min = s2[k]
			}

		}
		myTest(t, s1.Min(), min)
	}
}