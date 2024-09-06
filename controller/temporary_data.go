package controller

import (
	"fmt"
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
	var poolCtl PoolConnectionController

	if len(config.Global.PoolServer) > 1 {
		poolCtl.GetWorldPool()
	}

	counTry := 0
	var api ApiController
	go func() {
		for {
			// TODO: server local check status
			counTry += 1
			// oldestUptime := 0
			for _, value := range config.Global.PoolServer {
				serv, err := api.CheckOnlineServer(value.World.ExternalAddress, fmt.Sprintf("%d", value.World.ExternalPort))
				if err != nil {
					utils.Error(fmt.Sprintf("error check online server xml function, %s", value.IpWebApi), err.Error())
					continue
				}

				TempData.ServStatusTotal.Players.Online += serv.Players.Online
				TempData.ServStatusTotal.Monsters.Total += serv.Monsters.Total
				TempData.ServStatusTotal.NPCs.Total += serv.NPCs.Total

				//check rev old time server
				// if serv.ServerInfo.Uptime >= oldestUptime {
				// 	oldestUptime = serv.ServerInfo.Uptime
				// }

				TempData.ServStatus = append(TempData.ServStatus, serv) // check server
			}
			// TempData.ServStatusTotal.ServerInfo.Uptime = oldestUptime

			if counTry >= 3 {
				return
			}
			time.Sleep(utils.TimeCheckInfoServer * time.Minute)
		}
	}()

	return nil
}
