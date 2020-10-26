package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup(t *testing.T) godis {
	t.Helper()
	godis, _ := New("127.0.0.1:6379")
	godis.Flush()
	return *godis
}

func Test_Ping(t *testing.T) {
	godis := setup(t)
	pong, _ := godis.Ping()
	assert.Equal(t, "PONG", pong)
}

func Test_Get(t *testing.T) {
	t.Run("get nil key", func(t *testing.T) {
		godis := setup(t)
		value, _ := godis.Get("hello")
		assert.Equal(t, nil, value)
	})

	t.Run("get key, returns int", func(t *testing.T) {
		godis := setup(t)
		godis.Set("age", 8)
		value, _ := godis.Get("age")
		assert.Equal(t, 8, value)
	})

	t.Run("get key, returns string", func(t *testing.T) {
		godis := setup(t)
		name := "raymond"
		godis.Set("name", name)
		value, _ := godis.Get("name")
		assert.Equal(t, name, value)
	})
}

func Test_Set(t *testing.T) {
	t.Run("set string key", func(t *testing.T) {
		godis := setup(t)
		value, _ := godis.Set("name", "Raymond Tukpe")
		assert.Equal(t, true, value)
		v, _ := godis.Get("name")
		assert.Equal(t, "Raymond Tukpe", v)
	})

	t.Run("set int key", func(t *testing.T) {
		godis := setup(t)
		value, _ := godis.Set("age", 100)
		assert.Equal(t, true, value)
	})
}

func Test_Keys(t *testing.T) {
	t.Run("Empty array reply", func(t *testing.T) {
		godis := setup(t)
		value, _ := godis.Keys()
		assert.Equal(t, []string{}, value)
	})

	t.Run("Array reply", func(t *testing.T) {
		godis := setup(t)
		godis.Set("name", "raymond")
		godis.Set("age", 8)
		godis.Set("hello", "hello")
		value, _ := godis.Keys()
		assert.ElementsMatch(t, []string{"hello", "age", "name"}, value)
	})
}
