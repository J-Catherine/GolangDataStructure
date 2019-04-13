package sort

func SelectionSort(nums Sortable)  {
	length := nums.Len()
	minIndex := -1
	for i:=0; i<length-1; i++ {
		minIndex = i
		for j:=i+1; j< length; j++{
			if nums.Less(minIndex, j){
				minIndex = j
			}
		}
		nums.Swap(minIndex, i)
	}
}