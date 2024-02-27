package cache

import "time"

func (cache *Cache) StartGC() {
	go cache.GC()
}

func (cache *Cache) GC() {
	for {
		<-time.After(cache.cleanupInterval)
		if cache.Storage == nil {
			return
		}

		if keys := cache.expiredKeys(); len(keys) != 0 {
			cache.clearItems(keys)
		}
	}
}

func (cache *Cache) expiredKeys() (keys []string) {
	cache.RLock()
	defer cache.RUnlock()

	for k, i := range cache.Storage {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

func (cache *Cache) clearItems(keys []string) {
	cache.Lock()
	defer cache.Unlock()

	for _, k := range keys {
		delete(cache.Storage, k)
	}
}
