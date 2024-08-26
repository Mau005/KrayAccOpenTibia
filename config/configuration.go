package config

import (
	"crypto/rand"
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

var Server ExecuteServer
var VarEnviroment *Configuration
var SecretPassword []byte
var Welcome string = `
____  __.                      _____                
|    |/ _|___________  ___.__. /  _  \   ____  ____  
|      < \_  __ \__  \<   |  |/  /_\  \_/ ___\/ ___\ 
|    |  \ |  | \// __ \\___  /    |    \  \__\  \___ 
|____|__ \|__|  (____  / ____\____|__  /\___  >___  >
	   \/           \/\/            \/     \/    \/ 
   `

type ServerWeb struct {
	IP                   string `yaml:"IP"`
	Port                 uint16 `yaml:"Port"`
	Debug                bool   `yaml:"Debug"`
	ApiMode              bool   `yaml:"ApiMode"`
	LengthSecurity       int    `yaml:"LengthSecurity"`
	EnvironmentVariables bool   `yaml:"EnvironmentVariables"`
	UrlItemView          string `yaml:"UrlItemView"`
	UrlOutfitsView       string `yaml:"UrlOutfitsView"`
	TargetServer         string `yaml:"TargetServer"`
	LimitCreateCharacter uint8  `yaml:"LimitCreateCharacter"`
}
type MySQL struct {
	Host       string `yaml:"Host"`
	Port       uint16 `yaml:"Port"`
	UserName   string `yaml:"UserName"`
	DBPassword string `yaml:"DBPassword"`
	DataBase   string `yaml:"DataBase"`
}

type Certificate struct {
	ProtocolTLS bool   `yaml:"ProtocolHTTPS"`
	Chain       string `yaml:"Chain"`
	PrivKey     string `yaml:"PrivKey"`
}

type Configuration struct {
	DB                   MySQL                         `yaml:"MySQL"`
	ServerWeb            ServerWeb                     `yaml:"ServerWeb"`
	Certificate          Certificate                   `yaml:"Certificate"`
	ServerConnectKrayAcc []models.ServerConnectKrayAcc `yaml:"ApiConnectionPool"`
}

func Load(filename string) error {

	VarEnviroment = &Configuration{}
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &VarEnviroment)
	if err != nil {
		return err
	}

	utils.InfoBlue(Welcome)

	SecretPassword = []byte(GenerateRandomPassword(VarEnviroment.ServerWeb.LengthSecurity))

	//Cargo la configuracion
	if VarEnviroment.ServerWeb.Debug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.SetPrefix("DEBUG: ")
		SecretPassword = []byte("debugging")
		utils.Warn("Security mode Debug")
	} else {
		utils.Info(fmt.Sprintf("Security Active %d\n", VarEnviroment.ServerWeb.LengthSecurity))
	}
	if VarEnviroment.ServerWeb.EnvironmentVariables {
		utils.Info("MySQL variables are loaded from environment")
		VarEnviroment.DB.Host = os.Getenv("MYSQL_HOST")
		VarEnviroment.DB.Port = parseEnvUint(os.Getenv("MYSQL_PORT"), 3306)
		VarEnviroment.DB.UserName = os.Getenv("MYSQL_USER")
		VarEnviroment.DB.DBPassword = os.Getenv("MYSQL_PASSWORD")
		VarEnviroment.DB.DataBase = os.Getenv("MYSQL_DATABASE")
	} else {
		utils.Info("MySQL variables are loaded to config.yml")
	}
	Server, err = LoadConfigLua(VarEnviroment.ServerWeb.TargetServer)
	if err != nil {
		utils.ErrorFatal(err.Error())
	}
	err = loadMySQL()
	if err != nil {
		utils.ErrorFatal(err.Error())
	}

	utils.Info("Environment variables are added OK")
	return nil
}

func loadMySQL() error {
	return db.ConnectionMysql(
		VarEnviroment.DB.UserName,
		VarEnviroment.DB.DBPassword,
		VarEnviroment.DB.Host,
		VarEnviroment.DB.DataBase,
		VarEnviroment.DB.Port,
		VarEnviroment.ServerWeb.Debug,
	)
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

func LoadConfigLua(targetServer string) (preConfigServer ExecuteServer, err error) {
	checkOS := runtime.GOOS
	targetExecute := ""
	targetPath := ""
	switch checkOS {

	case "windows":
		slicePath := strings.Split(targetServer, "\\")
		targetExecute = fmt.Sprintf("%s.exe", slicePath[len(slicePath)-1])
		targetPath = strings.Join(slicePath[:len(slicePath)-1], "\\")

	default:
		slicePath := strings.Split(targetServer, "/")
		targetExecute = fmt.Sprintf("./%s", slicePath[len(slicePath)-1])
		targetPath = strings.Join(slicePath[:len(slicePath)-1], "/")
	}

	preConfigServer.PathServer = targetPath
	preConfigServer.NameExecute = targetExecute

	L := lua.NewState()
	defer L.Close()
	path_new := fmt.Sprintf("%s/%s", targetPath, "config.lua")
	if err := L.DoFile(path_new); err != nil {
		utils.ErrorFatal("error executing Lua script:", err.Error())
	}
	switch L.GetGlobal("worldType").String() {
	case "pvp":
		preConfigServer.Server.WorldType = 0
	case "no-pvp":
		preConfigServer.Server.WorldType = 1
	case "pvp-enforced":
		preConfigServer.Server.WorldType = 2
	default:
		preConfigServer.Server.WorldType = 0
	}
	preConfigServer.Server.IPServer = L.GetGlobal("ip").String()
	LoginProtocolPort, err := strconv.ParseUint(L.GetGlobal("loginProtocolPort").String(), 10, 16)
	if err != nil {
		return
	}
	preConfigServer.Server.LoginProtocolPort = uint16(LoginProtocolPort)

	gameProtocolPort, err := strconv.ParseUint(L.GetGlobal("gameProtocolPort").String(), 10, 16)
	if err != nil {
		return
	}
	preConfigServer.Server.GameProtocolPort = uint16(gameProtocolPort)

	statusProtocolPort, err := strconv.ParseUint(L.GetGlobal("statusProtocolPort").String(), 10, 16)
	if err != nil {
		return
	}
	preConfigServer.Server.StatusProtocolPort = uint16(statusProtocolPort)

	preConfigServer.NameServer = L.GetGlobal("serverName").String()
	preConfigServer.Server.HouseRentPeriod = L.GetGlobal("houseRentPeriod").String()

	rateExp, err := strconv.ParseUint(L.GetGlobal("rateExp").String(), 10, 16)
	if err != nil {
		return
	}
	preConfigServer.RateServer.RateExp = uint16(rateExp)

	rateSkill, err := strconv.ParseUint(L.GetGlobal("rateSkill").String(), 10, 16)
	if err != nil {
		return
	}
	preConfigServer.RateServer.RateSkill = uint16(rateSkill)

	rateLoot, err := strconv.ParseUint(L.GetGlobal("rateLoot").String(), 10, 16)
	if err != nil {
		return
	}
	preConfigServer.RateServer.RateLoot = uint16(rateLoot)

	RateMagic, err := strconv.ParseUint(L.GetGlobal("rateMagic").String(), 10, 16)
	if err != nil {
		return
	}
	preConfigServer.RateServer.RateMagic = uint16(RateMagic)

	rateSpawn, err := strconv.ParseUint(L.GetGlobal("rateSpawn").String(), 10, 16)
	if err != nil {
		return
	}
	preConfigServer.RateServer.RateSpawn = uint16(rateSpawn)

	utils.Info("loaded config.lua")
	return
}
