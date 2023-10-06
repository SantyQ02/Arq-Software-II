package main

import (
	"mvc-go/app"
	"mvc-go/db"
	"fmt"
)

func main() {
	err:=db.Init_db()
	app.StartRoute()

	defer db.Disconect_db()

	if err!=nil {
			fmt.Println("Cannot init db")
			fmt.Println(err)
			return;
		}
}
