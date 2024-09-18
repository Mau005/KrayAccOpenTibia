package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/gorilla/context"
)

type HomeHandler struct{}

func (hh *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	navWeb, _ := context.Get(r, utils.CtxNavWeb).(models.NavWeb)

	templ, err := template.New("index.html").ParseFiles("www/index.html")
	if err != nil {
		log.Println("error create template", err)
		return
	}
	var Layouthandler Layouthandler
	err = templ.Execute(w, Layouthandler.Generatelayout(navWeb, models.SolicitudeLayout{News: true, Discord: true, Login: true, ServerStatus: true, TopPlayers: true, Rates: true}))
	if err != nil {
		log.Println("error execute template", err)
		return
	}
}
