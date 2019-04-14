package sort


func HeapSort(nums Sortable)  {
	maxHeapify(nums)  // build heap
	sortHeap(nums)  // sort heap
}

func down(nums Sortable, i, length int)  {
	for j := i*2+1; j < length; i,j = j,j*2+1 {
		if j+1 < length && nums.Less(j,j+1) {
			j+= 1
		}
		if nums.Less(i,j) {
			nums.Swap(i,j)
		} else{
			break
		}
	}
}

// build heap
func maxHeapify(nums Sortable)  {
	length := nums.Len()
	now := int(length/2)-1
	for i:= now; i>=0; i--{
		down(nums,i,length)
	}
}

func sortHeap(nums Sortable)  {
	length:=nums.Len()
	for i:= length-1; i>=0; i--{
		nums.Swap(0,i)
		down(nums,0,i)
	}

}
