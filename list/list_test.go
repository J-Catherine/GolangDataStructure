package list

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewNodeSpace(t *testing.T) {
	sp := NewNodeSpace(1)
	t1 := sp.NewNode(1)
	myTest(t, t1, 1)
	t2 := sp.NewNode(3)
	myTest(t, t2, 0)
}

func TestList_Base(t *testing.T) {
	sp := NewNodeSpace(10000)
	l1 := NewList(sp)
	l1.Insert(12, 0)
	l1.Insert(13, l1.Head)
	l1.Insert(15, l1.Head)
	l1.Insert(17, l1.Head)
	res := []Node{12, 17, 15, 13}
	for p, i := l1.Head, 0; p != 0; p, i = sp.Next(p), i+1 {
		myTest(t, sp.Get(p), res[i])
	}
	myTest(t, l1.Size(), 4)
	fmt.Println()
	i2 := l1.Delete(l1.Head)
	l1.Delete(i2)

	myTest(t, l1.Size(), 2)
	res = []Node{12, 15}

	for p, i := l1.Head, 0; p != 0; p, i = sp.Next(p), i+1 {
		myTest(t, sp.Get(p), res[i])
	}
}

func TestList_Large(t *testing.T) {
	sizeV, maxE := 30, 10000
	sp := NewNodeSpace(maxE)
	P := NewListCluster(sizeV, sp)
	for i := 0; i != sizeV; i++ {
		p := P[i].Head
		for j := 0; j != sizeV; j++ {
			if i != j {
				p = P[i].Insert(Node(j), p)
			}
		}
	}
	myTest(t, sp.size, 870)
}

func TestList_Corner(t *testing.T) {
	rand.Seed(time.Now().Unix())
	sizeV, maxE := 3000000, 10000000
	sp := NewNodeSpace(maxE)
	P := NewListCluster(sizeV, sp)
	for i := 0; i != maxE; i++ {
		from := rand.Intn(sizeV)
		to := rand.Intn(sizeV)
		P[from].Insert(Node(to), 0)
	}
	myTest(t, sp.size, maxE)
}
