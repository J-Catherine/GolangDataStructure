package sort

func BubbleSort(nums Sortable) {
	length := nums.Len()
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums.Less(j+1, j) {
				nums.Swap(j, j+1)
			}
		}
	}
}
