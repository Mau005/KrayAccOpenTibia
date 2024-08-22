package main

import (
	"log"
	"net/http"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/router"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

func main() {

	err := config.Load("config.yml")
	if err != nil {
		utils.ErrorFatal(err.Error())
	}
	err = db.ConnectionMysql("root", "12345", "127.0.0.1", "tfs", 3306, true)
	if err != nil {
		utils.ErrorFatal(err.Error())
	}

	r := router.NewRouter()
	log.Fatalln(http.ListenAndServe(":7575", r))
}
