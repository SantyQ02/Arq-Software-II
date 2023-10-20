package main

import (
	"mvc-go/app"
	"mvc-go/solr"
)

func main() {
	solr.StartSolr()
	app.StartRoute()
}
