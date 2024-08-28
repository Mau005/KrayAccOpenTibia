package handler

import (
	"net/http"
	"text/template"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/gorilla/context"
)

type WhoOnlineHandler struct{}

func (woh *WhoOnlineHandler) GetViewPlayer(w http.ResponseWriter, r *http.Request) {
	navWeb, _ := context.Get(r, utils.CtxNavWeb).(models.NavWeb)
	var playerCtl controller.PlayerController
	var response struct {
		UrlOutfitsView string
		Players        []models.Players
		NavWeb         models.NavWeb
		ServerStatus   models.ServerStatus
		TopPlayers     []models.Players
	}

	response.UrlOutfitsView = config.VarEnviroment.ServerWeb.UrlOutfitsView
	response.Players = playerCtl.GetPlayerOnline()
	response.NavWeb = navWeb
	response.ServerStatus = controller.TempData.ServerStatus
	response.TopPlayers = playerCtl.GetPlayerLimits(utils.LimitRecordFive)

	templ, err := template.New("player_online.html").ParseFiles("www/player_online.html")
	if err != nil {
		utils.Error("error create tempalte", err.Error())
		return
	}
	err = templ.Execute(w, response)
	if err != nil {
		utils.Error("error execute template", err.Error())
	}
}
