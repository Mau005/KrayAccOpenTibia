package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
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
	type ResponseTicket struct {
		Icon        string
		Ticket      string
		ByCharacter string
		IDCharacter int
	}

	var newsTicketCtl controller.NewsTickerController
	var responseTicket []ResponseTicket

	tickets, _ := newsTicketCtl.GetTickerLimited(utils.LimitRecordFive)
	for _, value := range tickets {
		responseTicket = append(responseTicket, ResponseTicket{
			Icon:        newsTicketCtl.GetIconID(value.IconID),
			Ticket:      value.Ticket,
			IDCharacter: value.PlayersID,
			ByCharacter: value.Player.Name,
		})
	}
	var playerCtl controller.PlayerController

	type Home struct {
		Players        []models.Players
		UrlOutfitsView string
		NewsTicket     []ResponseTicket
		NavWeb         models.NavWeb
		ServerStatus   models.ServerStatus
	}
	response := Home{
		Players:        playerCtl.GetPlayerLimits(5),
		UrlOutfitsView: config.VarEnviroment.ServerWeb.UrlOutfitsView,
		NewsTicket:     responseTicket,
		NavWeb:         navWeb,
		ServerStatus:   controller.TempData.ServerStatus,
	}
	err = templ.Execute(w, response)
	if err != nil {
		log.Println("error execute template", err)
		return
	}
}
