package controller

import (
	"errors"
	"strings"

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

func (ac *AccountController) GetAllAccount() (account []models.Account) {

	db.DB.Find(&account)

	return
}

// create account valida web main, to connection to user server, sync preview
func (ac *AccountController) CreateAccountAPI(account models.Account) (models.Account, error) {

	if err := db.DB.Create(&account).Error; err != nil {
		return account, err
	}
	return account, nil
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

func (ac *AccountController) CreateAccountPoolConnection(account models.Account) (models.Account, error) {
	if err := db.DB.Create(&account).Error; err != nil {
		return account, err
	}
	return account, nil
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

func (ac *AccountController) LoginAccesAccountClient(answerExpected models.AnswerExpected) (account models.Account, err error) {
	if err = db.DB.Preload("Players").Where("email = ?", answerExpected.Email).First(&account).Error; err != nil {
		return
	}
	var apiCtl ApiController
	passSha := apiCtl.ConvertSha1(answerExpected.Password)
	if passSha != account.Password {
		err = errors.New("incorrect credentials")
		return
	}

	return
}
