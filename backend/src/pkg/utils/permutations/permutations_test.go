package permutations_test

import (
	permutations "neversitup/src/pkg/utils/permutations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations_Permutations(t *testing.T) {
	t.Run("should return slice with value string a when input a", func(t *testing.T) {
		got := permutations.Permutations("a")

		assert.Equal(t, []string{"a"}, got)
	})

	t.Run("should return slice with value string ab ba when input ab", func(t *testing.T) {
		got := permutations.Permutations("ab")

		assert.Equal(t, []string{"ab", "ba"}, got)
	})

	t.Run("should return slice with value string aa when input aa", func(t *testing.T) {
		got := permutations.Permutations("aa")

		assert.Equal(t, []string{"aa"}, got)
	})
}
