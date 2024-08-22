package router

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Mau005/KrayAccOpenTibia/handler"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	fs := http.FileServer(http.Dir("./www"))

	r := mux.NewRouter()
	r.PathPrefix("/www/").Handler(http.StripPrefix("/www/", fs))
	//Not Found
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyResponde, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("error io read", err)
			return
		}
		fmt.Println(string(bodyResponde))
	})

	//HandlerOldSession
	var HanlderOld handler.HandlerOldSession
	//r.HandleFunc("/", HanlderOld.LoginHandler)
	r.HandleFunc("/cacheinfo", HanlderOld.CacheInfoHandler)
	r.HandleFunc("/eventschedule", HanlderOld.EventScheduleHandler)
	r.HandleFunc("/boostedcreature", HanlderOld.BoostedCreatureHandler)
	r.HandleFunc("/login", HanlderOld.PreparingHanlderClient)

	var homeHandler handler.HomeHandler
	r.HandleFunc("/", homeHandler.GetHome).Methods("GET")

	return r
}
