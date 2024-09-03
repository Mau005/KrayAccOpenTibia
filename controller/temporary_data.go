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
	ServerStatus models.ServerStatus
}

func LoadTemporaryData() error {
	TempData = &TemporaryData{}

	if config.VarEnviroment.ServerWeb.TargetServer != "" {
		var api ApiController
		go func() {
			for {
				serv, err := api.CheckOnlineServer(config.Server.Server.IPServer, fmt.Sprintf("%d", config.Server.Server.LoginProtocolPort))
				if err != nil {
					utils.Error("error check online server xml function, broken gorutine", err.Error())
					break
				}

				TempData.ServerStatus = serv
				time.Sleep(utils.TimeCheckInfoServer * time.Minute)
			}
		}()
	}

	return nil
}
