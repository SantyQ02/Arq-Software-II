package cacheLocal

import (
	"fmt"
	"github.com/karlseguin/ccache/v3"
	"mvc-go/dto"

	"time"
)

var (
	cache *ccache.Cache[[]string]
)

func init() {
	cache = ccache.New(ccache.Configure[[]string]())
}

func Get(key string) []dto.Hotel {
	item := cache.Get(key)
	if item == nil {
		fmt.Println("Error getting item from cacheLocal")
		return nil
	} else {
		return item.Value()
	}
}

func Set(key string, value []dto.Hotel) {
	cache.Set(key, value)
}

func SetWithExpiration(key string, value []dto.Hotel, expiration int) {
	duration := time.Second * time.Duration(expiration)
	cache.Set(key, value, duration)
}
