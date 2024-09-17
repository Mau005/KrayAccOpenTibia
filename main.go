package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Mau005/KrayAccOpenTibia/config"
	"github.com/Mau005/KrayAccOpenTibia/controller"
	"github.com/Mau005/KrayAccOpenTibia/router"
	"github.com/Mau005/KrayAccOpenTibia/utils"
)

func main() {
	preparing := flag.Bool("security", false, "Security Pool Connection ")
	flag.Parse()
	if *preparing {
		security()
	}

	err := config.Load("config.yml")
	if err != nil {
		utils.ErrorFatal(err.Error())
	}

	err = controller.LoadTemporaryData()
	if err != nil {
		utils.ErrorFatal(err.Error())
	}

	configureIP := fmt.Sprintf("%s:%d", config.Global.ServerWeb.IP, config.Global.ServerWeb.Port)

	r := router.NewRouter()

	if config.Global.Certificate.ProtocolTLS {
		utils.InfoBlue(fmt.Sprintf("[HTTPS] Starting the HTTPS server: https://%s/", configureIP))
		server := &http.Server{
			Addr:    configureIP,
			Handler: r,
		}
		if err := server.ListenAndServeTLS(config.Global.Certificate.Chain, config.Global.Certificate.PrivKey); err != nil {
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

func security() {

	var pressContinue string
	fmt.Println("Welcome to the connection pool configuration for multiple APIs")
	fmt.Println("The token is generated from the environment password 'KRAY_PASSWORD', if you do not configure this attribute this command will not be executed.")
	varPass := os.Getenv("KRAY_PASSWORD")

	if varPass == "" {
		utils.ErrorFatal("set the password in your work environment as export KRAY_PASSWORD=YOUR_PASSWORD")
	}
	fmt.Println("Press a key to continue...")
	_, err := fmt.Scanln(&pressContinue)
	if err != nil {
	}

	var apiCtl controller.ApiController
	token, err := apiCtl.GenerateJWTPoolConnection(varPass)
	if err != nil {
		utils.ErrorFatal(err.Error())
	}
	utils.Info("has been generated successfully")
	utils.Info("Token:")
	utils.InfoSuccess(token)

	os.Exit(0)

}
