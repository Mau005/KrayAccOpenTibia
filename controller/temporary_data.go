package controller

import (
	"fmt"
	"log"
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
			playerNPC := 0
			MonsterTotal := 0
			oldestUptime := 0
			for _, pool := range config.Global.PoolServer {
				serv, err := api.CheckOnlineServer(pool.World.ExternalAddress, fmt.Sprintf("%d", pool.World.ExternalPort))
				if err != nil {
					log.Println(err)
					continue
				}

				playerOnline += serv.Players.Online
				MonsterTotal += serv.Monsters.Total

				playerNPC += serv.NPCs.Total

				timeNow, _ := strconv.ParseInt(serv.ServerInfo.Uptime, 10, 64)

				if oldestUptime <= int(timeNow) {
					oldestUptime = int(timeNow)
				}
			}

			if oldestUptime == 0 {
				TempData.ServStatusTotal.ServerInfo.Uptime = ""
			} else {

				TempData.ServStatusTotal.ServerInfo.Uptime = strconv.Itoa(oldestUptime)
			}

			TempData.ServStatusTotal.Players.Online = playerOnline

			TempData.ServStatusTotal.Monsters.Total = MonsterTotal
			TempData.ServStatusTotal.NPCs.Total = playerNPC

			time.Sleep(utils.TimeCheckInfoServer * time.Second)
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
