package router

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/handler"
	"github.com/Mau005/KrayAccOpenTibia/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	fs := http.FileServer(http.Dir("./www"))

	r := mux.NewRouter()
	if !config.VarEnviroment.ServerWeb.ApiMode {
		r.PathPrefix("/www/").Handler(http.StripPrefix("/www/", fs))

		var homeHandler handler.HomeHandler
		r.HandleFunc("/", homeHandler.GetHome).Methods("GET")
		//Not Found
		r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//check posible enter web
			bodyResponde, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("error io read", err)
				return
			}
			fmt.Println(string(bodyResponde))
		})

		var handlerAccount handler.AccountHandler
		r.HandleFunc("/login", handlerAccount.Authentication).Methods("POST")
	}

	// Router client connections
	ctl := r.PathPrefix("/client").Subrouter()
	ctl.Use(middleware.CommonMiddleware)

	var handlerClientConnect handler.HandlerClientConnect
	ctl.HandleFunc("/cacheinfo", handlerClientConnect.CacheInfoHandler)
	ctl.HandleFunc("/eventschedule", handlerClientConnect.EventScheduleHandler)
	ctl.HandleFunc("/boostedcreature", handlerClientConnect.BoostedCreatureHandler)
	ctl.HandleFunc("/login", handlerClientConnect.PreparingHanlderClient)

	s := r.PathPrefix("/auth").Subrouter()
	s.Use(middleware.AuthMiddleware)

	return r
}
