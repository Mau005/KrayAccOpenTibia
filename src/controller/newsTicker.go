package controller

import (
	"errors"

	"github.com/Mau005/KrayAccOpenTibia/src/db"
	"github.com/Mau005/KrayAccOpenTibia/src/models"
	"github.com/Mau005/KrayAccOpenTibia/src/utils"
)

type NewsTickerController struct{}

func (ntc *NewsTickerController) GetTickerLimited(count int) (ticker []models.NewsTicket) {
	db.DB.Preload("Player").Limit(count).Order("created_at DESC").Find(&ticker)
	return
}

func (ntc *NewsTickerController) GetIconID(id uint8) (result string) {
	switch id {
	case utils.IconNewsTicketCommunity:
		result = utils.PathIconNewsTicketCommunity

	case utils.IconNewsTicketSupport:
		result = utils.PathIconNewsTicketSupport

	case utils.IconNewsTicketTechnical:
		result = utils.PathIconNewsTicketTechnical

	case utils.IconNewsTicketDevelopment:
		result = utils.PathIconNewsTicketDevelopment
	default:
		result = utils.PathIconNewsTicketSupport
	}
	return
}

func (ntc *NewsTickerController) RulesTicker(ticket models.NewsTicket) (models.NewsTicket, error) {
	var err error

	if ticket.IconID == 0 {
		err = errors.New("icon not found")
		return ticket, err
	}
	return ticket, nil
}

func (ntc *NewsTickerController) CreateTicker(ticket models.NewsTicket, accountID int) (models.NewsTicket, error) {
	ticket, err := ntc.RulesTicker(ticket)
	if err != nil {
		return ticket, err
	}

	var playerCtl PlayerController
	if !playerCtl.GetPropertiesPlayer(accountID, int(ticket.PlayersID)) {
		err = errors.New("this character does not belong to you")
		return ticket, err
	}

	if err := db.DB.Create(&ticket).Error; err != nil {
		return ticket, err
	}
	return ticket, nil
}

func (ntc *NewsTickerController) GetTicker(id uint) (ticker models.NewsTicket, err error) {

	if err = db.DB.Preload("Player").Where("id = ?", id).Find(&ticker).Error; err != nil {
		return
	}

	return
}

func (ntc *NewsTickerController) PutTicker(ticker models.NewsTicket) (models.NewsTicket, error) {

	tickerOld, err := ntc.GetTicker(ticker.ID)
	if err != nil {
		return ticker, err
	}
	// TODO: add fix configure edit ticker
	tickerOld.IconID = ticker.IconID

	tickerOld.Ticket = ticker.Ticket

	if err := db.DB.Save(&tickerOld).Error; err != nil {
		return ticker, err
	}
	return ticker, nil
}

func (ntc *NewsTickerController) GetNewsShort() (news []models.NewsShort) {
	db.DB.Limit(5).Order("created_at DESC").Find(&news)
	return
}

func (ntc *NewsTickerController) CreateNewsShort(newsenter models.NewsShort) (models.NewsShort, error) {
	if err := db.DB.Create(&newsenter).Error; err != nil {
		return newsenter, err
	}
	return newsenter, nil
}
