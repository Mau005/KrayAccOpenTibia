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
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./www"))
	if !config.VarEnviroment.ServerWeb.ApiMode {
		//WEB Active!
		r.Use(middleware.AuthPathPublicMiddleware)
		r.PathPrefix("/www/").Handler(http.StripPrefix("/www/", fs))
		r.HandleFunc("/get_news_ticket", NewsTickerHandler.GetTicket).Methods("GET") //API PUBLIC

		var homeHandler handler.HomeHandler
		r.HandleFunc("/", homeHandler.GetHome).Methods("GET") //Public

		var whoPlayer handler.WhoOnlineHandler
		r.HandleFunc("/who_online", whoPlayer.GetViewPlayer).Methods("GET")

		var killerhandler handler.PlayerDeathHandler
		r.HandleFunc("/last_death", killerhandler.GetViewPlayerDeath).Methods("GET")

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
		r.HandleFunc("/create_account", handlerAccount.CreateAccount).Methods("POST")

		//SecurityPath
		s := r.PathPrefix("/auth").Subrouter()
		s.Use(middleware.AuthMiddleware)
		s.HandleFunc("/create_character", handlerAccount.CreateCharacter).Methods("POST")
		s.HandleFunc("/create_news_ticket", NewsTickerHandler.CreateTicket).Methods("POST")
	} else {
		//APIMODE
		api := r.PathPrefix("/api").Subrouter()
		api.Use(middleware.AuthPoolConnection)
	}

	// Router client connections
	ctl := r.PathPrefix("/client").Subrouter()
	ctl.Use(middleware.CommonMiddleware)

	var handlerClientConnect handler.HandlerClientConnect
	ctl.HandleFunc("/cacheinfo", handlerClientConnect.CacheInfoHandler)
	ctl.HandleFunc("/eventschedule", handlerClientConnect.EventScheduleHandler)
	ctl.HandleFunc("/boostedcreature", handlerClientConnect.BoostedCreatureHandler)
	ctl.HandleFunc("/login", handlerClientConnect.PreparingHanlderClient)

	return r
}
