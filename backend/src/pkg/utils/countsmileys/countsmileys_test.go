package countsmileys_test

import (
	"neversitup/src/pkg/utils/countsmileys"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmileys_CountSmileys(t *testing.T) {
	t.Run("should return 0 when input [':']", func(t *testing.T) {
		got := countsmileys.CountSmileys([]string{":"})
		assert.Equal(t, 0, got)
	})

	t.Run("should return 1 when input [':)']", func(t *testing.T) {
		got := countsmileys.CountSmileys([]string{":)"})
		assert.Equal(t, 1, got)
	})

	t.Run("should return 2 when input [':)', ';(', ';}', ':-D']", func(t *testing.T) {
		got := countsmileys.CountSmileys([]string{":)", ";(", ";}", ":-D"})
		assert.Equal(t, 2, got)
	})

	t.Run("should return 3 when input [';D', ':-(', ':-)', ';~)']", func(t *testing.T) {
		got := countsmileys.CountSmileys([]string{";D", ":-(", ":-)", ";~)"})
		assert.Equal(t, 3, got)
	})

	t.Run("should return 1 when input [';]', ':[', ';*', ':$', ';-D']", func(t *testing.T) {
		got := countsmileys.CountSmileys([]string{";]", ":[", ";*", ":$", ";-D"})
		assert.Equal(t, 1, got)
	})
}
