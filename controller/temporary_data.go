package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

var TempData *TemporaryData

type TemporaryData struct {
	ServStatusTotal models.ServerStatus
	ServStatus      []models.ServerStatus
}

func LoadTemporaryData() error {
	TempData = &TemporaryData{ServStatusTotal: models.ServerStatus{}}
	var api ApiController
	go func() {
		for {
			// TODO: server local check status
			playerOnline := 0
			MonsterTotal := 0
			oldestUptime := 0
			for _, pool := range config.Global.PoolServer {
				serv, err := api.CheckOnlineServer(pool.World.ExternalAddress, fmt.Sprintf("%d", pool.World.ExternalPort))
				if err != nil {
					utils.Error(fmt.Sprintf("error check online server xml function, %s", pool.IpWebApi), err.Error())
					continue
				}

				playerOnline += serv.Players.Online
				MonsterTotal += serv.Monsters.Total

				TempData.ServStatusTotal.NPCs.Total += serv.NPCs.Total

				timeNow, _ := strconv.ParseInt(serv.ServerInfo.Uptime, 10, 64)

				if oldestUptime <= int(timeNow) {
					oldestUptime = int(timeNow)
				}

				TempData.ServStatus = append(TempData.ServStatus, serv) // check server
			}
			TempData.ServStatusTotal.ServerInfo.Uptime = string(oldestUptime)
			TempData.ServStatusTotal.Players.Online = playerOnline
			TempData.ServStatusTotal.Monsters.Total = MonsterTotal

			time.Sleep(utils.TimeCheckInfoServer * time.Minute)
		}
	}()

	go func() {
		// Normalice database init proyect
		var PoolConnectionController PoolConnectionController
		PoolConnectionController.GetWorldPool()
		PoolConnectionController.SyncAccountPool()
		PoolConnectionController.SyncPlayerNamePoolConnection()
	}()
	return nil
}
