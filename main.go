package main

import (
	"log"
	"net/http"

	"github.com/Mau005/KrayAcc/db"
	"github.com/Mau005/KrayAcc/router"
)

func main() {

	db.ConnectionMysql("root", "12345", "127.0.0.1", "tfs", 3306, true)

	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":7575", r))
}
