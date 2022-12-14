package myalgorithm

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivot
	return low
}

func QuickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	pivot := partition(arr, low, high)
	QuickSort(arr, low, pivot-1)
	QuickSort(arr, pivot+1, high)
}
