package cache

import (
	"sync"
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
)

type _Item struct {
	Value      any
	Expiration int64
	Created    time.Time
}

type Cache struct {
	sync.RWMutex
	Storage map[string]_Item

	defaultExpiration time.Duration
	cleanupInterval   time.Duration
}

func NewCache(defaultExpiration, cleanupInterval time.Duration) domain.Cache {
	cache := &Cache{
		Storage:           make(map[string]_Item),
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	if cleanupInterval > 0 {
		cache.StartGC()
	}

	return cache
}

func (cache *Cache) Set(key string, value any, duration time.Duration) {
	var expiration int64

	if duration == 0 {
		duration = cache.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	cache.Lock()
	defer cache.Unlock()

	cache.Storage[key] = _Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}

}

func (cache *Cache) Get(key string) (any, error) {
	cache.RLock()
	defer cache.RUnlock()

	item, found := cache.Storage[key]
	if !found {
		return nil, domain.ErrorNotFound
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, domain.ErrorNotFound
		}
	}

	return item.Value, nil
}

func (cache *Cache) Del(key string) {
	cache.Lock()
	defer cache.Unlock()

	delete(cache.Storage, key)
}
