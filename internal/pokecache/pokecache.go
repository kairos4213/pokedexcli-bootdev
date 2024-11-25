package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (ca *Cache) Add(key string, val []byte) {
	ca.mu.Lock()
	defer ca.mu.Unlock()

	ca.entries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (ca *Cache) Get(key string) ([]byte, bool) {
	ca.mu.RLock()
	defer ca.mu.RUnlock()

	entry, ok := ca.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (ca *Cache) reapLoop(interval time.Duration) {
	ca.mu.Lock()
	defer ca.mu.Unlock()

	ticker := time.NewTicker(interval)
	for range ticker.C {
		expiryTime := time.Now().Add(-interval)
		for key, entry := range ca.entries {
			if entry.createdAt.Before(expiryTime) {
				delete(ca.entries, key)
			}
		}
	}
}
