package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type HandlerClientConnect struct{}

// Define your handlers
func (hcc *HandlerClientConnect) CacheInfoHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]interface{}{
		"playersonline":        controller.TempData.ServStatusTotal.Players.Online,
		"twitchstreams":        0,
		"twitchviewer":         0,
		"gamingyoutubestreams": 0,
		"gamingyoutubeviewer":  0,
	}
	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		utils.Error("error cache info client", err.Error())
	}
}

func (hcc *HandlerClientConnect) EventScheduleHandler(w http.ResponseWriter, r *http.Request) {
	// Placehcclder XML parsing, replace with your own logic
	eventList := []map[string]interface{}{} // Populate this with actual data

	response := map[string]interface{}{
		"eventlist":           eventList,
		"lastupdatetimestamp": time.Now().Unix(),
	}
	hcc.RespondJSON(w, response)
}

func (hcc *HandlerClientConnect) BoostedCreatureHandler(w http.ResponseWriter, r *http.Request) {
	var boostedCreature struct {
		RaceID int
	}
	response := map[string]interface{}{
		"boostedcreature": true,
		"raceid":          boostedCreature.RaceID,
	}
	hcc.RespondJSON(w, response)
}

func (hcc *HandlerClientConnect) loginHandler(answerExpected models.AnswerExpected, w http.ResponseWriter) (err error) {

	var PoolConnectionController controller.PoolConnectionController

	response, err := PoolConnectionController.CharacterLoginAccountPoolConnection(answerExpected)
	if err != nil {
		utils.Error(err.Error())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errorCode":    3,
			"errorMessage": "incorrect credentials",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&response); err != nil {
		utils.Error(err.Error())
		utils.Warn("error encode response", err.Error())
		return
	}
	return
}

func (hcc *HandlerClientConnect) PreparingHanlderClient(w http.ResponseWriter, r *http.Request) {
	var answer models.AnswerExpected
	err := json.NewDecoder(r.Body).Decode(&answer)

	if err != nil {
		log.Println("error decode body", err)
	}

	switch answer.Type {
	case "login":
		hcc.loginHandler(answer, w)

	case "cacheinfo":
		hcc.CacheInfoHandler(w, r)

	}
}

// Helper functions
func (hcc *HandlerClientConnect) RespondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

func (hcc *HandlerClientConnect) RespondError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"errorCode":    3,
		"errorMessage": msg,
	})
}
