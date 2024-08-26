package main

import (
	"fmt"
	"net/http"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/router"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

func main() {
	err := config.Load("config.yml")
	if err != nil {
		utils.ErrorFatal(err.Error())
	}

	err = controller.LoadTemporaryData()
	if err != nil {
		utils.ErrorFatal(err.Error())
	}
	configureIP := fmt.Sprintf("%s:%d", config.VarEnviroment.ServerWeb.IP, config.VarEnviroment.ServerWeb.Port)

	r := router.NewRouter()
	if config.VarEnviroment.Certificate.ProtocolTLS {
		utils.InfoBlue(fmt.Sprintf("[HTTPS] Starting the HTTPS server: https://%s/", configureIP))
		server := &http.Server{
			Addr:    configureIP,
			Handler: r,
		}
		if err := server.ListenAndServeTLS(config.VarEnviroment.Certificate.Chain, config.VarEnviroment.Certificate.PrivKey); err != nil {
			utils.ErrorFatal("Error starting TLS server: " + err.Error())
		}
	} else {
		utils.Warn("TLS is recommended for processing HTTPS protected routes")
		utils.InfoBlue(fmt.Sprintf("[HTTP] Starting the HTTP server: http://%s/", configureIP))
		if err := http.ListenAndServe(configureIP, r); err != nil {
			utils.ErrorFatal("Error starting HTTP server: " + err.Error())
		}
	}
}
