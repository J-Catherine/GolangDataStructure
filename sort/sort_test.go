package sort

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
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

func TestSort_Sample(t *testing.T) {
	intArr := IntArr{2, 3, 1, -9, 0}
	ans := IntArr{-9, 0, 1, 2, 3}

	QuickSort(intArr)
	for i := 0; i != intArr.Len(); i++ {
		myTest(t, intArr[i], ans[i])
	}
}

func TestSort_Large(t *testing.T) {
	M := 2000
	rand.Seed(time.Now().Unix())
	origin := make(IntArr, M)
	for i := 0; i != M; i++ {
		origin[i] = rand.Int()
	}
	fmt.Println()
	ans1 := make(IntArr, M)
	ans2 := make(IntArr, M)
	ans3 := make(IntArr, M)
	ans4 := make(IntArr, M)
	ans5 := make(IntArr, M)
	copy(ans1,origin)
	copy(ans2,origin)
	copy(ans3,origin)
	copy(ans4,origin)
	copy(ans5,origin)

	QuickSort(ans1)
	BubbleSort(ans2)
	SelectionSort(ans3)
	InsertionSort(ans4)
	HeapSort(ans5)
	for i := 0; i != origin.Len(); i++ {
		myTest(t, ans1[i], ans2[i])
		myTest(t, ans2[i], ans3[i])
		myTest(t, ans3[i], ans4[i])
		myTest(t, ans1[i], ans5[i])
	}
	//fmt.Println("A:",ans1)
	//fmt.Println("B:",ans2)
	//fmt.Println("C:",ans3)
	//fmt.Println("D:",ans4)
	//fmt.Println("E:",ans5)

}

func TestMergeSort(t *testing.T) {
	M := 2000000
	rand.Seed(time.Now().Unix())
	origin := make([]int, M)
	reverse := make([]int, M)
	for i := 0; i != M; i++ {
		origin[i] = i
		reverse[i] = M-1-i
	}
	//fmt.Println(origin)

	reverse = MergeSort(reverse)
	//fmt.Println(reverse)
	for i := 0; i < M; i++ {
		myTest(t, reverse[i], origin[i])
	}

}


func TestQuickSort_Corner(t *testing.T) {
	M := 2000000
	rand.Seed(time.Now().Unix())
	origin := make(IntArr, M)
	reverse := make(IntArr, M)
	for i := 0; i != M; i++ {
		origin[i] = i
		reverse[i] = M-1-i
	}
	ans := origin
	QuickSort(ans)
	for i := 0; i != origin.Len(); i++ {
		myTest(t, ans[i], origin[i])
	}
	QuickSort(reverse)
	for i := 0; i != origin.Len(); i++ {
		myTest(t, reverse[i], origin[i])
	}
}