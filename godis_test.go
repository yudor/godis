package godis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("connect successfully", func(t *testing.T) {
		godis, err := New("127.0.0.1:6379")
		assert.Nil(t, err)
		assert.NotNil(t, godis)
	})

	t.Run("Panic", func(t *testing.T) {
		godis, err := New("127.0.0.1:6378")
		assert.Nil(t, godis)
		assert.NotNil(t, err)
	})
}

func Test_Ping(t *testing.T) {
	godis, _ := New("127.0.0.1:6379")
	pong, _ := godis.Ping()
	fmt.Print(pong)
}

func Test_Get(t *testing.T) {
	godis, _ := New("127.0.0.1:6379")
	value, _ := godis.Get("hello")
	fmt.Print(value)
}
