package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
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

func (pc *PoolConnectionController) GetWorldPool() {
	for index, pool := range config.Global.PoolServer {
		if len(pool.Token) == 0 || len(pool.IpWebApi) == 0 && len(pool.World.ExternalAddress) >= 0 {
			config.Global.PoolServer[index].World.ID = index
			config.Global.PoolServer[index] = pool
			continue
		} else if len(pool.Token) == 0 || len(pool.IpWebApi) == 0 {
			utils.Error("error pool not have data ", "index array:", fmt.Sprintf("%d", index))
			continue
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlGetPoolConnect), nil)
		if err != nil {
			utils.Error("error create solicitude http ip:", pool.IpWebApi, err.Error())
			continue
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))
		req.Header.Set("Content-Type", "applicantion/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			utils.Error("error send solicitude", pool.IpWebApi)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			utils.Error("error read body", err.Error())
			continue
		}
		var poolResponse []models.PoolServer
		err = json.Unmarshal(body, &poolResponse)
		if err != nil {
			utils.Error("decode unmarshal response router", err.Error())
			continue
		}

		if len(poolResponse) == 1 {
			pool.World = poolResponse[0].World
			pool.RateServer = poolResponse[0].RateServer
			config.Global.PoolServer[index].World.ID = index
			config.Global.PoolServer[index] = pool
		} else {
			utils.Warn("error connection not support other connections apis", pool.IpWebApi)
		}
		utils.Info("connected to ", pool.IpWebApi)
	}
	// pc.orderPoolConnectionID()
}

// // work target var temporary pool connection
// func (pc *PoolConnectionController) orderPoolConnectionID() {
// 	//Order pool connection ID series
// 	for index, _ := range config.Global.PoolServer {
// 		config.Global.PoolServer[index].World.ID = index
// 	}
// }

func (pc *PoolConnectionController) GetAccountPool(answerExpected models.AnswerExpected) {

}

func (pc *PoolConnectionController) CreateAccountPool(account models.Account) {
	utils.InfoBlue("entro a create Account Pool")
	var Request struct {
		Account           models.Account
		PasswordEncrypted string
	}
	// account send not have password, send other methods, security web secuence do password
	Request.Account = account
	Request.PasswordEncrypted = account.Password

	for _, pool := range config.Global.PoolServer {

		jsonData, err := json.Marshal(Request)
		if err != nil {
			utils.Error("error json marshal pool create account")
			continue
		}

		if len(pool.Token) == 0 || len(pool.IpWebApi) == 0 {
			fmt.Println(pool)
			utils.Error("error len token or ipwebapi")
			continue
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlCreateAccount), bytes.NewBuffer(jsonData))
		if err != nil {
			utils.Error("error create account solicidute http: ", pool.IpWebApi)
			continue
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))
		req.Header.Set("Content-Type", "applicantion/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			utils.Error("error send solicitude", pool.IpWebApi)
			continue
		}
		defer resp.Body.Close()

	}

}
