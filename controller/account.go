package controller

import (
	"errors"

	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type AccountController struct{}

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

func (ac *AccountController) CreateAccount(userOrEmail, password, passwordTwo string) (account models.Account, err error) {
	if password != passwordTwo {
		err = errors.New(utils.ErrorPasswordEquals)
		return
	}

	account, err = ac.GetAccountEmail(userOrEmail)
	if err != nil {
		account, err = ac.GetAccountName(userOrEmail)
		if err != nil {
			err = errors.New(utils.ErrorEmailOrUser)
			return
		}
	}
	return
}
