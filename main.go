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
	db := driver.InitDatabase()
	repo := Repository.LoadDB(db)
	srv := service.LoadService(repo)
	r := api.NewGinHandler(srv)
	defer driver.CloseDatabase(db)
	log.Fatal(r.Run("localhost:8080"))
}
