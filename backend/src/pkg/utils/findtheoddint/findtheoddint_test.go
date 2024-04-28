package findtheoddint_test

import (
	findTheOddInt "neversitup/src/pkg/utils/findtheoddint"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheOddInt_FindTheOddInt(t *testing.T) {
	t.Run("should return 7 when input [7]", func(t *testing.T) {
		got := findTheOddInt.FindOddInt([]int{7})

		assert.Equal(t, 7, got)
	})

	t.Run("should return 2 when input [1,1,2]", func(t *testing.T) {
		got := findTheOddInt.FindOddInt([]int{1, 1, 2})

		assert.Equal(t, 2, got)
	})

	t.Run("should return 4 when input [1,2,2,3,3,3,4,3,3,3,2,2,1]", func(t *testing.T) {
		got := findTheOddInt.FindOddInt([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})

		assert.Equal(t, 4, got)
	})
}
