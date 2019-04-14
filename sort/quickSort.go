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
	qs(nums,left,right)

}

func qs(nums Sortable, left, right int)  {
	if right <= left{
		return
	}
	end:=right
	tmp := left + rand.Intn(right-left+1)
	nums.Swap(left,tmp)
	start := left
	left+=1
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
	qs(nums,start,left-1)
	qs(nums,left+1,end)

}

