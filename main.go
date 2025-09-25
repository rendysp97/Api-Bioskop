package main

import (
	"github.com/rendysp97/Api-Bioskop/database"
	"github.com/rendysp97/Api-Bioskop/routers"
)

func main() {
	var PORT = ":8080"

	database.ConnectDB()

	routers.StartServer().Run(PORT)
}
