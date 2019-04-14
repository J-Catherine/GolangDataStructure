package sort

import (
	"math/rand"
	//"fmt"
	"time"
)

func QuickSort(nums Sortable)  {
	length := nums.Len()
	left,right := 0,length-1
	rand.Seed(time.Now().Unix())
	partition(nums,left,right)

}

func partition(nums Sortable, start, end int)  {
	if end <= start{
		return
	}
	tmp := start + rand.Intn(end-start+1)
	nums.Swap(start,tmp)
	left, right := start+1, end
	for ; left < right; {
		for ;left < right && !nums.Less(right,start); right-=1{}
		for ;left < right && nums.Less(left,start); left+=1{}
		nums.Swap(left, right)
	}
	if nums.Less(left,start) {
		nums.Swap(left,start)
	} else {
		left -= 1
	}
	partition(nums,start,left-1)
	partition(nums,left+1,end)

}

