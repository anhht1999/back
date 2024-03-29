package algorithm

//bubbleSort
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//MergeSort
func Merge(l, r []int) []int {

	arr := make([]int, 0, len(l)+len(r))

	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(arr, r...)
		}
		if len(r) == 0 {
			return append(arr, l...)
		}
		if l[0] <= r[0] {
			arr = append(arr, l[0])
			l = l[1:]
		} else {
			arr = append(arr, r[0])
			r = r[1:]
		}
	}
	return arr
}

func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	n := len(arr) / 2
	l := MergeSort(arr[:n])
	r := MergeSort(arr[n:])

	return Merge(l, r)
}

//quickSort
func Partition(arr []int, low, high int) int {

	p := arr[high]

	for j := low; j < high; j++ {
		if arr[j] < p {
			arr[j], arr[low] = arr[low], arr[j]
			low++
		}
	}

	arr[low], arr[high] = arr[high], arr[low]
	return low
}

func QuickSort(arr []int, low, high int) []int {

	if low > high {
		return arr
	}

	p := Partition(arr, low, high)
	QuickSort(arr, low, p-1)
	QuickSort(arr, p+1, high)
	return arr
}

