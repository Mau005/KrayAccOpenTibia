package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Mau005/KrayAccOpenTibia/components"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type PlayerHandler struct {
	PlayCtl controller.PlayerController
	PoolCtl controller.PoolConnectionController
}

func (pdh *PlayerHandler) GetViewPlayerDeath(w http.ResponseWriter, r *http.Request) {
	navWeb, _ := context.Get(r, utils.CtxNavWeb).(models.NavWeb)

	templ, err := template.New("index.html").ParseFiles("www/index.html")
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

func (pdh *PlayerHandler) getCharacterInternal(navWeb models.NavWeb, nameCharacter string, w http.ResponseWriter) {
	player, errPlayer := pdh.PoolCtl.GetPlayerPoolConnection(nameCharacter)
	if errPlayer != nil {
		log.Println(errPlayer)
	}

	templ, err := template.New("index.html").ParseFiles("www/index.html")
	if err != nil {
		log.Println("error create template", err)
		return
	}
	var Layouthandler Layouthandler
	ConditionalLayout := models.NewLayoutDefault()
	ConditionalLayout.LastDeath = true
	layout := Layouthandler.Generatelayout(navWeb, ConditionalLayout)

	layout.Component = components.CreateGetPlayer(player)
	if errPlayer != nil {
		layout.Component = `

  <div class="text-center">
    <h1 class="display-4">Oops 😓</h1>
    <p class="lead">El personaje que buscaste no existe o fue deletiado.</p>
    <a href="/" class="btn btn-outline-light mt-3">Volver al inicio</a>
  </div>

		`
	}
	err = templ.Execute(w, layout)
	if err != nil {
		log.Println("error execute template", err)
		return
	}
}

func (pdh *PlayerHandler) GetCharacterPOST(w http.ResponseWriter, r *http.Request) {
	navWeb, _ := context.Get(r, utils.CtxNavWeb).(models.NavWeb)
	pdh.getCharacterInternal(navWeb, r.FormValue("search"), w)

}
func (pdh *PlayerHandler) GetCharacter(w http.ResponseWriter, r *http.Request) {
	navWeb, _ := context.Get(r, utils.CtxNavWeb).(models.NavWeb)

	content := mux.Vars(r)
	pdh.getCharacterInternal(navWeb, content["name"], w)

}
