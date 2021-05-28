package memauth

import "github.com/patrickmn/go-cache"

type MemAuth struct {
	Cache *cache.Cache
}
