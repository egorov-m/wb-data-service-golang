package cache

import (
	"fmt"
	"time"
	"wb-data-service-golang/wb-data-service/internal/domain"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
)

const (
	defaultUserExpiration = 24 * time.Hour
)

func getUserKey(entity core.User) string {
	return fmt.Sprintf("user-%d", entity.Id)
}

type CacheAdapter interface {
	Set(core.User)
	Get(core.User) (core.User, error)
	Del(core.User)
}

type _CacheAdapter struct {
	Cache domain.Cache
}

func NewCacheAdapter(cache domain.Cache) CacheAdapter {
	return &_CacheAdapter{Cache: cache}
}

func (adapter *_CacheAdapter) Del(entity core.User) {
	adapter.Cache.Del(
		getUserKey(entity),
	)
}

func (adapter *_CacheAdapter) Get(entity core.User) (core.User, error) {
	user, err := adapter.Cache.Get(getUserKey(entity))
	if err != nil {
		return core.User{}, err
	}

	entity, ok := user.(core.User)
	if !ok {
		return core.User{}, domain.ErrorInvalidType
	}

	return entity, nil
}

func (adapter *_CacheAdapter) Set(entity core.User) {
	adapter.Cache.Set(
		getUserKey(entity),
		entity,
		defaultUserExpiration,
	)
}
