package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type HandlerClientConnect struct{}

// Define your handlers
func (hcc *HandlerClientConnect) CacheInfoHandler(w http.ResponseWriter, r *http.Request) {
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
	hcc.RespondJSON(w, response)
}

func (hcc *HandlerClientConnect) EventScheduleHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	fmt.Println(string(body))
	// Placehcclder XML parsing, replace with your own logic
	eventList := []map[string]interface{}{} // Populate this with actual data

	response := map[string]interface{}{
		"eventlist":           eventList,
		"lastupdatetimestamp": time.Now().Unix(),
	}
	hcc.RespondJSON(w, response)
}

func (hcc *HandlerClientConnect) BoostedCreatureHandler(w http.ResponseWriter, r *http.Request) {
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
		hcc.RespondError(w, "Error fetching boosted creature data.")
		return
	}
	response := map[string]interface{}{
		"boostedcreature": true,
		"raceid":          boostedCreature.RaceID,
	}
	hcc.RespondJSON(w, response)
}

func (hcc *HandlerClientConnect) loginHandler(answerExpected models.AnswerExpected, w http.ResponseWriter) (err error) {

	var accountCtl controller.AccountController

	account, err := accountCtl.GetAccountEmail(answerExpected.Email)
	if err != nil {
		utils.Warn("error get account", err.Error())
		err = errors.New("incorrect credentials")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errorCode":    3,
			"errorMessage": err.Error(),
		})
	}

	var apiCtl controller.ApiController
	passSha := apiCtl.ConvertSha1(answerExpected.Password)
	if passSha != account.Password {
		err = errors.New("incorrect credentials")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errorCode":    3,
			"errorMessage": err.Error(),
		})
		return
	}

	var playerCtl controller.PlayerController
	player := playerCtl.GetPlayersWithAccountID(account.ID)
	nowTime := time.Now().Unix()

	var session models.ClientSession
	session.IsPremium = uint32(account.PremiumEndsAt) > uint32(nowTime)
	session.LastLoginTime = uint32(time.Now().Unix())
	session.PremiumUntil = uint64(time.Now().Add(4 * time.Hour).Unix())
	session.OptionTracking = false
	session.SessionKey = fmt.Sprintf("%s\n%s\n%s\n%d", answerExpected.Email, answerExpected.Password, answerExpected.Token, time.Now().Add(30*time.Minute).Unix())
	session.Status = "active"
	session.IsReturner = true
	session.ShowRewardNews = false

	var world models.ClientWorld
	world.AntiCheatProtection = false
	world.ExternalAddRessUnProtected = config.Server.Server.IPServer
	world.ExternalAddress = config.Server.Server.IPServer
	world.ExternalAddressProtected = config.Server.Server.IPServer
	world.ExternalPort = config.Server.Server.GameProtocolPort
	world.ID = 0
	world.PvpType = config.Server.Server.WorldType
	world.Location = "CL"
	world.Name = config.Server.NameServer
	world.PreviewState = 1
	world.ExternalPortProtected = config.Server.Server.GameProtocolPort
	world.ExternalPortUnprotected = config.Server.Server.LoginProtocolPort
	world.CurrentTournamentPhase = 2

	var playdata models.PlayData
	playdata.World = append(playdata.World, world)
	playdata.Characters = apiCtl.PreparingCharacter(player)

	var responseData models.ResponseData
	responseData.PlayData = playdata
	responseData.Session = session

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(responseData); err != nil {
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

	}
}

// Helper functions
func (hcc *HandlerClientConnect) RespondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

}

func (hcc *HandlerClientConnect) RespondError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusConflict)
	hcc.RespondJSON(w, map[string]interface{}{
		"errorCode":    3,
		"errorMessage": msg,
	})
}
