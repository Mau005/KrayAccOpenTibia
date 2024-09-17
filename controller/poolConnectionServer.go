package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
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
			log.Println("error pool not have data ", "index array:", fmt.Sprintf("%d", index))
			continue
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlGetPoolConnect), nil)
		if err != nil {
			log.Println("error create solicitude http ip:", pool.IpWebApi, err.Error())
			continue
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))
		req.Header.Set("Content-Type", "applicantion/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("error send solicitude", pool.IpWebApi)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("error read body", err.Error())
			continue
		}
		var poolResponse []models.PoolServer
		err = json.Unmarshal(body, &poolResponse)
		if err != nil {
			log.Println("decode unmarshal response router", err.Error())
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
				log.Println("not create solcitiude", pool.IpWebApi)
				continue
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Println("error send session", err.Error())
				continue
			}
			err = json.NewDecoder(resp.Body).Decode(&account)
			if err != nil {
				log.Println("error decode account", err.Error())
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
		log.Println("error generate id Index World register character", err)
		return err
	}

	if len(config.Global.PoolServer) < int(idIndexWorld) {
		log.Println("out of the world index", fmt.Sprintf("%d", idIndexWorld))
		return errors.New("out of the world index")
	}

	var player models.Players
	player.AccountID = accountID
	player.Name = nameCharacter
	player.Sex = isMale

	player.Level = config.Global.ServerWeb.DefaultPlayer.Level
	player.Experience = config.Global.ServerWeb.DefaultPlayer.Experience
	player.Health = config.Global.ServerWeb.DefaultPlayer.HealthMax
	player.HealthMax = config.Global.ServerWeb.DefaultPlayer.HealthMax
	player.Mana = config.Global.ServerWeb.DefaultPlayer.ManaMax
	player.ManaMax = config.Global.ServerWeb.DefaultPlayer.ManaMax
	player.Cap = config.Global.ServerWeb.DefaultPlayer.Cap
	player.TownID = config.Global.ServerWeb.DefaultPlayer.TownID
	player.Vocation = config.Global.ServerWeb.DefaultPlayer.Vocation

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
			return err
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", config.Global.PoolServer[idIndexWorld].IpWebApi, utils.ApiUrl, utils.ApiUrlRegisterCharacter), bytes.NewBuffer(playerJson))
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Global.PoolServer[idIndexWorld].Token))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
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
			log.Println("error json marshal pool create account")
			continue
		}

		if len(pool.Token) == 0 || len(pool.IpWebApi) == 0 {
			log.Println("error len token or ipwebapi")
			continue
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlCreateAccount), bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("error create account solicidute http: ", pool.IpWebApi)
			continue
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("error send solicitude", pool.IpWebApi)
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
		log.Println(err)
	}

	for _, pool := range config.Global.PoolServer {
		go func() {
			err := pc.parallelSynAccount(pool, body)
			if err != nil {
				log.Println(err)
			}
		}()
	}

}

func (pc *PoolConnectionController) WhoIsOnlinePoolConnection() map[string][]models.Players {
	//var poolPlayerOnline map[string][]models.Players
	poolPlayerOnline := make(map[string][]models.Players)

	for _, pool := range config.Global.PoolServer {
		if pool.IpWebApi == "" {
			var playerCtl PlayerController
			whoisOnline := playerCtl.GetPlayerOnline()
			poolPlayerOnline[pool.World.Name] = whoisOnline
			continue
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlWhoIsOnline), nil)
		if err != nil {
			log.Println("error send woisonline pool connection", err)
			continue
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))

		reso := &http.Client{}
		r, err := reso.Do(req)
		if err != nil {
			log.Println("error send solicitude whoIsOnline", err)
			continue
		}
		var players []models.Players

		if r.StatusCode != http.StatusOK {
			continue
		}

		json.NewDecoder(r.Body).Decode(&players)

		poolPlayerOnline[pool.World.Name] = players
	}
	return poolPlayerOnline
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

			req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlGetAllPlayers), nil)
			if err != nil {
				log.Println(err)
				continue
			}

			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Println("poolConnectionServer 300", "error send solicitude", pool.IpWebApi)
				continue
			}
			defer resp.Body.Close()

			var players []models.Players

			err = json.NewDecoder(resp.Body).Decode(&players)
			if err != nil {
				log.Println("poolConnectionServer 309", "error send solicitude", pool.IpWebApi)
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

func (pc *PoolConnectionController) parallelSynAccount(pool models.PoolServer, body []byte) (err error) {
	if pool.IpWebApi == "" || pool.Token == "" {
		return nil
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlMySyncAccount), bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusExpectationFailed {
		var exp models.Exception
		err = json.NewDecoder(resp.Body).Decode(&exp)
		if err != nil {
			return
		}
		return errors.New(fmt.Sprintf("error poolConnectionServer 354 error connection %s", exp.Msg))
	} else {
		var msg []string
		err = json.NewDecoder(resp.Body).Decode(&msg)
		if err != nil {
			return
		}
		for _, msgCaptured := range msg {
			//Info order in console
			utils.InfoBlue(msgCaptured)
		}

	}
	return nil
}

func (pc *PoolConnectionController) GetACcountPlayerPoolConenction(accountID int) (account models.Account) {

	for _, pool := range config.Global.PoolServer {

		if pool.IpWebApi == "" {
			var accountCtl AccountController
			acc, err := accountCtl.GetAccountWithPlayer(accountID)
			if err != nil {
				log.Println(err)
				continue
			}
			account.Players = append(account.Players, acc.Players...)
			account.Name = acc.Name
			account.Email = acc.Email
			continue
		}

		accByte, err := json.Marshal(models.Account{ID: accountID})
		if err != nil {
			log.Println(err)
			continue
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s", pool.IpWebApi, utils.ApiUrl, utils.ApiUrlGetPlayerAccount), bytes.NewBuffer(accByte))
		if err != nil {
			log.Println(err)
			continue
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pool.Token))

		client := &http.Client{}
		body, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}

		var accountFinaly models.Account

		if err := json.NewDecoder(body.Body).Decode(&accountFinaly); err != nil {
			log.Println(err)
			continue
		}

		account.Players = append(account.Players, accountFinaly.Players...)
		account.Email = accountFinaly.Email
		account.Name = accountFinaly.Name
	}
	return account
}
