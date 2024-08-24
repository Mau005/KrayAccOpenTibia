package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type AccountHandler struct{}

func (ah *AccountHandler) Authentication(w http.ResponseWriter, r *http.Request) {
	var exceptCtl controller.ExceptionController
	var request struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Warn("error decode", err.Error())
		exceptCtl.Exeption(err.Error(), http.StatusConflict, w)
		return
	}

	var accountCtl controller.AccountController
	account, err := accountCtl.AuthenticationAccount(request.User, request.Password)
	if err != nil {
		utils.Warn("error authentication account", err.Error())
		exceptCtl.Exeption(err.Error(), http.StatusConflict, w)
		return
	}

	var apiCtl controller.ApiController
	token, err := apiCtl.GenerateJWT(account)
	if err != nil {
		utils.Warn("error generate token", err.Error())
		exceptCtl.Exeption(err.Error(), http.StatusConflict, w)
		return
	}

	secure := config.VarEnviroment.Certificate.ProtolTLS

	if !config.VarEnviroment.ServerWeb.ApiMode {
		http.SetCookie(w, &http.Cookie{
			Name:     utils.NameCookieToken,
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   secure,
			SameSite: http.SameSiteStrictMode,
		})
	}

	json.NewEncoder(w).Encode(struct {
		Token       string    `json:"token"`
		TimeCurrent time.Time `json:"timecurrent"`
	}{
		Token:       token,
		TimeCurrent: time.Now(),
	})

}
