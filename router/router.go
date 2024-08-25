package router

import (
	"net/http"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/handler"
	"github.com/Mau005/KrayAccOpenTibia/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	var NewsTickerHandler handler.NewsTicketHandler
	var handlerAccount handler.AccountHandler

	fs := http.FileServer(http.Dir("./www"))

	r := mux.NewRouter()
	r.Use(middleware.AuthPathPublicMiddleware)
	r.HandleFunc("/get_news_ticket", NewsTickerHandler.GetTicket).Methods("GET") //API PUBLIC
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(middleware.AuthMiddleware)
	if !config.VarEnviroment.ServerWeb.ApiMode {
		r.PathPrefix("/www/").Handler(http.StripPrefix("/www/", fs))

		var homeHandler handler.HomeHandler
		r.HandleFunc("/", homeHandler.GetHome).Methods("GET") //Public
		r.HandleFunc("/create_account", handlerAccount.CreateAccount).Methods("POST")
		s.HandleFunc("/create_character", handlerAccount.CreateCharacter).Methods("POST")

		//Not Found
		// r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 	bodyResponde, err := io.ReadAll(r.Body)
		// 	if err != nil {
		// 		log.Println("error io read", err)
		// 		return
		// 	}
		// 	fmt.Println(string(bodyResponde))
		// })

		r.HandleFunc("/login", handlerAccount.Authentication).Methods("POST")
		r.HandleFunc("/logout", handlerAccount.Desconnected).Methods("GET")

	}

	// Router client connections
	ctl := r.PathPrefix("/client").Subrouter()
	ctl.Use(middleware.CommonMiddleware)

	var handlerClientConnect handler.HandlerClientConnect
	ctl.HandleFunc("/cacheinfo", handlerClientConnect.CacheInfoHandler)
	ctl.HandleFunc("/eventschedule", handlerClientConnect.EventScheduleHandler)
	ctl.HandleFunc("/boostedcreature", handlerClientConnect.BoostedCreatureHandler)
	ctl.HandleFunc("/login", handlerClientConnect.PreparingHanlderClient)

	s.HandleFunc("/create_news_ticket", NewsTickerHandler.CreateTicket).Methods("POST")

	return r
}
