package main

import (
	"mvc-go/app"
	"mvc-go/solr"
	cc "mvc-go/cache"
	"mvc-go/utils/queue"
)

func main() {
	solr.StartSolr()
	cc.StartLocalCache()
	go queue.Consumer()
	app.StartRoute()
}
