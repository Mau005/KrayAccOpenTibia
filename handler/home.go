package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/models"
)

type HomeHandler struct{}

func (hh *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	templ, err := template.New("index.html").ParseFiles("www/index.html")
	if err != nil {
		log.Println("error create template", err)
		return
	}

	type home struct {
		Players        []models.Player
		UrlOutfitsView string
	}
	var player []models.Player
	db.DB.Find(&player)
	response := home{Players: player, UrlOutfitsView: config.VarEnviroment.ServerWeb.UrlOutfitsView}
	err = templ.Execute(w, response)
	if err != nil {
		log.Println("error execute template", err)
		return
	}
}
