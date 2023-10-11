package cache

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	log "github.com/sirupsen/logrus"
)

var (
	cacheClient *memcache.Client
)

func init() {
	cacheClient = memcache.New("memcached:11211")
	if cacheClient == nil {
		fmt.Println("Error initializing Memcached client")
	}
	log.Info("Cache initialized correctly!")
}

func Get(key string) []byte {
	item, err := cacheClient.Get(key)
	if err != nil {
		fmt.Printf("Error getting item from cache for key %s: %v\n", key, err)
		return nil
	}
	if item == nil {
		fmt.Printf("Item not found in cache for key %s\n", key)
		return nil
	}
	return item.Value
}

func Set(key string, value []byte) {
	if err := cacheClient.Set(&memcache.Item{
		Key:   key,
		Value: value,
	}); err != nil {
		fmt.Println("Error setting item in cache", err)
	}
}

func SetWithExpiration(key string, value []byte, expiration int) {
	if err := cacheClient.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(expiration), // Aquí especificamos la expiración
	}); err != nil {
		fmt.Println("Error setting item in cache", err)
	}
}
