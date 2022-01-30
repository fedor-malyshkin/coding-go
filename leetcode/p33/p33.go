package p33

func search(nums []int, target int) int {

	n := len(nums)

	if n == 1 {
		switch nums[0] == target {
		case true:
			return 0
		case false:
			return -1
		}
	}

	rotateIndex := findRotateIndex(nums, 0, n-1)

	// if target is the smallest element
	switch {
	case nums[rotateIndex] == target:
		return rotateIndex
	// if array is not rotated, search in the entire array
	case rotateIndex == 0:
		return binarySearch(nums, 0, n-1, target)
	case target < nums[0]:
		// search in the right side
		return binarySearch(nums, rotateIndex, n-1, target)
	// search in the left side
	default:
		return binarySearch(nums, 0, rotateIndex, target)

	}
}

func binarySearch(nums []int, left int, right int, target int) int {
	for left <= right {
		pivot := (left + right) / 2
		switch {
		case nums[pivot] == target:
			return pivot
		case target < nums[pivot]:
			right = pivot - 1
		default:
			left = pivot + 1
		}
	}
	return -1
}

func findRotateIndex(nums []int, left int, right int) int {
	if nums[left] < nums[right] {
		return 0
	}

	for left <= right {
		pivot := (left + right) / 2
		switch {
		case nums[pivot] > nums[pivot+1]:
			return pivot + 1
		case nums[pivot] < nums[left]:
			right = pivot - 1
		default:
			left = pivot + 1
		}
	}
	return 0
}
