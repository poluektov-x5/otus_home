package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

// Set - Добавление значения в кеш по ключу.
func (c *lruCache) Set(key Key, value interface{}) bool {
	// Получение элемента из словаря.
	item := c.items[key]

	// Если элемент уже в кеше.
	if item != nil {
		// Обновление значения элемента.
		item.Value = cacheItem{key, value}
		c.queue.MoveToFront(item)

		return true
	}

	// Добавление нового значения в начало списка.
	item = c.queue.PushFront(cacheItem{key, value})

	// Добавление элемента в словарь.
	c.items[key] = item

	// Проверка размера кеша.
	c.checkCacheSize()

	return false
}

// Get - Получение значения из кеша по ключу.
func (c *lruCache) Get(key Key) (interface{}, bool) {
	// Получение элемента из словаря.
	item := c.items[key]

	if item == nil {
		return nil, false
	}

	c.queue.MoveToFront(item)

	if cacheItem, ok := item.Value.(cacheItem); ok {
		return cacheItem.value, true
	}

	return nil, false
}

// Clear - Очистка кеша.
func (c *lruCache) Clear() {
	// Удаление всех элементов из списка.
	for c.queue.Len() > 0 {
		item := c.queue.Front()
		c.queue.Remove(item)
	}

	// Очистка словаря.
	c.items = make(map[Key]*ListItem, c.capacity)
}

// Проверка и ограничение размера кеша.
func (c *lruCache) checkCacheSize() {
	// Если размер кеша не превышен.
	if c.queue.Len() <= c.capacity {
		return
	}

	// Получение последнего элемента из списка.
	lastItem := c.queue.Back()

	// Удаление элемента из списка.
	c.queue.Remove(lastItem)

	// Удаление элемента из словаря.
	if cacheItem, ok := lastItem.Value.(cacheItem); ok {
		delete(c.items, cacheItem.key)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
