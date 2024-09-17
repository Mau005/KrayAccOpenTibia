package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/components"
	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/gorilla/context"
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
		exceptCtl.Exeption(err.Error(), http.StatusInternalServerError, w)
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

	secure := false

	if !config.Global.ServerWeb.ApiMode {
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

func (ah *AccountHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	claim := context.Get(r, utils.CtxClaim).(models.Claim)
	var exceptCtl controller.ExceptionController

	var accountCtl controller.AccountController
	account, err := accountCtl.GetAccountWithPlayer(claim.AccountID)
	if err != nil {
		utils.WarnLog("error try create character is not found",
			fmt.Sprintf("claim try: %v", claim), err.Error())
		exceptCtl.Exeption(err.Error(), http.StatusFailedDependency, w)
		return
	}

	if len(account.Players) >= int(config.Global.ServerWeb.LimitCreateCharacter) {
		utils.Warn("account completed limite character")
		exceptCtl.Exeption("complete limit character", http.StatusPreconditionRequired, w)
		return
	}

	var request struct {
		NameCharacter string
		IsMale        int
		World         string `json:"world"`
	}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Warn("error decore", err.Error())
		exceptCtl.Exeption(err.Error(), http.StatusInternalServerError, w)
		return
	}

	var PoolConnectionController controller.PoolConnectionController
	err = PoolConnectionController.CreateCharacter(request.NameCharacter, request.World, request.IsMale, claim.AccountID)
	if err != nil {
		utils.Warn("error create player ", err.Error())
		exceptCtl.Exeption(err.Error(), http.StatusConflict, w)
		return
	}

	exceptCtl.MessageAproved("ok", w)

}

func (ah *AccountHandler) Desconnected(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  utils.NameCookieToken,
		Value: "nil",
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (ah *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var exceptCtl controller.ExceptionController
	var request struct {
		IsTerms     bool
		UserName    string
		Password    string
		PasswordTwo string
		Email       string
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.Warn("error decode request", err.Error())
		exceptCtl.Exeption(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if !(request.Password == request.PasswordTwo) {
		utils.Warn("password not equals")
		exceptCtl.Exeption("password equals", http.StatusExpectationFailed, w)
		return
	}
	if !request.IsTerms {
		utils.Warn("not accept term")
		exceptCtl.Exeption("not accept term", http.StatusConflict, w)
		return
	}

	var accountCtl controller.AccountController
	account, err := accountCtl.CreateAccount(models.Account{
		Name:     request.UserName,
		Password: request.Password,
		Email:    request.Email,
	})
	if err != nil {
		utils.Warn("error create account", err.Error())
		exceptCtl.Exeption(err.Error(), http.StatusNotAcceptable, w)
		return
	}

	//pool connections
	var poolConnectionCTL controller.PoolConnectionController
	go poolConnectionCTL.CreateAccountPool(account)

	err = json.NewEncoder(w).Encode(&account)
	if err != nil {
		utils.Warn("error encoder", err.Error())
		exceptCtl.Exeption(err.Error(), http.StatusInternalServerError, w)
	}

}

func (ah *AccountHandler) MyAccount(w http.ResponseWriter, r *http.Request) {
	navWeb, _ := context.Get(r, utils.CtxNavWeb).(models.NavWeb)

	if !navWeb.Authentication {
		http.Redirect(w, r, "/", http.StatusOK)
		return
	}

	templ, err := template.New("my_account.html").ParseFiles("www/my_account.html")
	if err != nil {
		log.Println("error create template", err)
		return
	}

	var Layouthandler Layouthandler
	layout := Layouthandler.Generatelayout(navWeb, models.SolicitudeLayout{News: true, Login: true, ServerStatus: true, TopPlayers: true, Rates: true})
	var poolConnectionCTL controller.PoolConnectionController
	accountCheck := poolConnectionCTL.GetACcountPlayerPoolConenction(navWeb.AccountID)
	layout.MyAccount = components.CreateMyAccount(accountCheck)

	if err := templ.Execute(w, layout); err != nil {
		log.Println("error execute template", err)
		return
	}
}
