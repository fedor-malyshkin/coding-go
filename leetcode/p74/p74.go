package p74

/**
https://leetcode.com/problems/search-a-2d-matrix/


*/
func searchMatrix(matrix [][]int, target int) bool {
	// find a row
	l := 0
	r := len(matrix) - 1
	m := (l + r) / 2
	for l <= r {
		m = (l + r) / 2
		if matrix[m][0] <= target && matrix[m][len(matrix[m])-1] >= target {
			break
		}
		if matrix[m][0] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	// m - is probably our row
	row := m
	l = 0
	r = len(matrix[row]) - 1
	for l <= r {
		m := (l + r) / 2
		switch {
		case matrix[row][m] > target:
			r = m - 1
		case matrix[row][m] < target:
			l = m + 1
		case matrix[row][m] == target:
			return true
		}
	}
	return false
}
