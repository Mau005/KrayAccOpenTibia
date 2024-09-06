package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
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

	_, err = playerCtl.CreatePlayer(players)
	if err != nil {
		ExceptionController.Exeption(err.Error(), http.StatusInternalServerError, w)
		return
	}

}
