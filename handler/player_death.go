package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/gorilla/context"
)

type PlayerDeathHandler struct{}

func (pdh *PlayerDeathHandler) GetViewPlayerDeath(w http.ResponseWriter, r *http.Request) {
	navWeb, _ := context.Get(r, utils.CtxNavWeb).(models.NavWeb)

	templ, err := template.New("death_player.html").ParseFiles("www/death_player.html")
	if err != nil {
		log.Println("error create template", err)
		return
	}
	var Layouthandler Layouthandler
	ConditionalLayout := models.NewLayoutDefault()
	ConditionalLayout.LastDeath = true
	err = templ.Execute(w, Layouthandler.Generatelayout(navWeb, ConditionalLayout))
	if err != nil {
		log.Println("error execute template", err)
		return
	}

}
