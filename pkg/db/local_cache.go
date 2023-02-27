package db

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	ProjectsCache *cache.Cache
	MerchantCache *cache.Cache
	UserCache     *cache.Cache
)

func InitLocalCache() {
	ProjectsCache = cache.New(time.Minute*10, time.Hour*1)
	MerchantCache = cache.New(time.Minute*10, time.Hour*1)
	UserCache = cache.New(time.Minute*10, time.Hour*1)
}
