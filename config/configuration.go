package config

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"github.com/go-yaml/yaml"
	lua "github.com/yuin/gopher-lua"
)

var PoolWorld []models.ClientWorld
var Global *models.Configuration
var SecurityPoolConnection string
var SecretPassword []byte
var Welcome string = `
____  __.                      _____                
|    |/ _|___________  ___.__. /  _  \   ____  ____  
|      < \_  __ \__  \<   |  |/  /_\  \_/ ___\/ ___\ 
|    |  \ |  | \// __ \\___  /    |    \  \__\  \___ 
|____|__ \|__|  (____  / ____\____|__  /\___  >___  >
	   \/           \/\/            \/     \/    \/ 
Created By Krayno https://www.github.com/Mau005
`

func Load(filename string) error {
	Global = &models.Configuration{}
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &Global)
	if err != nil {
		return err
	}

	utils.InfoBlueNotLog(Welcome)

	SecretPassword = []byte(GenerateRandomPassword(Global.ServerWeb.LengthSecurity))

	//Cargo la configuracion
	if Global.ServerWeb.Debug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.SetPrefix("DEBUG: ")
		SecretPassword = []byte("debugging")
		utils.Warn("Security mode Debug")
	} else {
		utils.Info(fmt.Sprintf("Security Active %d\n", Global.ServerWeb.LengthSecurity))
	}
	if Global.ServerWeb.EnvironmentVariables {
		Global.DB.Host = os.Getenv("MYSQL_HOST")
		Global.DB.Port = parseEnvUint(os.Getenv("MYSQL_PORT"), 3306)
		Global.DB.UserName = os.Getenv("MYSQL_USER")
		Global.DB.DBPassword = os.Getenv("MYSQL_PASSWORD")
		Global.DB.DataBase = os.Getenv("MYSQL_DATABASE")
		utils.Info("MySQL variables are loaded from environment")
	} else {
		utils.Info("MySQL variables are loaded to config.yml")
	}

	//securityToken
	SecurityPoolConnection = os.Getenv("KRAY_PASSWORD")
	if SecurityPoolConnection == "" {
		//If there is no API token, it will generate a random token of 500 length so that it cannot be accessed by anyone.
		SecurityPoolConnection = GenerateRandomPassword(500)
	}

	for _, value := range Global.PoolServer {
		if value.IpWebApi == "" || value.Token == "" {
			return errors.New("error pool server no empty string, configure in config.yml")
		}
	}

	if Global.ServerWeb.TargetServer != "" {
		err = LoadConfigLua(Global.ServerWeb.TargetServer)
		if err != nil {
			return err
		}
	} else {
		utils.Warn("configurate Target Server not found")
	}

	err = db.ConnectionMysql(
		Global.DB.UserName,
		Global.DB.DBPassword,
		Global.DB.Host,
		Global.DB.DataBase,
		Global.DB.Port,
		Global.ServerWeb.Debug,
	)
	if err != nil {
		return err
	}

	utils.Info("Environment variables are added OK")
	return nil
}
func parseEnvUint(key string, defaultValue uint16) uint16 {
	if value, exists := os.LookupEnv(key); exists {
		if parsedValue, err := strconv.ParseUint(value, 10, 16); err == nil {
			return uint16(parsedValue)
		}
	}
	return defaultValue
}

func generateRandomCharacter(charset string) string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		log.Fatal(err)
	}
	return string(charset[n.Int64()])
}

func GenerateRandomPassword(length int) string {
	allChars := utils.UpperCase + utils.LowerCase + utils.Digits + utils.Special
	password := ""
	for i := 0; i < length; i++ {
		randomChar := generateRandomCharacter(allChars)
		password += randomChar
	}
	return password
}

func LoadConfigLua(targetServer string) (err error) {
	checkOS := runtime.GOOS
	// targetExecute := ""
	targetPath := ""
	switch checkOS {

	case "windows":
		slicePath := strings.Split(targetServer, "\\")
		// targetExecute = fmt.Sprintf("%s.exe", slicePath[len(slicePath)-1])
		targetPath = strings.Join(slicePath[:len(slicePath)-1], "\\")

	default:
		slicePath := strings.Split(targetServer, "/")
		// targetExecute = fmt.Sprintf("./%s", slicePath[len(slicePath)-1])
		targetPath = strings.Join(slicePath[:len(slicePath)-1], "/")
	}

	// PathServer := targetPath
	// NameExecute := targetExecute

	L := lua.NewState()
	defer L.Close()
	path_new := fmt.Sprintf("%s/%s", targetPath, "config.lua")
	if err := L.DoFile(path_new); err != nil {
		return err
	}
	var WorldType int
	switch L.GetGlobal("worldType").String() {
	case "pvp":
		WorldType = 0
	case "no-pvp":
		WorldType = 1
	case "pvp-enforced":
		WorldType = 2
	default:
		WorldType = 0
	}
	IPServer := L.GetGlobal("ip").String()
	// LoginProtocolPort, err := strconv.ParseUint(L.GetGlobal("loginProtocolPort").String(), 10, 16)
	// if err != nil {
	// 	return
	// }

	GameProtocolPort, err := strconv.ParseUint(L.GetGlobal("gameProtocolPort").String(), 10, 16)
	if err != nil {
		return
	}

	StatusProtocolPort, err := strconv.ParseUint(L.GetGlobal("statusProtocolPort").String(), 10, 16)
	if err != nil {
		return
	}

	Location := L.GetGlobal("location").String()
	NameServer := L.GetGlobal("serverName").String()
	//HouseRentPeriod := L.GetGlobal("houseRentPeriod").String()

	rateExp, err := strconv.ParseUint(L.GetGlobal("rateExp").String(), 10, 16)
	if err != nil {
		return
	}

	rateSkill, err := strconv.ParseUint(L.GetGlobal("rateSkill").String(), 10, 16)
	if err != nil {
		return
	}

	rateLoot, err := strconv.ParseUint(L.GetGlobal("rateLoot").String(), 10, 16)
	if err != nil {
		return
	}

	RateMagic, err := strconv.ParseUint(L.GetGlobal("rateMagic").String(), 10, 16)
	if err != nil {
		return
	}

	rateSpawn, err := strconv.ParseUint(L.GetGlobal("rateSpawn").String(), 10, 16)
	if err != nil {
		return
	}
	rate := models.RateServer{RateExp: uint16(rateExp),
		RateSkill: uint16(rateSkill),
		RateLoot:  uint16(rateLoot),
		RateMagic: uint16(RateMagic),
		RateSpawn: uint16(rateSpawn)}

	World := models.ClientWorld{
		ID:                         len(Global.PoolServer),
		AntiCheatProtection:        false,
		ExternalAddRessUnProtected: IPServer,
		ExternalAddress:            IPServer,
		ExternalAddressProtected:   IPServer,
		PvpType:                    uint8(WorldType),
		Location:                   Location,
		Name:                       NameServer,
		PreviewState:               1,
		ExternalPort:               uint16(StatusProtocolPort),
		ExternalPortProtected:      uint16(GameProtocolPort),
		ExternalPortUnprotected:    uint16(GameProtocolPort),
		CurrentTournamentPhase:     2,
	}

	Global.PoolServer = append(Global.PoolServer, models.PoolServer{World: World, RateServer: rate})
	utils.Info("configure server local")
	utils.Info("loaded config.lua")
	return
}
