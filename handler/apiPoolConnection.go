package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type ApiPoolConnectionHandler struct{}

func (apc *ApiPoolConnectionHandler) GetPoolConnection(w http.ResponseWriter, r *http.Request) {
	var errorException controller.ExceptionController
	if len(config.Global.PoolServer) == 0 {
		errorException.Exeption(fmt.Sprintf("error server not register %d", len(config.Global.PoolServer)), http.StatusInternalServerError, w)
		return
	}

	err := json.NewEncoder(w).Encode(&config.Global.PoolServer)
	if err != nil {
		errorException.Exeption("error decode web", http.StatusConflict, w)
		return
	}
}

func (apc *ApiPoolConnectionHandler) RegisterNewAccount(w http.ResponseWriter, r *http.Request) {

	var accountAPI struct {
		Account           models.Account
		PasswordEncrypted string
	}
	err := json.NewDecoder(r.Body).Decode(&accountAPI)
	if err != nil {
		return
	}
	account := accountAPI.Account
	account.Password = accountAPI.PasswordEncrypted

	var accController controller.AccountController

	_, err = accController.CreateAccountAPI(account)
	if err != nil {
		return
	}

}

func (apc *ApiPoolConnectionHandler) RegisterNewCharacter(w http.ResponseWriter, r *http.Request) {
	var ExceptionController controller.ExceptionController
	var players models.Players

	err := json.NewDecoder(r.Body).Decode(&players)
	if err != nil {
		ExceptionController.Exeption(err.Error(), http.StatusConflict, w)
		return
	}
	var playerCtl controller.PlayerController
	players.Level = config.Global.ServerWeb.DefaultPlayer.Level
	players.Experience = config.Global.ServerWeb.DefaultPlayer.Experience
	players.Health = config.Global.ServerWeb.DefaultPlayer.HealthMax
	players.HealthMax = config.Global.ServerWeb.DefaultPlayer.HealthMax
	players.Mana = config.Global.ServerWeb.DefaultPlayer.ManaMax
	players.ManaMax = config.Global.ServerWeb.DefaultPlayer.ManaMax
	players.Cap = config.Global.ServerWeb.DefaultPlayer.Cap
	players.TownID = config.Global.ServerWeb.DefaultPlayer.TownID
	players.Vocation = config.Global.ServerWeb.DefaultPlayer.Vocation
	_, err = playerCtl.CreatePlayer(players)
	if err != nil {
		ExceptionController.Exeption(err.Error(), http.StatusInternalServerError, w)
		return
	}
}

func (apc *ApiPoolConnectionHandler) LoginAccountPoolConnection(w http.ResponseWriter, r *http.Request) {
	var errorException controller.ExceptionController
	var answer models.AnswerExpected
	err := json.NewDecoder(r.Body).Decode(&answer)

	if err != nil {
		log.Println("error decode body", err)
	}

	var accountCtl controller.AccountController
	account, err := accountCtl.LoginAccesAccountClient(answer)
	if err != nil {
		errorException.Exeption(err.Error(), http.StatusConflict, w)
		return
	}
	json.NewEncoder(w).Encode(&account)
}

func (apc *ApiPoolConnectionHandler) SyncAccountPoolConnection(w http.ResponseWriter, r *http.Request) {
	var PoolConnectionController controller.PoolConnectionController
	PoolConnectionController.SyncAccountPool()
}

func (apc *ApiPoolConnectionHandler) MySyncAccountData(w http.ResponseWriter, r *http.Request) {
	var accounts []models.Account
	json.NewDecoder(r.Body).Decode(&accounts)
	var accountCTL controller.AccountController
	var accountNormalice []string
	for _, value := range accounts {
		_, err := accountCTL.CreateAccountPoolConnection(value)
		if err == nil {
			accountNormalice = append(accountNormalice, fmt.Sprintf("AccountID %d normalice in database accountName: %s", value.ID, value.Name))
		}
	}

	if len(accountNormalice) > 0 {
		json.NewEncoder(w).Encode(&accountNormalice)
	} else {
		var errorCtl controller.ExceptionController
		errorCtl.Exeption("not have changes", http.StatusExpectationFailed, w)
	}
}

func (apc *ApiPoolConnectionHandler) SynPlayerName(w http.ResponseWriter, r *http.Request) {
	var PoolConnectionController controller.PoolConnectionController
	PoolConnectionController.SyncPlayerNamePoolConnection()
}

func (apc *ApiPoolConnectionHandler) GetAllPlayer(w http.ResponseWriter, r *http.Request) {
	var playerCtl controller.PlayerController
	players := playerCtl.GetAllPlayer()
	json.NewEncoder(w).Encode(&players)
}

func (apc *ApiPoolConnectionHandler) WhoIsOnline(w http.ResponseWriter, r *http.Request) {
	var playerCtl controller.PlayerController
	players := playerCtl.GetPlayerOnline()
	json.NewEncoder(w).Encode(&players)
}

func (apc *ApiPoolConnectionHandler) GetPlayerAccount(w http.ResponseWriter, r *http.Request) {
	var accountCtl controller.AccountController
	var account models.Account
	var errorCtl controller.ExceptionController
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		errorCtl.Exeption(err.Error(), http.StatusConflict, w)
		return
	}

	account, err := accountCtl.GetAccountWithPlayer(account.ID)
	if err != nil {
		errorCtl.Exeption(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if err := json.NewEncoder(w).Encode(&account); err != nil {
		errorCtl.Exeption(err.Error(), http.StatusConflict, w)
		return
	}
}

func (apc *ApiPoolConnectionHandler) GetNews(w http.ResponseWriter, r *http.Request) {
	var newsCtl controller.NewsTickerController

	news := newsCtl.GetTickerLimited(utils.LimitRecordFive)

	json.NewEncoder(w).Encode(news)
}

func (apc *ApiPoolConnectionHandler) GetHighScore(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ID int
	}
	json.NewDecoder(r.Body).Decode(&request)

	var playerCtl controller.PlayerController
	players := playerCtl.GetHighScore(request.ID)
	json.NewEncoder(w).Encode(&players)

}
