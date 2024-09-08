package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

type PoolConnectionController struct{}

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

func (pc *PoolConnectionController) CharacterLoginAccountPoolConnection(answerExpected models.AnswerExpected) (models.ResponseData, error) {
	var account models.Account
	var response models.ResponseData
	var err error
	var apiCtl ApiController
	for _, pool := range config.Global.PoolServer {
		if pool.IpWebApi == "" {
			var accCtl AccountController
			account, err = accCtl.LoginAccesAccountClient(answerExpected)
			if err != nil {
				continue
			}
			response.PlayData.Characters = append(response.PlayData.Characters, apiCtl.PreparingCharacter(account.Players, uint(pool.World.ID))...)
		} else {
			jsonBody, err := json.Marshal(answerExpected)
			if err != nil {
				continue
			}
			req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlLoginClientConnection), bytes.NewBuffer(jsonBody))
			if err != nil {
				utils.Error("not create solcitiude", pool.IpWebApi)
				continue
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				utils.Error("error send session", err.Error())
				continue
			}
			err = json.NewDecoder(resp.Body).Decode(&account)
			if err != nil {
				utils.Error("error decode account", err.Error())
				continue
			}
			response.PlayData.Characters = append(response.PlayData.Characters, apiCtl.PreparingCharacter(account.Players, uint(pool.World.ID))...)
		}
		response.PlayData.World = append(response.PlayData.World, pool.World)
	}
	response.Session = pc.preparingSessionClien(account, answerExpected.Password, answerExpected.Token)
	return response, nil
}

func (pc *PoolConnectionController) preparingSessionClien(account models.Account, password string, token string) models.ClientSession {
	var session models.ClientSession

	nowTime := time.Now().Unix()
	session.IsPremium = uint32(account.PremiumEndsAt) > uint32(nowTime)
	session.LastLoginTime = uint32(time.Now().Unix())
	session.PremiumUntil = uint64(time.Now().Add(4 * time.Hour).Unix())
	session.OptionTracking = false
	session.SessionKey = fmt.Sprintf("%s\n%s\n%s\n%d", account.Email, password, token, time.Now().Add(30*time.Minute).Unix())
	session.Status = "active"
	session.IsReturner = true
	session.ShowRewardNews = false
	return session
}

func (pc *PoolConnectionController) CreateCharacter(nameCharacter, idWorld string, isMale int, accountID int) error {
	worldSub := strings.Split(idWorld, "-")

	idIndexWorld, err := strconv.ParseInt(worldSub[0], 10, 64)
	if err != nil {
		utils.Error("error generate id Index World register character", err.Error())
		return err
	}
	fmt.Println(idIndexWorld)

	if len(config.Global.PoolServer) < int(idIndexWorld) {
		utils.Error("out of the world index", fmt.Sprintf("%d", idIndexWorld))
		return errors.New("out of the world index")
	}

	var player models.Players
	player.AccountID = accountID
	player.Name = nameCharacter
	player.Sex = isMale

	var count int64
	db.DB.Where("name = ?", player.Name).Find(&models.PlayersNames{}).Count(&count)
	if count > 0 {
		return errors.New("error character clone name")
	}

	if config.Global.PoolServer[idIndexWorld].IpWebApi == "" {
		var playerCtl PlayerController
		player, err = playerCtl.CreatePlayer(player)
		if err != nil {
			return err
		}
	} else {
		playerJson, err := json.Marshal(player)
		if err != nil {
			utils.Error("error marshal playerjson", err.Error())
			return err
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", config.Global.PoolServer[idIndexWorld].IpWebApi, utils.ApiUrl, utils.ApiUrlRegisterCharacter), bytes.NewBuffer(playerJson))
		if err != nil {
			utils.Error("error generate solicitude create character", err.Error())
			return err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Global.PoolServer[idIndexWorld].Token))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			utils.Error("Error connection pool server", config.Global.PoolServer[idIndexWorld].IpWebApi)
			return err
		}
		if resp.StatusCode != http.StatusOK {
			utils.Error("error create character in server")
			return err
		}
	}

	db.DB.Create(&models.PlayersNames{
		Name:      player.Name,
		World:     worldSub[1],
		AccountID: player.AccountID,
	})

	return err
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

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			utils.Error("error send solicitude", pool.IpWebApi)
			continue
		}
		defer resp.Body.Close()

	}

}

func (pc *PoolConnectionController) SyncAccountPool() {
	var accountCtl AccountController
	account := accountCtl.GetAllAccount()
	accountCtl.GetAllAccount()
	body, err := json.Marshal(account)
	if err != nil {
		utils.Error(err.Error())
	}

	for _, pool := range config.Global.PoolServer {
		go pc.parallelSynAccount(pool, body)
	}

}

func (pc *PoolConnectionController) SyncPlayerNamePoolConnection() {

	go func() {

		for _, pool := range config.Global.PoolServer {
			if pool.IpWebApi == "" {
				var playerCtl PlayerController
				players := playerCtl.GetAllPlayer()
				for _, player := range players {
					db.DB.Create(&models.PlayersNames{
						Name:      player.Name,
						World:     pool.World.Name,
						AccountID: player.AccountID,
					})
				}
				continue
			}
			fmt.Println(fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlGetAllPlayers))

			req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlGetAllPlayers), nil)
			if err != nil {
				utils.Error(err.Error())
				continue
			}

			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				utils.Error("poolConnectionServer 300", "error send solicitude", pool.IpWebApi)
				continue
			}
			defer resp.Body.Close()

			var players []models.Players

			err = json.NewDecoder(resp.Body).Decode(&players)
			if err != nil {
				utils.Error("poolConnectionServer 309", "error send solicitude", pool.IpWebApi)
				continue
			}

			for _, player := range players {
				db.DB.Create(&models.PlayersNames{
					Name:      player.Name,
					World:     pool.World.Name,
					AccountID: player.AccountID,
				})
			}

		}
	}()
}

func (pc *PoolConnectionController) parallelSynAccount(pool models.PoolServer, body []byte) {
	if pool.IpWebApi == "" || pool.Token == "" {
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlMySyncAccount), bytes.NewBuffer(body))
	if err != nil {
		utils.Error("error sync send api: ", pool.IpWebApi)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.Error("error send solicitude", pool.IpWebApi)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusExpectationFailed {
		var exp models.Exception
		err = json.NewDecoder(resp.Body).Decode(&exp)
		if err != nil {
			utils.Error("error decode captured exception")
			return
		}
		utils.Error("error poolConnectionServer 354", "error connection", exp.Msg)
	} else {
		var msg []string
		err = json.NewDecoder(resp.Body).Decode(&msg)
		if err != nil {
			utils.Error("error decode msg alerts")
			return
		}
		for _, msgCaptured := range msg {
			//Info order in console
			utils.InfoBlue(msgCaptured)
		}

	}
}
