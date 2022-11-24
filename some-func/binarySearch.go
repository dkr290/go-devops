package main

/// position of the item in items

func binarySearch(array []int, target, lowIndex, highIndex int) int {

	if highIndex < lowIndex {
		return -1
	}

	mid := int((lowIndex + highIndex) / 2)

	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}

}

func binarySearchV1(arr []int, target, startIndex, endIndex int) int {
	var mid int
	st := startIndex
	se := endIndex
	for st <= se {
		mid = int((startIndex + endIndex) / 2)

		if arr[mid] > target {
			endIndex = mid
		} else if arr[mid] < target {
			startIndex = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
