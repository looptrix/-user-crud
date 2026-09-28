package cache

import "sync"

// Cache — простой потокобезопасный cache.
// Данные хранятся в map в памяти приложения.
type Cache struct {
	mu   sync.RWMutex
	data map[uint]interface{}
}

// New создаёт новый cache.
func New() *Cache {
	return &Cache{
		data: make(map[uint]interface{}),
	}
}

// Set сохраняет значение по ключу.
func (c *Cache) Set(key uint, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

// Get получает значение по ключу.
func (c *Cache) Get(key uint) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.data[key]
	return value, ok
}

// Delete удаляет значение из cache.
func (c *Cache) Delete(key uint) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}
