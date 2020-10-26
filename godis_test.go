package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("connect successfully", func(t *testing.T) {
		godis, err := New("127.0.0.1:6379")
		assert.Nil(t, err)
		assert.NotNil(t, godis)
	})

	t.Run("connect failed", func(t *testing.T) {
		godis, err := New("127.0.0.1:6378")
		assert.Nil(t, godis)
		assert.NotNil(t, err)
	})
}
