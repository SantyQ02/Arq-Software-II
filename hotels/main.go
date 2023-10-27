package main

import (
	"mvc-go/app"
	"mvc-go/db"
	"fmt"
	"mvc-go/queue"
)

func main() {
	err:=db.Init_db()
	queue.StartQueue()
	app.StartRoute()

	defer db.Disconect_db()

	if err!=nil {
			fmt.Println("Cannot init db")
			fmt.Println(err)
			return;
		}
}
