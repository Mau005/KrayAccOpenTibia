package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/gorilla/context"
)

type WhoOnlineHandler struct{}

func (woh *WhoOnlineHandler) GetViewPlayer(w http.ResponseWriter, r *http.Request) {
	navWeb, _ := context.Get(r, utils.CtxNavWeb).(models.NavWeb)

	var Layouthandler Layouthandler
	ConditionalLayout := models.NewLayoutDefault()
	ConditionalLayout.WhoIsOnline = true
	templ, err := template.New("player_online.html").ParseFiles("www/player_online.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = templ.Execute(w, Layouthandler.Generatelayout(navWeb, ConditionalLayout))
	if err != nil {
		log.Println(err)
	}
}
