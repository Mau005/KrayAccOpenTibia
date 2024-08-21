package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Mau005/KrayAcc/controller"
	"github.com/Mau005/KrayAcc/db"
	"github.com/Mau005/KrayAcc/models"
)

type HandlerOldSession struct{}

// Define your handlers
func (HO *HandlerOldSession) CacheInfoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	fmt.Println(string(body))
	var playersonline int
	db.DB.Raw("SELECT COUNT(*) FROM players_online").Scan(&playersonline)

	response := map[string]interface{}{
		"playersonline":        playersonline,
		"twitchstreams":        0,
		"twitchviewer":         0,
		"gamingyoutubestreams": 0,
		"gamingyoutubeviewer":  0,
	}
	HO.RespondJSON(w, response)
}

func (HO *HandlerOldSession) EventScheduleHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	fmt.Println(string(body))
	// Placeholder XML parsing, replace with your own logic
	eventList := []map[string]interface{}{} // Populate this with actual data

	response := map[string]interface{}{
		"eventlist":           eventList,
		"lastupdatetimestamp": time.Now().Unix(),
	}
	HO.RespondJSON(w, response)
}

func (HO *HandlerOldSession) BoostedCreatureHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	fmt.Println(string(body))
	var boostedCreature struct {
		RaceID int
	}
	if err := db.DB.Raw("SELECT * FROM boosted_creature").Scan(&boostedCreature).Error; err != nil {
		HO.RespondError(w, "Error fetching boosted creature data.")
		return
	}
	response := map[string]interface{}{
		"boostedcreature": true,
		"raceid":          boostedCreature.RaceID,
	}
	HO.RespondJSON(w, response)
}

func (HO *HandlerOldSession) loginHandler(answerExpected models.AnswerExpected, w http.ResponseWriter) (err error) {

	var accountCtl controller.AccountController

	account, err := accountCtl.GetAccountEmail(answerExpected.Email)
	if err != nil {
		log.Println("error get account", err)
		return
	}

	var apiCtl controller.ApiController
	passSha := apiCtl.ConvertSha1(answerExpected.Password)
	if passSha != account.Password {
		log.Println("error equals password", err)
		err = errors.New("email or password is not correct")
		return
	}

	var playerCtl controller.PlayerController
	player := playerCtl.GetPlayerWithAccountID(account.ID)

	var session models.ClientSession
	session.IsPremium = true
	session.LastLoginTime = uint32(time.Now().Unix())
	session.PremiumUntil = uint64(time.Now().Unix())
	session.OptionTracking = false
	session.SessionKey = fmt.Sprintf("%s\n%s\n%s\n%d", answerExpected.Email, answerExpected.Password, answerExpected.Token, time.Now().Add(30*time.Minute).Unix())
	session.Status = "active"
	session.IsReturner = false
	session.ShowRewardNews = false

	var world models.ClientWorld
	world.AntiCheatProtection = false
	world.ExternalAddRessUnProtected = "127.0.0.1"
	world.ExternalAddress = "127.0.0.1"
	world.ExternalAddressProtected = "127.0.0.1"
	world.ExternalPort = 7171
	world.ID = 0
	world.PvpType = 1
	world.Location = "CL"
	world.Name = "TheLastRookgard"
	world.PreviewState = 0
	world.ExternalPortProtected = 7172
	world.ExternalPortUnprotected = 7171
	world.CurrentTournamentPhase = 2

	var playdata models.PlayData
	playdata.World = append(playdata.World, world)
	playdata.Characters = apiCtl.PreparingCharacter(player)

	var responseData models.ResponseData
	responseData.PlayData = playdata
	responseData.Session = session

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println("error encode response", err)
		return
	}
	return
}

func (HO *HandlerOldSession) PreparingHanlderClient(w http.ResponseWriter, r *http.Request) {
	var answer models.AnswerExpected

	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		log.Println("error decode body", err)
	}

	switch answer.Type {
	case "login":
		HO.loginHandler(answer, w)

	}
}

// Helper functions
func (HO *HandlerOldSession) RespondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

func (HO *HandlerOldSession) RespondError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusConflict)
	HO.RespondJSON(w, map[string]interface{}{
		"errorCode":    3,
		"errorMessage": msg,
	})
}
