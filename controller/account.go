package controller

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type AccountController struct{}

func (ac *AccountController) GetAccountWithPlayer(accountID int) (account models.Account, err error) {

	if err = db.DB.Preload("Players").Where("id = ?", accountID).First(&account).Error; err != nil {
		return
	}
	return
}

func (ac *AccountController) GetAccountEmail(email string) (account models.Account, err error) {

	if err = db.DB.Where("email = ?", email).First(&account).Error; err != nil {
		return
	}
	return
}

func (ac *AccountController) GetAccountName(name string) (account models.Account, err error) {
	if err = db.DB.Where("name = ?", name).First(&account).Error; err != nil {
		return
	}
	return
}

func (ac *AccountController) GetAccountID(id int) (account models.Account, err error) {
	if err = db.DB.Where("id = ?", id).First(&account).Error; err != nil {
		return
	}
	return
}

func (ac *AccountController) AuthenticationAccount(userOrEmail, password string) (account models.Account, err error) {

	account, err = ac.GetAccountEmail(userOrEmail)
	if err != nil {
		account, err = ac.GetAccountName(userOrEmail)
		if err != nil {
			err = errors.New(utils.ErrorEmailOrUser)
			return
		}
	}
	var apiCtl ApiController
	passEncryp := apiCtl.ConvertSha1(password)
	if account.Password != passEncryp {
		err = errors.New(utils.ErrorPasswordEquals)
		return
	}
	return
}

func (ac *AccountController) CreateAccount(account models.Account) (models.Account, error) {
	var api ApiController
	account.Email = strings.ToLower(account.Email)
	account.Name = strings.ToLower(account.Name)
	account.Password = api.ConvertSha1(account.Password)
	if err := db.DB.Create(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}

func (ac *AccountController) LoginAccountClient(answerExpected models.AnswerExpected) (playData models.PlayData, session models.ClientSession, err error) {
	var account models.Account
	if err = db.DB.Where("email = ?", answerExpected.Email).First(&account).Error; err != nil {
		return
	}
	var apiCtl ApiController
	passSha := apiCtl.ConvertSha1(answerExpected.Password)
	if passSha != account.Password {
		err = errors.New("incorrect credentials")
		return
	}

	var playerCtl PlayerController
	players := playerCtl.GetPlayersWithAccountID(account.ID)

	nowTime := time.Now().Unix()

	session.IsPremium = uint32(account.PremiumEndsAt) > uint32(nowTime)
	session.LastLoginTime = uint32(time.Now().Unix())
	session.PremiumUntil = uint64(time.Now().Add(4 * time.Hour).Unix())
	session.OptionTracking = false
	session.SessionKey = fmt.Sprintf("%s\n%s\n%s\n%d", answerExpected.Email, answerExpected.Password, answerExpected.Token, time.Now().Add(30*time.Minute).Unix())
	session.Status = "active"
	session.IsReturner = true
	session.ShowRewardNews = false

	var worlds []models.ClientWorld
	for _, value := range config.Global.PoolSerer {
		worlds = append(worlds, value.World)
	}
	playData.World = worlds
	playData.Characters = apiCtl.PreparingCharacter(players)

	return
}
