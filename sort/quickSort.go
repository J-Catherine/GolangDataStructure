package sort

import (
	"math/rand"
	"time"
)

func QuickSort(nums Sortable)  {
	length := nums.Len()
	left,right := 0,length-1
	rand.Seed(time.Now().Unix())
	qs(nums,left,right)

}

func qs(nums Sortable, start, end int)  {
	if end <= start{
		return
	}
	tmp := start + rand.Intn(end-start+1)
	nums.Swap(start,tmp)
	left, right := start, end
	for ; left < right; {
		for ;left < right && nums.Less(start,right); right-=1{}
		for ;left < right && !nums.Less(start,left); left+=1{}
		nums.Swap(left, right)
	}
	nums.Swap(left,start)
	qs(nums,start,left-1)
	qs(nums,left+1,end)

}

