package sort

func BubbleSort(nums Sortable) {
	length := nums.Len()
	for i := 0; i < length; i++ {
		for j := i; j < length-i-1; j++ {
			if nums.Less(j, j+1) {
				nums.Swap(j, j+1)
			}
		}
	}
}
