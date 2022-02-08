package main

func main() {

}

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	mid := (low + high) / 2
	for {
		if low < high {

		} else {
			break
		}
	}
	if arr[mid] == target {
		return mid
	}

	if arr[mid] > target {
		high = mid - 1
	}

	if arr[mid] < target {
		low = mid + 1
	}

	return -1
}
