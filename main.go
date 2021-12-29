package main

import (
	"PhotoDiary/Repository"
	"PhotoDiary/api"
	"PhotoDiary/driver"
	"PhotoDiary/service"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	repo := Repository.LoadDB(driver.InitDatabase())
	srv := service.LoadService(repo)
	r := api.NewGinHandler(srv)

	log.Fatal(r.Run("localhost:8080"))
}
