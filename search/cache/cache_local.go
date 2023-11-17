package cacheLocal

import (
	"github.com/karlseguin/ccache/v3"
	"mvc-go/dto"
)

var (
	Cache *ccache.Cache[[]dto.Hotel]
)

func StartLocalCache() {
	Cache = ccache.New(ccache.Configure[[]dto.Hotel]())
}

