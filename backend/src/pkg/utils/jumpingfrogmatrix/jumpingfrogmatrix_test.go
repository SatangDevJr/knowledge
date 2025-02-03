package jumpingfrogmatrix_test

import (
	jumpingFrogMatrix "neversitup/src/pkg/utils/jumpingfrogmatrix"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpingFrogMatrix_LongestIncreasingPath(t *testing.T) {
	t.Run("should return 0 (route 0) when input empty matrix", func(t *testing.T) {
		matrix := [][]int{}

		got := jumpingFrogMatrix.LongestIncreasingPath(matrix)

		assert.Equal(t, 0, got)
	})

	t.Run("should return 0 (route 0) when input matrix with first row is empty", func(t *testing.T) {
		matrix := [][]int{
			{},
			{6, 6, 8},
			{2, 1, 1},
		}

		got := jumpingFrogMatrix.LongestIncreasingPath(matrix)

		assert.Equal(t, 0, got)
	})

	t.Run("should return 9 (route 1->2->3->4->5->6->7->8->9) when input matrix {1, 2, 3}, {8, 9, 4}, {7, 6, 5}", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		}

		got := jumpingFrogMatrix.LongestIncreasingPath(matrix)

		assert.Equal(t, 9, got)
	})

	t.Run("should return 4 (route 1->2->6->9) when input matrix {9, 9, 4}, {6, 6, 8}, {2, 1, 1}", func(t *testing.T) {
		matrix := [][]int{
			{9, 9, 4},
			{6, 6, 8},
			{2, 1, 1},
		}

		got := jumpingFrogMatrix.LongestIncreasingPath(matrix)

		assert.Equal(t, 4, got)
	})

	t.Run("should return 4 (route 3->4->5->6) when input matrix {3, 4, 5}, {3, 2, 6}, {2, 2, 1}", func(t *testing.T) {
		matrix := [][]int{
			{3, 4, 5},
			{3, 2, 6},
			{2, 2, 1},
		}

		got := jumpingFrogMatrix.LongestIncreasingPath(matrix)

		assert.Equal(t, 4, got)
	})
}
