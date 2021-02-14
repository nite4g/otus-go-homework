package hw04_lru_cache

type Key string

type Cache interface {
	Clear()
	// Set(key Key, item interface{}) bool // Добавить значение в кэш по ключу.
	Set(key Key, item int) bool      // Добавить значение в кэш по ключу.
	Get(key Key) (interface{}, bool) // Получить значение из кэша по ключу.
	Len() int                        // cache length
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   string
	value interface{}
}

// func (c *lruCache) Set(key Key, value interface{}) bool {
func (c *lruCache) Set(key Key, value int) bool {
	ci := cacheItem{key: string(key), value: value}

	if len(c.items) == 0 {
		c.items[key] = c.queue.PushFront(ci) // передаваться должен cacheItem
	} else if val, found := c.items[key]; found {
		val.Value = ci
		c.queue.MoveToFront(val)
		c.items[key] = val
		return true
	} else {
		c.items[key] = c.queue.PushFront(ci)
	}

	if c.queue.Len() > c.capacity {
		// как я люблю  type assertions в go :)
		k := c.queue.Back().Value.(cacheItem).key
		delete(c.items, Key(k))
		c.queue.Remove(c.queue.Back())
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, found := c.items[key]; found {
		c.queue.MoveToFront(item)
		c.items[key] = item
		result := item.Value.(cacheItem).value
		return result, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	// присваиваем новые объекты, старые удалит GC
	c.items = make(map[Key]*ListItem)
	c.queue = NewList()
}

func (c *lruCache) Len() int {
	return c.queue.Len()
}
func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func NewList() List {
	return new(list) // синтакический сахар для &list
}
