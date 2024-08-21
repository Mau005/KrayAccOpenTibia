package controller

import (
	"github.com/Mau005/KrayAcc/db"
	"github.com/Mau005/KrayAcc/models"
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
