package main

import (
	"containers/app"
	"containers/service"
	"containers/docker"
)

func main() {
	docker.StartDockerClient()
	go service.AutoScaling("frontend")
	go service.AutoScaling("search")
	go service.AutoScaling("business")
	go service.AutoScaling("hotels")
	app.StartRoute()
}