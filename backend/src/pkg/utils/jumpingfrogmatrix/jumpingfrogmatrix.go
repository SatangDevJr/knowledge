package jumpingfrogmatrix

var directions = [][]int{
	{-1, 0}, // ⬆
	{1, 0},  // ⬇️
	{0, -1}, // ⬅️
	{0, 1},  // ➡️
}

func LongestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	memory := make([][]int, rows)
	for i := range memory {
		memory[i] = make([]int, cols)
	}

	var dfs func(int, int) int
	dfs = func(row, col int) int {
		if memory[row][col] != 0 {
			return memory[row][col]
		}

		maxLength := 1

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && matrix[newRow][newCol] > matrix[row][col] {
				maxLength = max(maxLength, 1+dfs(newRow, newCol))
			}
		}

		memory[row][col] = maxLength
		return maxLength
	}

	longestPath := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			longestPath = max(longestPath, dfs(i, j))
		}
	}

	return longestPath
}

func max(input1, input2 int) int {
	if input1 > input2 {
		return input1
	}
	return input2
}
