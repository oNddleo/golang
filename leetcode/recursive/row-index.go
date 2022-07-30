package recursive

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	// recursive call rowIndex -1
	return recursive([]int{1, 1}, 2, rowIndex)
}
func recursive(prev []int, current, target int) []int {
	row := make([]int, current+1)
	row[0], row[len(row)-1] = 1, 1
	for i := 1; i < current; i++ {
		row[i] = prev[i-1] + prev[i]
	}
	if current == target {
		return row
	} else {
		return recursive(row, current+1, target)
	}

}
