package handler

import (
	"fmt"
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
	claim, _ := context.Get(r, utils.CtxClaim).(models.Claim)
	Authentication := false
	if claim.TypeAccess > 0 {
		Authentication = true
	}

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

	type Home struct {
		Players        []models.Players
		UrlOutfitsView string
		NewsTicket     []ResponseTicket
		Authentication bool
	}
	var newsTicketCtl controller.NewsTickerController

	var responseTicket []ResponseTicket

	tickets, _ := newsTicketCtl.GetTickerLimited(utils.LimitRecordFive)
	for _, value := range tickets {
		fmt.Println(value.PlayersID)
		responseTicket = append(responseTicket, ResponseTicket{
			Icon:        newsTicketCtl.GetIconID(value.IconID),
			Ticket:      value.Ticket,
			IDCharacter: value.PlayersID,
			ByCharacter: value.Player.Name,
		})
	}
	var playerCtl controller.PlayerController
	players := playerCtl.GetPlayerLimits(5)
	response := Home{
		Players:        players,
		UrlOutfitsView: config.VarEnviroment.ServerWeb.UrlOutfitsView,
		NewsTicket:     responseTicket,
		Authentication: Authentication,
	}
	err = templ.Execute(w, response)
	if err != nil {
		log.Println("error execute template", err)
		return
	}
}
