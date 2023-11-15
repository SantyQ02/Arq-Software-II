package client

import (
	"time"
	"mvc-go/dto"

	cc "mvc-go/cache"

	log "github.com/sirupsen/logrus"
)

type cacheClient struct{}

type cacheClientInterface interface {
	Get(key string) []dto.Hotel
	SetWithExpiration(key string, value []dto.Hotel, expiration int)
}

var (
	CacheClient cacheClientInterface
)

func init() {
	CacheClient = &cacheClient{}
}

func (c *cacheClient) Get(key string) []dto.Hotel {
	item := cc.Cache.Get(key)
	if item == nil {
		log.Error("Error getting item from local cache")
		return nil
	} 
	if item.Expired() {
		log.Error("Item is expired")
		return nil
	}
	log.Info("Geting Item from local cache successfully")
	return item.Value()
}

func (c *cacheClient) SetWithExpiration(key string, value []dto.Hotel, expiration int) {
	duration := time.Second * time.Duration(expiration)
	cc.Cache.Set(key, value, duration)
}