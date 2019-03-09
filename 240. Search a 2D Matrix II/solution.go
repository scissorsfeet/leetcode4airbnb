package _40__Search_a_2D_Matrix_II

func searchMatrix(matrix [][]int, target int) bool {
	rowLen := len(matrix)
	colLen := len(matrix[0])
	if rowLen == 0 || colLen == 0 {
		return false
	}

	low := 0
	high := rowLen - 1
	for i := 0; i < colLen; i++ {
		low = 0
		for low <= high {
			mid := (low + high) >> 1
			if target == matrix[mid][i] {
				return true
			} else if target > matrix[mid][i] {
				if mid < rowLen-1 && target < matrix[mid+1][i] {
					high = mid
					break
				} else {
					low = mid + 1
				}
			} else {
				high = mid - 1
			}
		}
	}

	return false
}
