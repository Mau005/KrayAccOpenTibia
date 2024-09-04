package controller

import (
	"github.com/Mau005/KrayAccOpenTibia/models"
)

type PoolConnectionController struct{}

func (pc *PoolConnectionController) GeneratePool(answerExpected models.AnswerExpected) (response models.ResponseData, err error) {
	var accountCtl AccountController

	playdata, session, err := accountCtl.LoginAccountClient(answerExpected)
	if err != nil {
		return
	}

	response.PlayData = playdata
	response.Session = session

	return

}

func (pc *PoolConnectionController) GetAccountPool(answerExpected models.AnswerExpected) {

}
