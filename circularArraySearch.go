package main

func CircularArraySearch(arr []int, target int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {

		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] <= arr[high] {

			if target > arr[mid] && target <= arr[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if target >= arr[low] && target < arr[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}

		}

	}
	return -1

}