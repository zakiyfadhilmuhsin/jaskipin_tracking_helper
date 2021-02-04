package main

import (
	"jaskipin-front/db"
	"jaskipin-front/routes"
	"os"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
