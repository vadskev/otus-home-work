package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mutex    sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	var nItem cacheItem
	nItem.key = key
	nItem.value = value

	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	if cachedValue, isExistCache := cache.items[key]; isExistCache {
		cachedValue.Value = value
		cache.queue.MoveToFront(cachedValue)
		return true
	}

	if cache.capacity == cache.queue.Len() {
		delete(cache.items, cache.queue.Back().Value.(cacheItem).key)
		cache.queue.Remove(cache.queue.Back())
	}

	cache.items[key] = cache.queue.PushFront(nItem)

	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	if cachedValue, isExistCache := cache.items[key]; isExistCache {
		cache.queue.MoveToFront(cachedValue)
		return cachedValue.Value.(cacheItem).value, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
