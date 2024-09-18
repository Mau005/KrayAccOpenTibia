package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type HighScorehandler struct{}

func (hs *HighScorehandler) ViewHighScore(w http.ResponseWriter, r *http.Request) {
	navWeb, _ := context.Get(r, utils.CtxNavWeb).(models.NavWeb)

	templ, err := template.New("highscore.html").ParseFiles("www/highscore.html")
	if err != nil {
		log.Println("error create template", err)
		return
	}
	var Layouthandler Layouthandler
	err = templ.Execute(w, Layouthandler.Generatelayout(navWeb, models.SolicitudeLayout{News: true,
		Login:        true,
		ServerStatus: true,
		Discord:      true,
		TopPlayers:   true,
		Rates:        true,
		HighScore:    true}))
	if err != nil {
		log.Println("error execute template", err)
		return
	}
}

func (hs *HighScorehandler) GetHighScoreHandler(w http.ResponseWriter, r *http.Request) {
	var exepCtl controller.ExceptionController
	vars := mux.Vars(r)
	world := vars["world"]
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println(err)
		exepCtl.Exeption(err.Error(), http.StatusConflict, w)
		return
	}

	var request struct {
		ID int
	}
	var response struct {
		World    string
		Category string
		Players  []models.Players
	}
	request.ID = int(id)
	for _, pool := range config.Global.PoolServer {
		if pool.World.Name == world {
			if pool.IpWebApi == "" {
				var playerCtl controller.PlayerController
				response.Players = playerCtl.GetHighScore(request.ID)
				response.World = pool.World.Name
				response.Category = playerCtl.IndexHighScore(request.ID)
				if err := json.NewEncoder(w).Encode(&response); err != nil {
					log.Println(err)
					exepCtl.Exeption(err.Error(), http.StatusConflict, w)
				}
			} else {
				structMarshal, err := json.Marshal(&request)
				if err != nil {
					log.Println(err)
					exepCtl.Exeption(err.Error(), http.StatusConflict, w)
					return
				}
				req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlGetHighScore), bytes.NewBuffer(structMarshal))
				if err != nil {
					log.Println(err)
					exepCtl.Exeption(err.Error(), http.StatusConflict, w)
					return
				}

				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))

				client := http.Client{}
				body, err := client.Do(req)
				if err != nil {
					log.Println(err)
					exepCtl.Exeption(err.Error(), http.StatusConflict, w)
					return
				}
				if body.StatusCode != http.StatusOK {
					log.Println(err)
					exepCtl.Exeption("error", http.StatusConflict, w)
					return
				}

				if err := json.NewDecoder(body.Body).Decode(&response.Players); err != nil {
					log.Println(err)
					exepCtl.Exeption(err.Error(), http.StatusConflict, w)
					return
				}
				response.World = pool.World.Name
				if err := json.NewEncoder(w).Encode(&response); err != nil {
					log.Println(err)
					exepCtl.Exeption(err.Error(), http.StatusConflict, w)
					return
				}
			}
			break
		}
	}
}
