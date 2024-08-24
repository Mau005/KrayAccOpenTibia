package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
)

type HomeHandler struct{}

func (hh *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	templ, err := template.New("index.html").ParseFiles("www/index.html")
	if err != nil {
		log.Println("error create template", err)
		return
	}
	type ResponseTicket struct {
		Icon        string
		Ticket      string
		ByCharacter string
	}

	type Home struct {
		Players        []models.Player
		UrlOutfitsView string
		NewsTicket     []ResponseTicket
	}
	var newsTicketCtl controller.NewsTickerController

	var responseTicket []ResponseTicket

	tickets, _ := newsTicketCtl.GetTickerLimited(5)

	for _, value := range tickets {
		responseTicket = append(responseTicket, ResponseTicket{
			Icon:        value.Icon,
			Ticket:      value.Ticket,
			ByCharacter: value.Player.Name,
		})
	}
	response := Home{
		Players:        []models.Player{},
		UrlOutfitsView: config.VarEnviroment.ServerWeb.UrlOutfitsView,
		NewsTicket:     responseTicket,
	}
	err = templ.Execute(w, response)
	if err != nil {
		log.Println("error execute template", err)
		return
	}
}
