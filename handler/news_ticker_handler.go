package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/gorilla/context"
)

type NewsTicketHandler struct{}

func (ntc *NewsTicketHandler) GetTicketLimited(w http.ResponseWriter, r *http.Request) {

}

func (ntc *NewsTicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var exceptCtl controller.ExceptionController
	claim := context.Get(r, utils.CtxClaim).(models.Claim)

	if claim.TypeAccess <= utils.UserCommunityManager {
		utils.WarnLog(fmt.Sprintf("account name %s try create ticker", claim.AccountName))
		return
	}

	var Ticket models.NewsTicket

	err := json.NewDecoder(r.Body).Decode(&Ticket)
	if err != nil {
		utils.Warn("error decode json in create news ticker")
		exceptCtl.Exeption(err.Error(), http.StatusConflict, w)
		return
	}

	var newsTicketCtl controller.NewsTickerController
	Ticket, err = newsTicketCtl.CreateTicker(Ticket, claim.AccountID)
	if err != nil {
		utils.Warn(err.Error())
		exceptCtl.Exeption(err.Error(), http.StatusConflict, w)
		return
	}

	err = json.NewEncoder(w).Encode(&Ticket)
	if err != nil {
		utils.Warn("error encoder json in create news ticker")
		exceptCtl.Exeption(err.Error(), http.StatusConflict, w)
		return
	}
}

func (ntc *NewsTicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {

	var newsTicketCtl controller.NewsTickerController

	type Response struct {
		Icon        string
		Ticket      string
		ByCharacter string
	}

	var response []Response

	tickets, _ := newsTicketCtl.GetTickerLimited(5)

	for _, value := range tickets {
		response = append(response, Response{
			Icon:        newsTicketCtl.GetIconID(value.IconID),
			Ticket:      value.Ticket,
			ByCharacter: value.Player.Name,
		})
	}

	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		utils.Warn("error encoder json in create news ticker")
		return
	}

}
