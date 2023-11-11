package main

import (
	"containers/app"
	"containers/service"
	"containers/docker"
)

func main() {
	docker.StartDockerClient()
	go service.AutoScaling("search")
	app.StartRoute()
}