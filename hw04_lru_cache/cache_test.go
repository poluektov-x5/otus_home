package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		cache := NewCache(3)

		cache.Set("a", 100)
		cache.Set("b", 200)
		cache.Set("c", 300)
		cache.Set("d", 400)

		value, isFound := cache.Get("a")
		require.False(t, isFound)
		require.Nil(t, value)

		value, isFound = cache.Get("b")
		require.True(t, isFound)
		require.Equal(t, 200, value)

		value, isFound = cache.Get("c")
		require.True(t, isFound)
		require.Equal(t, 300, value)

		value, isFound = cache.Get("d")
		require.True(t, isFound)
		require.Equal(t, 400, value)
	})

	t.Run("complex purge logic", func(t *testing.T) {
		cache := NewCache(3)

		// Добавление 3-х значений.
		wasInCache := cache.Set("a", 100) // a
		require.False(t, wasInCache)

		wasInCache = cache.Set("b", 200) // b a
		require.False(t, wasInCache)

		wasInCache = cache.Set("c", 300) // c b a
		require.False(t, wasInCache)

		// Перемешивание значений кеша.
		value, isFound := cache.Get("a") // a c b
		require.True(t, isFound)
		require.Equal(t, 100, value)

		value, isFound = cache.Get("c") // c a b
		require.True(t, isFound)
		require.Equal(t, 300, value)

		wasInCache = cache.Set("b", 20) // b c a
		require.True(t, wasInCache)

		wasInCache = cache.Set("a", 1) // a b c
		require.True(t, wasInCache)

		// Добавление 4-го значения.
		wasInCache = cache.Set("d", 400) // d a b
		require.False(t, wasInCache)

		// Попытка получения значения, которое уже было удалено из кеша.
		value, isFound = cache.Get("c")
		require.False(t, isFound)
		require.Nil(t, value)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
