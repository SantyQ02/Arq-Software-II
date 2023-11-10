package main

import (
	"containers/app"
	"containers/utils"
)

func main() {
	go utils.AutoScaling()
	app.StartRoute()
}