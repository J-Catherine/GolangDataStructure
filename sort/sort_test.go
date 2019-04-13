package sort

import (
	"testing"
	"math/rand"
	"time"
)

type IntArr []int

func (arr IntArr) Len() int {
	return len(arr)
}
func (arr IntArr) Less(i int, j int) bool {
	return arr[i] < arr[j]
}
func (arr IntArr) Swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr IntArr) Get(i int) int {
	return arr[i]
}
func (arr IntArr) Set(i int, t int) {
	arr[i] = t
}

func TestQuickSort_Sample(t *testing.T) {
	intArr := IntArr{2, 3, 1, -9, 0}
	ans := IntArr{-9, 0, 1, 2, 3}
	QuickSort(intArr)
	for i := 0; i != intArr.Len(); i++ {
		myTest(t, intArr[i], ans[i])
	}
}

func TestQuickSort_Large(t *testing.T) {
	rand.Seed(time.Now().Unix())
	origin := make(IntArr, 10000)
	for i := 0; i != 10000; i++ {
		origin[i] = rand.Int()
	}
	ans1 := origin
	ans2 := origin
	ans3 := origin
	QuickSort(ans1)
	BubbleSort(ans2)
	SelectionSort(ans3)
	for i := 0; i != origin.Len(); i++ {
		myTest(t, ans1[i], ans2[i])
		myTest(t, ans2[i], ans3[i])
	}
}
