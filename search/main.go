package main

import (
	"mvc-go/app"
	"mvc-go/solr"
	"mvc-go/utils/queue"
)

func main() {
	solr.StartSolr()
	go queue.Consumer()
	app.StartRoute()
}
