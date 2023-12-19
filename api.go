package main

import (
	"Todo/db"
	"Todo/routers"
)

func main() {
	db.CreateCon()
	e := routers.Init()
	e.Logger.Fatal(e.Start(":1323"))
}