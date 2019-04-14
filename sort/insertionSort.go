package sort

type Link struct {
	data []int
	next []int  // 下标数组
}

func InsertionSort(nums Sortable)  {
	length:=nums.Len()
	for i:= 0; i<length-1; i++{
		for j:=i+1; j>0; j-- {
			if nums.Less(j,j-1){
				nums.Swap(j,j-1)
			} else {
				break
			}
		}
	}
}

// 更适合链表
//func InsertionSort(nums Sortable)  {
//	length := nums.Len()
//	link := &Link{make([]int,length),make([]int,length)}
//	for i:=0; i < length; i+= 1{
//		for j:=i; j>=0; j+=1 {
//			if !nums.Less(i,j){
//
//			}
//		}
//		nums.
//	}
//}
