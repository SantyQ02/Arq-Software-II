package cache

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

var (
	cacheClient *memcache.Client
)

func InitCache() {
	cacheClient = memcache.New("memcached:11211")
}

func Get(key string) []byte {
	item, err := cacheClient.Get(key)
	if err != nil {
		fmt.Println("Error getting item from cache", err)
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

func Test() {
	fmt.Println("Cache testing...")
	InitCache()

	// Prueba Set() y Get()
	key := "myKey"
	value := []byte("Hello, World!")
	Set(key, value)
	fmt.Printf("cache key %s: %s\n", key, string(Get(key)))

	// Prueba de tiempo de vida de la caché
	expiration := 10
	keyWithExpiration := "keyWithExpiration"
	valueWithExpiration := []byte("10 sec expiration")
	SetWithExpiration(keyWithExpiration, valueWithExpiration, expiration)
	fmt.Printf("cache value %s: %s\n", keyWithExpiration, string(Get(keyWithExpiration)))

	// Esperar 11 segundos para que expire el valor con expiración
	time.Sleep(11 * time.Second)
	cachedValue := Get(keyWithExpiration)
	if cachedValue == nil {
		fmt.Printf("key value expirated\n", keyWithExpiration)
	} else {
		fmt.Printf("key value after expiration: %s\n", keyWithExpiration, string(cachedValue))
	}
}
