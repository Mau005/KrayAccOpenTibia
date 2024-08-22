package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type AccountHandler struct{}

func (ah *AccountHandler) Authentication(w http.ResponseWriter, r *http.Request) {

	var accountCtl controller.AccountController
	account, err := accountCtl.AuthenticationAccount(r.FormValue("UserOrEmail"), r.FormValue("password"))
	if err != nil {
		log.Println("error authentication account", err)
		http.Error(w, "Error credentials account", http.StatusInternalServerError)
		return
	}

	var apiCtl controller.ApiController
	token, err := apiCtl.GenerateJWT(account)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     utils.NameCookieToken,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	json.NewEncoder(w).Encode(struct {
		Token       string    `json:"token"`
		TimeCurrent time.Time `json:"timecurrent"`
	}{
		Token:       token,
		TimeCurrent: time.Now(),
	})

}
